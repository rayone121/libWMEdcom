// Package winmentor provides a Go client for the WinMENTOR DocImpServer DCOM interface.
//
// DocImpServer is a COM/DCOM automation server that allows external applications to
// read nomenclatures, query stock/balances, and import documents (invoices, orders, etc.)
// into the WinMENTOR accounting application.
//
// The client is safe to use from any goroutine. All COM calls are dispatched to a
// dedicated goroutine that owns the COM thread.
//
// Basic usage:
//
//	client, err := winmentor.NewClient()
//	if err != nil { log.Fatal(err) }
//	defer client.Close()
//
//	err = client.SetNumeFirma("MYCOMPANY")
//	err = client.SetLunaLucru(2024, 10)
//	partners, err := client.GetListaParteneri()
package winmentor

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var (
	ErrNotConnected = errors.New("winmentor: not connected to DocImpServer")
	ErrSetFirma     = errors.New("winmentor: failed to set company name")
	ErrSetLuna      = errors.New("winmentor: failed to set work month")
)

// Client wraps a connection to the DocImpServer COM object.
// All COM operations are dispatched to a dedicated goroutine that owns the OS thread.
type Client struct {
	obj     *ole.IDispatch // only accessed from the COM goroutine
	unknown *ole.IUnknown  // only accessed from the COM goroutine
	comCh   chan comReq     // channel for dispatching work to the COM goroutine
	done    chan struct{}   // closed when the COM goroutine exits
}

// NewClient creates a new COM connection to DocImpServer.
// It spawns a dedicated goroutine that owns the COM thread. The client is safe
// to use from any goroutine.
func NewClient() (*Client, error) {
	c := &Client{
		comCh: make(chan comReq),
		done:  make(chan struct{}),
	}

	errCh := make(chan error, 1)
	go c.comLoop(errCh)

	if err := <-errCh; err != nil {
		return nil, err
	}
	return c, nil
}

// comLoop runs on a dedicated goroutine, locked to an OS thread.
// It initializes COM, creates the DocImpServer object, and processes requests
// until the channel is closed.
func (c *Client) comLoop(errCh chan<- error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		oleErr, ok := err.(*ole.OleError)
		if !ok || oleErr.Code() != 0x00000001 { // S_FALSE = already initialized
			errCh <- fmt.Errorf("CoInitializeEx: %w", err)
			return
		}
	}
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("DocImpServer.DocImpObjectClass")
	if err != nil {
		errCh <- fmt.Errorf("CreateObject DocImpServer.DocImpObjectClass: %w", err)
		return
	}

	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		unknown.Release()
		errCh <- fmt.Errorf("QueryInterface IDispatch: %w", err)
		return
	}

	c.obj = disp
	c.unknown = unknown
	close(errCh) // signal success

	// Process requests until channel is closed.
	for req := range c.comCh {
		req.fn()
		close(req.done)
	}

	// Cleanup COM resources on the thread that created them.
	disp.Release()
	unknown.Release()
	c.obj = nil
	c.unknown = nil
	close(c.done)
}

// Close releases the COM object, uninitializes COM, and terminates the COM goroutine.
// After Close returns, the client must not be used.
func (c *Client) Close() {
	close(c.comCh)
	<-c.done
}

// --- Session Setup ---

// GetListaFirme returns the list of company short names available in the WinMENTOR data directory.
func (c *Client) GetListaFirme() ([]string, error) {
	return c.callReturningStrings("GetListaFirme")
}

// GetListaLuni returns the list of work months (format "yyyy_mm") for the given company.
func (c *Client) GetListaLuni(numeFirma string) ([]string, error) {
	return c.callReturningStrings("GetListaLuni", numeFirma)
}

// SetNumeFirma sets the active company. Returns nil on success.
func (c *Client) SetNumeFirma(numeFirma string) error {
	result, err := c.callMethodInt("SetNumeFirma", numeFirma)
	if err != nil {
		return err
	}
	if result == 0 {
		errs, _ := c.GetListaErori()
		if len(errs) > 0 {
			return fmt.Errorf("%w: %s", ErrSetFirma, strings.Join(errs, "; "))
		}
		return ErrSetFirma
	}
	return nil
}

// SetLunaLucru sets the working month (year, month). Returns nil on success.
func (c *Client) SetLunaLucru(an, luna int) error {
	result, err := c.callMethodInt("SetLunaLucru", an, luna)
	if err != nil {
		return err
	}
	if result == 0 {
		errs, _ := c.GetListaErori()
		if len(errs) > 0 {
			return fmt.Errorf("%w: %s", ErrSetLuna, strings.Join(errs, "; "))
		}
		return ErrSetLuna
	}
	return nil
}

// GetListaErori returns the list of error messages from the last operation.
func (c *Client) GetListaErori() ([]string, error) {
	return c.callReturningStrings("GetListaErori")
}

// SetDocsData sends a document data packet (array of strings) to the server.
// This must be called before DateValide/ImportaFacturi/ComenziValide/etc.
//
// The Delphi equivalent creates a VarArray of OleStr. go-ole's CallMethod
// automatically converts []string to a SAFEARRAY of BSTR.
func (c *Client) SetDocsData(lines []string) error {
	return c.callMethodVoid("SetDocsData", lines)
}

// --- Configuration ---

// SetIDPartField sets the field used for partner identification.
// Valid values: "CodExtern", "CodFiscal", "CodIntern".
func (c *Client) SetIDPartField(fieldName string) error {
	return c.callMethodVoid("SetIDPartField", fieldName)
}

// SetIDArtField sets the field used for article identification.
// Valid values: "CodExtern", "CodIntern".
func (c *Client) SetIDArtField(fieldName string) error {
	return c.callMethodVoid("SetIDArtField", fieldName)
}

// GetVersiuni calls the COM method that returns WinMENTOR and server versions.
// DLL signature: GetVersiuni(out VerMentor, VerServer: Double): Integer
// Returns (result_code, mentor_version, server_version, error).
func (c *Client) GetVersiuni() (resultCode int, verMentor float64, verServer float64, err error) {
	c.comDo(func() {
		var vMentor, vServer float64
		mentorVariant := ole.NewVariant(ole.VT_R8|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(&vMentor))))
		serverVariant := ole.NewVariant(ole.VT_R8|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(&vServer))))

		var v *ole.VARIANT
		v, err = c.rawCall("GetVersiuni", &mentorVariant, &serverVariant)
		if err != nil {
			return
		}
		defer v.Clear()

		resultCode = int(v.Val)
		verMentor = vMentor
		verServer = vServer
	})
	return
}

// GetStringConstanta returns a string constant by ID and Symbol.
func (c *Client) GetStringConstanta(id int, simbol string) (string, error) {
	return c.callMethodString("GetStringConstanta", id, simbol)
}

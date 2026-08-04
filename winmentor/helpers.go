package winmentor

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// comReq represents a unit of work to be executed on the COM goroutine.
type comReq struct {
	fn   func()
	done chan struct{}
}

// comDo executes fn on the dedicated COM goroutine and blocks until it completes.
// This ensures all COM calls happen on the thread that initialized COM.
func (c *Client) comDo(fn func()) {
	done := make(chan struct{})
	c.comCh <- comReq{fn: fn, done: done}
	<-done
}

// rawGetListaErori retrieves the server's error list. Must only be called from
// the COM goroutine, which is why it goes through the vtable directly rather
// than GetListaErori — that one takes the comDo lock this is already holding.
func (c *Client) rawGetListaErori() ([]string, error) {
	if c.vtbl == nil {
		return nil, ErrNotConnected
	}
	var resVar ole.VARIANT
	if _, err := c.vtblCall("GetListaErori", unsafe.Pointer(&resVar)); err != nil {
		return nil, err
	}
	defer resVar.Clear()
	return variantToStrings(&resVar)
}

// vtblMethod looks a method up and checks it takes the number of vtable slots
// the caller is about to fill.
//
// Every helper below used to fall through to an IDispatch retry when this did
// not match. That retry could never work — the server does not implement
// IDispatch::Invoke — so it only ever replaced the real problem with "the server
// threw an exception". With it gone a mismatch must be reported, or the helper
// returns success and no data.
func (c *Client) vtblMethod(name string, want ...int) (vtableMethod, error) {
	if c.vtbl == nil {
		return vtableMethod{}, ErrNotConnected
	}
	m, ok := c.vtbl.methods[name]
	if !ok {
		return vtableMethod{}, fmt.Errorf("%s is not in the type library", name)
	}
	for _, w := range want {
		if m.nParams == w {
			return m, nil
		}
	}
	return vtableMethod{}, fmt.Errorf("%s takes %d vtable parameters, caller supplied %v",
		name, m.nParams, want)
}

// callMethodInt invokes a COM method that returns an Integer.
func (c *Client) callMethodInt(name string, args ...interface{}) (result int, err error) {
	c.comDo(func() {
		// Most Delphi dual-interface methods returning Integer are
		// HRESULT Method(..., [out, retval] int*); a few return directly.
		m, e := c.vtblMethod(name, len(args)+1, len(args))
		if e != nil {
			err = e
			return
		}
		if m.nParams == len(args)+1 {
			var resVal int32
			if _, err = c.vtblCall(name, append(args, unsafe.Pointer(&resVal))...); err == nil {
				result = int(resVal)
			}
			return
		}
		var res uintptr
		if res, err = c.vtblCall(name, args...); err == nil {
			result = int(res)
		}
	})
	return
}

// callMethodVoid invokes a void COM method, ensuring the VARIANT is properly cleared.
func (c *Client) callMethodVoid(name string, args ...interface{}) (err error) {
	c.comDo(func() {
		// Even void methods return HRESULT, so nParams matches exactly.
		if _, e := c.vtblMethod(name, len(args)); e != nil {
			err = e
			return
		}
		_, err = c.vtblCall(name, args...)
	})
	return
}

// callMethodString invokes a COM method that returns a WideString (BSTR).
func (c *Client) callMethodString(name string, args ...interface{}) (result string, err error) {
	c.comDo(func() {
		if _, e := c.vtblMethod(name, len(args)+1); e != nil {
			err = e
			return
		}
		var resPtr *uint16
		if _, err = c.vtblCall(name, append(args, unsafe.Pointer(&resPtr))...); err != nil {
			return
		}
		result = ole.BstrToString(resPtr)
		ole.SysFreeString((*int16)(unsafe.Pointer(resPtr)))
	})
	return
}

// variantToStrings converts an OleVariant (SafeArray of strings) to a Go []string.
func variantToStrings(v *ole.VARIANT) ([]string, error) {
	if v == nil || v.VT == ole.VT_EMPTY || v.VT == ole.VT_NULL {
		return nil, nil
	}

	// If it's a BSTR directly, return single element.
	if v.VT == ole.VT_BSTR {
		s := v.ToString()
		if s == "" {
			return nil, nil
		}
		return []string{s}, nil
	}

	// Expect a SafeArray of variants or BSTRs.
	safeArray := v.ToArray()
	if safeArray == nil {
		return nil, fmt.Errorf("expected SafeArray, got VT=%d", v.VT)
	}

	arr := safeArray.ToValueArray()
	result := make([]string, 0, len(arr))
	for _, item := range arr {
		switch val := item.(type) {
		case string:
			result = append(result, val)
		default:
			result = append(result, fmt.Sprintf("%v", val))
		}
	}
	return result, nil
}

// callWithOutError calls a COM method that has an "out Error: Integer" parameter.
//
// Five readers do NOT have one — GetListaFirme, GetListaErori, GetListaLuni,
// GetIncasariFactura and GetPlatiFactura — and for those the parameter count
// alone is not enough to tell. GetIncasariFactura takes five [in] params plus a
// retval, so nParams is 6, and a caller passing four arguments satisfies
// len(extraArgs)+2 == 6 by coincidence. The call then went ahead and put
// &errParam into the fifth slot, where the server expects a BSTR, dereferenced
// it as a string, and took the process down with an access violation.
//
// So the shape is checked, not just the count: the slot before the retval has
// to be a by-reference integer.
func (c *Client) callWithOutError(name string, extraArgs ...interface{}) (result []string, err error) {
	c.comDo(func() {
		// HRESULT Method(..., [out] int* err, [out, retval] VARIANT* result)
		m, e := c.vtblMethod(name, len(extraArgs)+2)
		if e != nil {
			err = e
			return
		}
		if !hasOutError(m) {
			err = fmt.Errorf("%s has no [out] Error parameter; calling it this way would pass "+
				"a pointer where the server expects a value", name)
			return
		}

		var errParam int32
		var resVar ole.VARIANT
		if _, err = c.vtblCall(name, append(extraArgs, unsafe.Pointer(&errParam), unsafe.Pointer(&resVar))...); err != nil {
			return
		}
		defer resVar.Clear()
		if errParam != 0 {
			if errs, _ := c.rawGetListaErori(); len(errs) > 0 {
				err = fmt.Errorf("%s failed: %s", name, strings.Join(errs, "; "))
				return
			}
			err = fmt.Errorf("%s failed with error code %d", name, errParam)
			return
		}
		result, err = variantToStrings(&resVar)
	})
	return
}

// callReturningStrings calls a COM method that returns an OleVariant directly
// (without an out Error parameter).
func (c *Client) callReturningStrings(name string, args ...interface{}) (result []string, err error) {
	c.comDo(func() {
		// HRESULT Method(..., [out, retval] VARIANT*)
		if _, e := c.vtblMethod(name, len(args)+1); e != nil {
			err = e
			return
		}
		var resVar ole.VARIANT
		if _, err = c.vtblCall(name, append(args, unsafe.Pointer(&resVar))...); err != nil {
			return
		}
		defer resVar.Clear()
		result, err = variantToStrings(&resVar)
	})
	return
}

// RawQuery calls a COM method by name and returns the raw semicolon-separated
// string records before any field splitting. Useful for debugging field counts.
func (c *Client) RawQuery(name string, extraArgs ...interface{}) ([]string, error) {
	return c.callWithOutError(name, extraArgs...)
}

// splitFields splits a semicolon-separated record into fields.
func splitFields(record string, count int) []string {
	parts := strings.SplitN(record, ";", count)
	// Pad with empty strings if fewer fields returned.
	for len(parts) < count {
		parts = append(parts, "")
	}
	return parts
}

package winmentor

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
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

// rawCall invokes a COM method by name. Must only be called from the COM goroutine.
func (c *Client) rawCall(name string, args ...interface{}) (*ole.VARIANT, error) {
	if c.obj == nil {
		return nil, ErrNotConnected
	}
	result, err := oleutil.CallMethod(c.obj, name, args...)
	if err != nil {
		return nil, fmt.Errorf("CallMethod %s: %w", name, err)
	}
	return result, nil
}

// rawGetListaErori retrieves error messages. Must only be called from the COM goroutine.
func (c *Client) rawGetListaErori() ([]string, error) {
	v, err := c.rawCall("GetListaErori")
	if err != nil {
		return nil, err
	}
	defer v.Clear()
	return variantToStrings(v)
}

// callMethodInt invokes a COM method that returns an Integer.
func (c *Client) callMethodInt(name string, args ...interface{}) (result int, err error) {
	c.comDo(func() {
		var v *ole.VARIANT
		v, err = c.rawCall(name, args...)
		if err != nil {
			return
		}
		defer v.Clear()
		result = int(v.Val)
	})
	return
}

// callMethodVoid invokes a void COM method, ensuring the VARIANT is properly cleared.
func (c *Client) callMethodVoid(name string, args ...interface{}) (err error) {
	c.comDo(func() {
		var v *ole.VARIANT
		v, err = c.rawCall(name, args...)
		if err != nil {
			return
		}
		v.Clear()
	})
	return
}

// callMethodString invokes a COM method that returns a WideString (BSTR).
func (c *Client) callMethodString(name string, args ...interface{}) (result string, err error) {
	c.comDo(func() {
		var v *ole.VARIANT
		v, err = c.rawCall(name, args...)
		if err != nil {
			return
		}
		defer v.Clear()
		result = v.ToString()
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
func (c *Client) callWithOutError(name string, extraArgs ...interface{}) (result []string, err error) {
	c.comDo(func() {
		var errParam int32
		errVariant := ole.NewVariant(ole.VT_I4|ole.VT_BYREF, int64(uintptr(unsafe.Pointer(&errParam))))

		args := make([]interface{}, 0, len(extraArgs)+1)
		args = append(args, extraArgs...)
		args = append(args, &errVariant)

		var v *ole.VARIANT
		v, err = c.rawCall(name, args...)
		if err != nil {
			return
		}
		defer v.Clear()

		if errParam != 0 {
			errs, _ := c.rawGetListaErori()
			if len(errs) > 0 {
				err = fmt.Errorf("%s failed: %s", name, strings.Join(errs, "; "))
				return
			}
			err = fmt.Errorf("%s failed with error code %d", name, errParam)
			return
		}

		result, err = variantToStrings(v)
	})
	return
}

// callReturningStrings calls a COM method that returns an OleVariant directly
// (without an out Error parameter).
func (c *Client) callReturningStrings(name string, args ...interface{}) (result []string, err error) {
	c.comDo(func() {
		var v *ole.VARIANT
		v, err = c.rawCall(name, args...)
		if err != nil {
			return
		}
		defer v.Clear()
		result, err = variantToStrings(v)
	})
	return
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

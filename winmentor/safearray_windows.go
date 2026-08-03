//go:build windows

package winmentor

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// go-ole keeps its SAFEARRAY wrappers unexported, so call oleaut32 directly. Needed
// for SetDocsData, whose Delphi signature is SetDocsData(DocData: OleVariant) and which
// expects a VARIANT wrapping a SAFEARRAY of BSTR — the COM equivalent of the manual's
// VarArrayCreate([0, N-1], varOleStr).
var (
	modOleAut32               = syscall.NewLazyDLL("oleaut32.dll")
	procSafeArrayCreateVector = modOleAut32.NewProc("SafeArrayCreateVector")
	procSafeArrayPutElement   = modOleAut32.NewProc("SafeArrayPutElement")
	procSafeArrayDestroy      = modOleAut32.NewProc("SafeArrayDestroy")
)

// newBSTRSafeArray builds a zero-based SAFEARRAY of BSTR from items.
//
// The lower bound is 0 to match the manual's VarArrayCreate([0, Count-1], varOleStr).
// SafeArrayPutElement copies the BSTR, so each temporary is freed immediately after;
// SafeArrayDestroy then owns the copies.
func newBSTRSafeArray(items []string) (*ole.SafeArray, error) {
	sa, _, _ := procSafeArrayCreateVector.Call(
		uintptr(ole.VT_BSTR),
		uintptr(0),
		uintptr(uint32(len(items))),
	)
	if sa == 0 {
		return nil, fmt.Errorf("SafeArrayCreateVector(VT_BSTR, 0, %d) returned NULL", len(items))
	}
	arr := (*ole.SafeArray)(unsafe.Pointer(sa))

	for i, s := range items {
		bstr := ole.SysAllocStringLen(s)
		if bstr == nil {
			destroySafeArray(arr)
			return nil, fmt.Errorf("SysAllocStringLen failed for element %d", i)
		}

		idx := int32(i)
		hr, _, _ := procSafeArrayPutElement.Call(
			sa,
			uintptr(unsafe.Pointer(&idx)),
			uintptr(unsafe.Pointer(bstr)),
		)
		ole.SysFreeString(bstr)

		if int32(hr) < 0 {
			destroySafeArray(arr)
			return nil, fmt.Errorf("SafeArrayPutElement(%d) failed: HRESULT 0x%08X", i, uint32(hr))
		}
	}

	return arr, nil
}

func destroySafeArray(sa *ole.SafeArray) {
	if sa != nil {
		procSafeArrayDestroy.Call(uintptr(unsafe.Pointer(sa)))
	}
}

// variantOfBSTRArray wraps a SAFEARRAY of BSTR in a VARIANT of type VT_ARRAY|VT_BSTR.
// The VARIANT does not own the array — the caller destroys it.
func variantOfBSTRArray(sa *ole.SafeArray) ole.VARIANT {
	var v ole.VARIANT
	v.VT = ole.VT_ARRAY | ole.VT_BSTR
	v.Val = int64(uintptr(unsafe.Pointer(sa)))
	return v
}

// variantStackSlots splits a VARIANT into the machine words a by-value [in] VARIANT
// parameter occupies on the stack: 4 on 386 (16 bytes), 3 on amd64 (24 bytes).
func variantStackSlots(v *ole.VARIANT) []uintptr {
	n := unsafe.Sizeof(*v) / ptrSize()
	words := (*[8]uintptr)(unsafe.Pointer(v))
	out := make([]uintptr, n)
	copy(out, words[:n])
	return out
}

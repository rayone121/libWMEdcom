//go:build !windows

package winmentor

import (
	"fmt"
	"github.com/go-ole/go-ole"
)

type vtableInfo struct {
	methods map[string]vtableMethod
}

type vtableMethod struct {
	nParams int
}

func newVTableInfo(disp *ole.IDispatch) (*vtableInfo, error) {
	return nil, fmt.Errorf("vtable only supported on windows")
}

func (vi *vtableInfo) Release() {}

func (c *Client) vtblCall(name string, args ...interface{}) (uintptr, error) {
	return 0, fmt.Errorf("vtable only supported on windows")
}

// hasOutError is unanswerable without a type library, so the stub never blocks
// a call. See the windows implementation for what it guards against.
func hasOutError(vtableMethod) bool { return true }

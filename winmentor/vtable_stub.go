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

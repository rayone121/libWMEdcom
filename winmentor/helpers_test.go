package winmentor

import (
	"errors"
	"strings"
	"testing"
)

// The IDispatch fallback is gone, so a method that is absent or whose arity does
// not match must now produce an error. It used to fall through to a retry the
// server cannot serve, which returned "the server threw an exception" and buried
// the real cause; with the retry removed, silence would be worse — the helper
// would return success and no data.
func TestVtblMethodReportsMismatchesInsteadOfGoingQuiet(t *testing.T) {
	var c Client
	if _, err := c.vtblMethod("GetListaFirme", 1); !errors.Is(err, ErrNotConnected) {
		t.Errorf("no type library should give ErrNotConnected, got %v", err)
	}

	c.vtbl = &vtableInfo{methods: map[string]vtableMethod{
		"GetListaGestiuni": {nParams: 2},
	}}

	if _, err := c.vtblMethod("NuExista", 2); err == nil {
		t.Error("an unknown method must be reported")
	} else if !strings.Contains(err.Error(), "type library") {
		t.Errorf("unhelpful error for an unknown method: %v", err)
	}

	if _, err := c.vtblMethod("GetListaGestiuni", 5); err == nil {
		t.Error("an arity mismatch must be reported")
	} else if !strings.Contains(err.Error(), "2") {
		t.Errorf("the error should name the real parameter count: %v", err)
	}

	if m, err := c.vtblMethod("GetListaGestiuni", 3, 2); err != nil {
		t.Errorf("a matching arity should be accepted: %v", err)
	} else if m.nParams != 2 {
		t.Errorf("returned the wrong method: %+v", m)
	}
}

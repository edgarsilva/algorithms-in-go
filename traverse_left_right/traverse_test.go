package traverseleftright

import (
	"testing"
)

func TestAssert(t *testing.T) {
	passes := Assert()

	if !passes {
		t.Errorf("Assert Traverse left and right failed")
	}
}

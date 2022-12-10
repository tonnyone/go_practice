package error_my

import (
	"runtime/debug"
	"testing"
)

func TestDebug(t *testing.T) {
	err := Test3()
	if err != nil {
		t.Log(string(debug.Stack()))
	}
}

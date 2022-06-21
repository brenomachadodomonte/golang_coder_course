package architecture

import (
	"runtime"
	"testing"
)

func TestDependency(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("It doesn't work with amd64")
	}

	t.Fail()
}

package s2

import (
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("skip MacOS")
	}

	t.Log("there a non-skipped log")
}

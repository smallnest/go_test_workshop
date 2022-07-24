package s2

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	t.Log("it is a log")
	t.Logf("it is a log at %v", time.Now().Format(time.RFC1123))

	t.Error("it is an error")
	t.Errorf("it is an error at %v", time.Now().Format(time.RFC1123))
}

package s2

import (
	"testing"
	"time"
)

func TestFail(t *testing.T) {
	t.Fail()
	t.Log("after Fail")
	t.FailNow()
	t.Log("after FailNow")
}

func TestFatal(t *testing.T) {
	t.Fatal("fataled")
	t.Log("after Fatal")
}

func TestFatalf(t *testing.T) {
	t.Fatalf("there is: %v", "Fatalf")
	t.Log("after Fatalf")
}

func TestFailNowInAnotherGoroutine(t *testing.T) {
	go func() {
		t.FailNow()
		t.Log("after FailNow in another goroutine")
	}()

	time.Sleep(time.Second)
	t.Log("after one second")
}

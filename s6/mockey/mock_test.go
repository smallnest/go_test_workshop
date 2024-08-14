package mockey

import (
	"fmt"
	"os"
	"testing"

	"github.com/bytedance/mockey"
)

// go test -v . -ldflags=-checklinkname=0 -gcflags="all=-N -l"

func TestCallerFunc(t *testing.T) {
	origin := fmt.Println
	mock := func(a ...any) (n int, err error) {
		fmt.Fprintln(os.Stdout, "I have changed the arguments")
		return origin(a...)
	}
	mock2 := mockey.Mock(origin).To(mock).Origin(&origin).Build()
	defer mock2.UnPatch()

	mock2.When(func(a ...any) bool {
		if len(a) == 1 && a[0] == "hello world" {
			return true
		}

		return false
	}).To(func(a ...any) (n int, err error) {
		return fmt.Fprintln(os.Stdout, "matched when")
	})

	fmt.Println("hello world")
	fmt.Println("hello mockey")
}

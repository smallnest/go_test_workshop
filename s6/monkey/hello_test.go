package monkey

import (
	"fmt"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey"
)

func TestPrintln(t *testing.T) {
	patches := gomonkey.ApplyFunc(fmt.Println, func(a ...any) (n int, err error) {

		return fmt.Fprintln(os.Stdout, "I have changed the arguments")
	})
	defer patches.Reset()

	fmt.Println("hello world")
}

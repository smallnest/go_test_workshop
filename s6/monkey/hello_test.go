package monkey

import (
	"fmt"
	"os"
	"testing"

	// "bou.ke/monkey"
	"github.com/agiledragon/gomonkey/v2"
)

func TestPrintlnByGomonkey(t *testing.T) {
	patches := gomonkey.ApplyFunc(fmt.Println, func(a ...any) (n int, err error) {

		return fmt.Fprintln(os.Stdout, "I have changed the arguments")
	})
	defer patches.Reset()

	fmt.Println("hello world")
}

// func TestPrintlnByMonkey(t *testing.T) {
// 	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
// 		s := make([]interface{}, len(a))
// 		for i, v := range a {
// 			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
// 		}
// 		return fmt.Fprintln(os.Stdout, s...)
// 	})
// 	fmt.Println("what the hell?") // what the *bleep*?
// }

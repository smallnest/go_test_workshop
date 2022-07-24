// package s1_test

package s1

import "testing"

func Test_greet(t *testing.T) {
	s := greet("test")

	if s != "hello test" {
		t.Errorf("want 'hello test' but got %s", s)
	}
}

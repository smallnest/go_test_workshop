package s2

import "testing"

func notHelper(t *testing.T, msg string) {
	t.Error(msg)
}

func helper(t *testing.T, msg string) {
	t.Helper()
	t.Error(msg)
}

func notHelperCallingHelper(t *testing.T, msg string) {
	helper(t, msg)
}

func helperCallingHelper(t *testing.T, msg string) {
	t.Helper()
	helper(t, msg)
}

func TestHelper(t *testing.T) {
	notHelper(t, "0")
	helper(t, "1")
	notHelperCallingHelper(t, "2")
	helperCallingHelper(t, "3")

	fn := func(msg string) {
		t.Helper()
		t.Error(msg)
	}
	fn("4")

	t.Helper()
	t.Error("5")

	t.Run("sub", func(t *testing.T) {
		helper(t, "6")
		notHelperCallingHelper(t, "7")
		t.Helper()
		t.Error("8")
	})
}

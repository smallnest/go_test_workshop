package s2

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	tempDir, err := ioutil.TempDir(".", "temp")
	if err != nil {
		t.Errorf("create tempDir: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(tempDir) })

	err = ioutil.WriteFile(filepath.Join(tempDir, "test.log"), []byte("hello test"), 0644)
	if err != nil {
		t.Errorf("want writing success but got %v", err)
	}
}

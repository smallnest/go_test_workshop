<<<<<<< HEAD
package s2

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("do stuff AFTER the tests!")

	os.Exit(exitVal)
}
package util

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateTestDB(t *testing.T) {
	db := CreateTestDB(t)
	defer db.Close()

	_, err := db.Exec("DELETE FROM tor_list")
	assert.NoError(t, err)
}

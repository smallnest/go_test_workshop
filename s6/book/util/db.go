package util

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func CreateTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "file:../testdata/test.db?cache=shared")
	assert.NoError(t, err)

	db.Exec(`
	CREATE TABLE IF NOT EXISTS books (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		isbn TEXT,
		subject TEXT,
		author_id INTEGER
	 );
	 DELETE FROM books;
	`)

	return db
}

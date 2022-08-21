package dao

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/smallnest/go_test_workshop/s7/book/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookStore_InsertBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "an error '%s' was not expected when opening a stub database connection", err)
	defer db.Close()

	store := &bookStore{
		db: db,
	}

	book := &model.Book{
		Title:    "The Go Programming Language",
		AuthorID: 1,
		ISBN:     "978-0134190440",
		Subject:  "computers",
	}

	mock.ExpectExec("INSERT INTO books").WillReturnResult(sqlmock.NewResult(1, 1))

	err = store.InsertBook(context.TODO(), book)

	require.NoError(t, err)
	assert.NotZero(t, book.ID)
}

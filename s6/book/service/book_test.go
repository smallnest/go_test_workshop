package service

import (
	"context"
	"testing"

	"github.com/smallnest/go_test_workshop/s6/book/model"
	"github.com/smallnest/go_test_workshop/s6/book/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookService_CreateBook(t *testing.T) {
	db := util.CreateTestDB(t)
	defer db.Close()

	bookService := NewBookService(db)

	book := &model.Book{
		Title:    "The Go Programming Language",
		AuthorID: 1,
		ISBN:     "978-0134190440",
		Subject:  "computers",
	}
	err := bookService.CreateBook(context.TODO(), book)

	require.NoError(t, err)
	assert.NotZero(t, book.ID)
}

package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/smallnest/go_test_workshop/s6/book/model"
	"github.com/smallnest/go_test_workshop/s6/book/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookController_GetBook_By_Monkey(t *testing.T) {
	bookService := service.NewBookService(nil)
	bc := &BookController{bookService: bookService}

	data := url.Values{}
	data.Set("title", "The Go Programming Language")
	data.Set("autherID", "1")
	data.Set("isbn", "978-0134190440")
	data.Set("subject", "computers")

	r := httptest.NewRequest("POST", "http://example.com/foo", strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	book := &model.Book{
		ID:       6,
		Title:    "The Go Programming Language",
		AuthorID: 1,
		ISBN:     "978-0134190440",
		Subject:  "computers",
	}
	patches := gomonkey.ApplyMethodFunc(bookService, "CreateBook", func(ctx context.Context, b *model.Book) error {
		b.ID = 6

		return nil
	})
	defer patches.Reset()

	w := httptest.NewRecorder()
	bc.CreateBook(w, r)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)

	err := json.Unmarshal(body, &book)
	require.NoError(t, err)
	assert.NotZero(t, book.ID)
}

package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/smallnest/go_test_workshop/s7/book/model"
	"github.com/smallnest/go_test_workshop/s7/book/service"
	"github.com/smallnest/go_test_workshop/s7/book/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookController_GetBook(t *testing.T) {
	db := util.CreateTestDB(t)
	defer db.Close()

	bookService := service.NewBookService(db)
	bc := &BookController{bookService: bookService}

	data := url.Values{}
	data.Set("title", "The Go Programming Language")
	data.Set("authorID", "1")
	data.Set("isbn", "978-0134190440")
	data.Set("subject", "computers")

	r := httptest.NewRequest("POST", "http://example.com/foo", strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	bc.CreateBook(w, r)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)

	var book model.Book
	err := json.Unmarshal(body, &book)
	require.NoError(t, err)
	assert.NotZero(t, book.ID)
}

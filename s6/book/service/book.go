package service

import (
	"context"
	"database/sql"

	"github.com/smallnest/go_test_workshop/s6/book/dao"
	"github.com/smallnest/go_test_workshop/s6/book/model"
)

type BookService interface {
	CreateBook(ctx context.Context, b *model.Book) error
}

func NewBookService(db *sql.DB) *bookService {
	return &bookService{
		store: dao.NewBookStore(db),
	}
}

type bookService struct {
	store dao.BookStore
}

func (s *bookService) CreateBook(ctx context.Context, b *model.Book) error {
	return s.store.InsertBook(ctx, b)
}

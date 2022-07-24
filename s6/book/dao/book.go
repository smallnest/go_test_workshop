package dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/smallnest/go_test_workshop/s6/book/model"
)

type BookStore interface {
	InsertBook(context.Context, *model.Book) error
}

func NewBookStore(db *sql.DB) *bookStore {
	return &bookStore{
		db: db,
	}
}

type bookStore struct {
	db *sql.DB
}

func (s *bookStore) InsertBook(ctx context.Context, b *model.Book) error {
	const stmt = "INSERT INTO books (author_id, title, isbn) VALUES ($1, $2, $3)"

	result, err := s.db.ExecContext(ctx, stmt, b.AuthorID, b.Title, b.ISBN)
	if err != nil {
		return fmt.Errorf("could not insert row: %w", err)
	}

	if _, err := result.RowsAffected(); err != nil {
		return fmt.Errorf("could not get affected rows: %w", err)
	}

	b.ID, _ = result.LastInsertId()

	return nil
}

package repository

import (
	"database/sql"
	"fmt"

	"github.com/shironxn/astra/internal/domain"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) domain.BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (r *BookRepository) Create(book domain.Book) error {
	q := "INSERT INTO books (title, author_id, genre, synopsis) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(q, book.Title, book.AuthorId, book.Genre, book.Synopsis)
	if err != nil {
		return fmt.Errorf("failed to insert book: %v", err)
	}
	return nil
}

func (r *BookRepository) GetAll() ([]domain.Book, error) {
	q := "SELECT * FROM books"
	rows, err := r.DB.Query(q)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %v", err)
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorId, &book.Genre, &book.Synopsis, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan book row: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over book rows: %v", err)
	}

	return books, nil
}

func (r *BookRepository) GetById(id int) (*domain.Book, error) {
	q := "SELECT * FROM books WHERE id = ?"
	row := r.DB.QueryRow(q, id)

	var book domain.Book
	if err := row.Scan(&book.ID, &book.Title, &book.AuthorId, &book.Genre, &book.Synopsis, &book.CreatedAt, &book.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to get book: %v", err)
	}

	return &book, nil
}

func (r *BookRepository) Update(id int, book domain.Book) error {
	q := "UPDATE books SET title = ?, genre = ?, synopsis = ? WHERE id = ?"
	_, err := r.DB.Exec(q, book.Title, book.Genre, book.Synopsis, id)
	if err != nil {
		return fmt.Errorf("failed to update book: %v", err)
	}
	return nil
}

func (r *BookRepository) Delete(id int) error {
	q := "DELETE FROM books WHERE id = ?"
	_, err := r.DB.Exec(q, id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %v", err)
	}
	return nil
}

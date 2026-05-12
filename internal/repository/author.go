package repository

import (
	"database/sql"
	"fmt"

	"github.com/shironxn/astra/internal/domain"
)

type AuthorRepository struct {
	DB *sql.DB
}

func NewAuthorRepository(db *sql.DB) domain.AuthorRepository {
	return &AuthorRepository{
		DB: db,
	}
}

func (r *AuthorRepository) Create(author domain.Author) error {
	q := "INSERT INTO authors (name, bio) VALUES (?, ?)"
	_, err := r.DB.Exec(q, author.Name, author.Bio)
	if err != nil {
		return fmt.Errorf("failed to insert author: %v", err)
	}
	return nil
}

func (r *AuthorRepository) GetAll() ([]domain.Author, error) {
	q := "SELECT * FROM authors"
	rows, err := r.DB.Query(q)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch authors: %v", err)
	}
	defer rows.Close()

	var authors []domain.Author
	for rows.Next() {
		var author domain.Author
		if err := rows.Scan(&author.ID, &author.Name, &author.Bio, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan author row: %v", err)
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over author rows: %v", err)
	}

	return authors, nil
}

func (r *AuthorRepository) GetById(id int) (*domain.Author, error) {
	q := "SELECT * FROM authors WHERE id = ?"
	row := r.DB.QueryRow(q, id)

	var author domain.Author
	if err := row.Scan(&author.ID, &author.Name, &author.Bio, &author.CreatedAt, &author.UpdatedAt); err != nil {
		return nil, fmt.Errorf("failed to get author: %v", err)
	}

	return &author, nil
}

func (r *AuthorRepository) Update(id int, author domain.Author) error {
	q := "UPDATE authors SET name = ?, bio = ? WHERE id = ?"
	_, err := r.DB.Exec(q, author.Name, author.Bio, id)
	if err != nil {
		return fmt.Errorf("failed to update author: %v", err)
	}
	return nil
}

func (r *AuthorRepository) Delete(id int) error {
	q := "DELETE FROM authors WHERE id = ?"
	_, err := r.DB.Exec(q, id)
	if err != nil {
		return fmt.Errorf("failed to delete author: %v", err)
	}
	return nil
}

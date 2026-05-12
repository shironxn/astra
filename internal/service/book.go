package service

import "github.com/shironxn/astra/internal/domain"

type BookService struct {
	repository domain.BookRepository
}

func NewBookService(repository domain.BookRepository) domain.BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) Create(book domain.Book) error {
	return s.repository.Create(book)
}

func (s *BookService) GetAll() ([]domain.Book, error) {
	return s.repository.GetAll()
}

func (s *BookService) GetById(id int) (*domain.Book, error) {
	return s.repository.GetById(id)
}

func (s *BookService) Update(id int, book domain.Book) error {
	return s.repository.Update(id, book)
}

func (s *BookService) Delete(id int) error {
	return s.repository.Delete(id)
}

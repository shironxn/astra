package service

import "github.com/shironxn/astra/internal/domain"

type AuthorService struct {
	repository domain.AuthorRepository
}

func NewAuthorService(repository domain.AuthorRepository) domain.AuthorService {
	return &AuthorService{
		repository: repository,
	}
}

func (s *AuthorService) Create(author domain.Author) error {
	return s.repository.Create(author)
}

func (s *AuthorService) GetAll() ([]domain.Author, error) {
	return s.repository.GetAll()
}

func (s *AuthorService) GetById(id int) (*domain.Author, error) {
	return s.repository.GetById(id)
}

func (s *AuthorService) Update(id int, author domain.Author) error {
	return s.repository.Update(id, author)
}

func (s *AuthorService) Delete(id int) error {
	return s.repository.Delete(id)
}

package usecase

import "fastapi-digital-library/go-backend/internal/domain"

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(r domain.BookRepository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) Create(book domain.Book) (domain.Book, error) {
	if err := book.Validate(); err != nil {
		return domain.Book{}, err
	}
	if err := s.repo.Create(book); err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (s *BookService) GetByID(id int) (domain.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) GetAll() ([]domain.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *BookService) Update(book domain.Book) (domain.Book, error) {
	if err := book.Validate(); err != nil {
		return domain.Book{}, err
	}
	if err := s.repo.Update(book); err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

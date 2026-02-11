package memory

import (
	"sync"

	"fastapi-digital-library/go-backend/internal/domain"
)

type BookRepository struct {
	mu    sync.RWMutex
	books []domain.Book
}

func NewBookRepository() *BookRepository {
	return &BookRepository{books: make([]domain.Book, 0)}
}

func (r *BookRepository) Create(book domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, b := range r.books {
		if b.ID == book.ID || b.ISBN == book.ISBN {
			return domain.ErrConflict
		}
	}
	r.books = append(r.books, book)
	return nil
}

func (r *BookRepository) GetByID(id int) (domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, b := range r.books {
		if b.ID == id {
			return b, nil
		}
	}
	return domain.Book{}, domain.ErrNotFound
}

func (r *BookRepository) GetAll() ([]domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]domain.Book, len(r.books))
	copy(out, r.books)
	return out, nil
}

func (r *BookRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, b := range r.books {
		if b.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return domain.ErrNotFound
}

func (r *BookRepository) Update(book domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, b := range r.books {
		if b.ID == book.ID {
			for _, other := range r.books {
				if other.ID != book.ID && (other.ID == book.ID || other.ISBN == book.ISBN) {
					return domain.ErrConflict
				}
			}
			r.books[i] = book
			return nil
		}
	}
	return domain.ErrNotFound
}

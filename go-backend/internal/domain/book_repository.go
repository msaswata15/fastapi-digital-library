package domain

type BookRepository interface {
	Create(book Book) error
	GetByID(id int) (Book, error)
	GetAll() ([]Book, error)
	Delete(id int) error
	Update(book Book) error
}

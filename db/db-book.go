package db

import "github.com/phongcao0114/book-service/model"

type BookDB interface {
	GetBookList() ([]model.Book, error)
	GetBookByID(id string) (model.Book, error)
	InsertBook(book model.Book) error
	UpdateBook(id string, book model.Book) error
	DeleteBook(id string) error
}

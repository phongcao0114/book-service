package service

import (
	"github.com/phongcao0114/book-service/db"
	"github.com/phongcao0114/book-service/model"
)

type BookManagerService interface {
	GetBookList() (interface{}, error)
	GetBookByID(id string) (interface{}, error)
	InsertBook(book model.Book) error
	UpdateBook(id string, book model.Book) error
	DeleteBook(id string) error
}

type BookManagerServiceImpl struct {
	*db.MasterDB
}

func NewBookManagerService(db *db.MasterDB) BookManagerService {
	return &BookManagerServiceImpl{db}
}

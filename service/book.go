package service

import (
	"github.com/google/uuid"
	"github.com/phongcao0114/book-service/model"
)

func (b BookManagerServiceImpl) GetBookList() (interface{}, error) {
	return b.BookDBImpl.GetBookList()
}

func (b BookManagerServiceImpl) GetBookByID(id string) (interface{}, error) {
	return b.BookDBImpl.GetBookByID(id)
}

func (b BookManagerServiceImpl) InsertBook(book model.Book) error {
	book.ID = uuid.New().String()
	return b.BookDBImpl.InsertBook(book)
}

func (b BookManagerServiceImpl) UpdateBook(id string, book model.Book) error {
	return b.BookDBImpl.UpdateBook(id, book)
}

func (b BookManagerServiceImpl) DeleteBook(id string) error {
	return b.BookDBImpl.DeleteBook(id)
}

package dbmysql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/phongcao0114/book-service/db"
	"github.com/phongcao0114/book-service/model"
)

type BookDBImpl struct {
	db *sql.DB
}

func initBookDBImpl() db.BookDB {
	db, err := sql.Open("mysql", "root:12345678x@X@/book")
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	instance := new(BookDBImpl)
	instance.db = db
	return instance
}

func (b BookDBImpl) GetBookList() ([]model.Book, error) {
	var bookList []model.Book
	stmt, err := b.db.Prepare("SELECT * FROM book")
	if err != nil {
		return []model.Book{}, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return []model.Book{}, err
	}
	defer stmt.Close()
	for rows.Next() {
		var book model.Book
		if err = rows.Scan(&book.ID, &book.Name, &book.Author); err != nil {
			return []model.Book{}, err
		}
		bookList = append(bookList, book)
	}
	return bookList, nil
}

func (b BookDBImpl) GetBookByID(id string) (model.Book, error) {
	stmt, err := b.db.Prepare("SELECT * FROM book WHERE id=?")
	if err != nil {
		return model.Book{}, err
	}
	defer stmt.Close()
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return model.Book{}, err
	}
	var book model.Book
	for rows.Next() {
		if err = rows.Scan(&book.ID, &book.Name, &book.Author); err != nil {
			return model.Book{}, err
		}
	}
	return book, nil
}

func (b BookDBImpl) InsertBook(book model.Book) error {
	stmt, err := b.db.Prepare("INSERT INTO book VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err = stmt.ExecContext(ctx, book.ID, book.Name, book.Author)
	if err != nil {
		return err
	}
	return nil
}

func (b BookDBImpl) UpdateBook(id string, book model.Book) error {
	stmt, err := b.db.Prepare("UPDATE book SET name=?, author=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err = stmt.ExecContext(ctx, book.Name, book.Author, id)
	if err != nil {
		return err
	}
	return nil
}

func (b BookDBImpl) DeleteBook(id string) error {
	stmt, err := b.db.Prepare("DELETE FROM book WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

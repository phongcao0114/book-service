package v1

import (
	"net/http"

	"github.com/phongcao0114/book-service/model"

	"github.com/gin-gonic/gin"
	"github.com/phongcao0114/book-service/service"
)

func GetBookList(c *gin.Context) {
	s := c.MustGet("bookManagerService").(service.BookManagerService)
	resp, err := s.GetBookList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func GetBookByID(c *gin.Context) {
	s := c.MustGet("bookManagerService").(service.BookManagerService)
	id := c.Param("id")
	resp, err := s.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func InsertBook(c *gin.Context) {
	s := c.MustGet("bookManagerService").(service.BookManagerService)
	var book model.Book
	c.BindJSON(&book)
	err := s.InsertBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.Status(http.StatusCreated)
	}
}

func UpdateBook(c *gin.Context) {
	s := c.MustGet("bookManagerService").(service.BookManagerService)
	id := c.Param("id")
	var book model.Book
	c.BindJSON(&book)
	err := s.UpdateBook(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.Status(http.StatusOK)
	}
}

func Delete(c *gin.Context) {
	s := c.MustGet("bookManagerService").(service.BookManagerService)
	id := c.Param("id")
	err := s.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.Status(http.StatusOK)
	}
}

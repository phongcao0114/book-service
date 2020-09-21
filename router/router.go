package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/phongcao0114/book-service/router/api/v1"
	"github.com/phongcao0114/book-service/service"
)

func ApiMiddleware(s service.BookManagerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("bookManagerService", s)
		c.Next()
	}
}

func InitRouter(s service.BookManagerService) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(ApiMiddleware(s))
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/books", v1.GetBookList)
	apiv1.GET("/book/:id", v1.GetBookByID)
	apiv1.POST("/book", v1.InsertBook)
	apiv1.PUT("/book/:id", v1.UpdateBook)
	apiv1.DELETE("/book/:id", v1.Delete)
	return r
}

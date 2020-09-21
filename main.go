package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/phongcao0114/book-service/dbmysql"
	"github.com/phongcao0114/book-service/router"
	"github.com/phongcao0114/book-service/service"
)

func main() {
	db := dbmysql.Init()
	service := service.NewBookManagerService(&db)
	routersInit := router.InitRouter(service)
	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}

	server.ListenAndServe()
}

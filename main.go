package main

import (
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: middleware,
	}
}

func main() {
	server := InitServer()

	err := server.ListenAndServe()

	helper.PanicIfError(err)
}

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/utils"
	"net/http"
)

var config, err = utils.LoadConfig()

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:" + config.AppPort,
		Handler: middleware,
	}
}

func main() {
	server := InitServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

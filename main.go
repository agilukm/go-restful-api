package main

import (
	"fmt"
	"go-restful-api/middleware"
	"go-restful-api/utils"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var config, err = utils.LoadConfig()

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:" + config.AppPort,
		Handler: middleware,
	}
}

func main() {
	if err != nil {
		panic("cannot load config")
	}

	fmt.Println(config)
	server := InitServer()

	server.ListenAndServe()

}

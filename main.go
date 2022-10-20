package main

import (
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/repository"
	"go-restful-api/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := services.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := httprouter.New()

	router.GET("/api/products/", productController.FindAll)
	router.GET("/api/products/:id", productController.FindById)
	router.POST("/api/products/", productController.Create)
	router.PATCH("/api/products/:id", productController.Update)
	router.DELETE("/api/products/:id", productController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	helper.PanicIfError(err)
}

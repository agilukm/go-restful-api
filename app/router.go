package app

import (
	"go-restful-api/controller"
	"go-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:id", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PATCH("/api/products/:id", productController.Update)
	router.DELETE("/api/products/:id", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

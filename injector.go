//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/services"
	"net/http"
)

var productSet = wire.NewSet(
	repository.NewProductRepository,
	services.NewProductService,
	controller.NewProductController,
)

func InitServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		productSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}

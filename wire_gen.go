// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/services"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitServer() *http.Server {
	workspaceRepository := repository.NewWorkspaceRepository()
	db := app.NewDB()
	validate := validator.New()
	workspaceService := services.NewWorkspaceService(workspaceRepository, db, validate)
	workspaceController := controller.NewWorkspaceController(workspaceService)
	router := app.NewRouter(workspaceController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var workspaceSet = wire.NewSet(repository.NewWorkspaceRepository, services.NewWorkspaceService, controller.NewWorkspaceController)

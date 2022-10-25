package app

import (
	"go-restful-api/controller"
	"go-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(workspaceController controller.WorkspaceController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/workspaces", workspaceController.FindAll)
	router.GET("/api/workspaces/:id", workspaceController.FindById)
	router.POST("/api/workspaces", workspaceController.Create)
	router.PATCH("/api/workspaces/:id", workspaceController.Update)
	router.DELETE("/api/workspaces/:id", workspaceController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

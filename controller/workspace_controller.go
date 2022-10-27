package controller

import (
	"go-restful-api/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewWorkspaceController(workspaceService services.WorkspaceService) WorkspaceController {
	return &WorkspaceControllerImpl{
		WorkspaceService: workspaceService,
	}
}

type WorkspaceController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GenerateToken(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Join(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	RemoveMember(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

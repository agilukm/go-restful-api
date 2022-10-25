package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-restful-api/helper"
	"go-restful-api/model/request"
	"go-restful-api/model/response"
	"go-restful-api/services"
	"net/http"
	"strconv"
)

type WorkspaceControllerImpl struct {
	WorkspaceService services.WorkspaceService
}

func (controller *WorkspaceControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	workspaceCreateRequest := request.WorkspaceCreateRequest{}

	var data interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	var decodedData = data.(map[string]interface{})
	userId, _ := strconv.ParseInt(fmt.Sprint(decodedData["user_id"]), 0, 64)
	decodedData["UserId"] = userId
	decodedData["user_id"] = userId

	b, err := json.Marshal(decodedData)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(b, &workspaceCreateRequest)

	helper.PanicIfError(err)

	workspaceResponse := controller.WorkspaceService.Create(r.Context(), workspaceCreateRequest)
	httpResponse := response.Response{
		Code:   201,
		Status: "OK",
		Data:   workspaceResponse,
	}

	helper.WriteToResponseBody(w, httpResponse, http.StatusCreated)
}

func (controller *WorkspaceControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	workspaceUpdateRequest := request.WorkspaceUpdateRequest{}
	helper.ReadFromRequestBody(r, &workspaceUpdateRequest)

	workspaceId := params.ByName("id")

	id, err := strconv.Atoi(workspaceId)

	helper.PanicIfError(err)

	workspaceResponse := controller.WorkspaceService.Update(r.Context(), workspaceUpdateRequest, id)
	httpResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   workspaceResponse,
	}

	helper.WriteToResponseBody(w, httpResponse, http.StatusOK)
}

func (controller *WorkspaceControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	workspaceId := params.ByName("id")

	id, err := strconv.Atoi(workspaceId)

	helper.PanicIfError(err)

	controller.WorkspaceService.Delete(r.Context(), id)
	httpResponse := response.Response{
		Code:   204,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, httpResponse, http.StatusOK)
}

func (controller *WorkspaceControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	workspaceId := params.ByName("id")

	id, err := strconv.Atoi(workspaceId)

	helper.PanicIfError(err)

	workspaceResponse := controller.WorkspaceService.FindById(r.Context(), id)
	httpResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   workspaceResponse,
	}

	helper.WriteToResponseBody(w, httpResponse, http.StatusOK)
}

func (controller *WorkspaceControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	workspaceResponse, _ := controller.WorkspaceService.Browse(r.Context(), r.URL.Query())
	httpResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   workspaceResponse,
	}

	helper.WriteToResponseBody(w, httpResponse, http.StatusOK)
}

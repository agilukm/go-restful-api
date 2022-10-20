package controller

import (
	"go-restful-api/helper"
	"go-restful-api/model/request"
	"go-restful-api/model/response"
	"go-restful-api/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productCreateRequest := request.ProductCreateRequest{}
	helper.ReadFromRequestBody(r, &productCreateRequest)

	productResponse := controller.ProductService.Create(r.Context(), productCreateRequest)
	httpResponse := response.Response{
		Code:   201,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, httpResponse)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productUpdateRequest := request.ProductUpdateRequest{}
	helper.ReadFromRequestBody(r, &productUpdateRequest)

	productId := params.ByName("id")

	id, err := strconv.Atoi(productId)

	helper.PanicIfError(err)

	productUpdateRequest.Id = id
	productResponse := controller.ProductService.Update(r.Context(), productUpdateRequest)
	httpResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, httpResponse)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("id")

	id, err := strconv.Atoi(productId)

	helper.PanicIfError(err)

	controller.ProductService.Delete(r.Context(), id)
	httpResponse := response.Response{
		Code:   204,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, httpResponse)
}

func (controller *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("id")

	id, err := strconv.Atoi(productId)

	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindById(r.Context(), id)
	httpResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, httpResponse)
}

func (controller *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productResponse := controller.ProductService.Browse(r.Context())
	httpResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, httpResponse)
}

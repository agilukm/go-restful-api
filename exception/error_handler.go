package exception

import (
	"go-restful-api/helper"
	"go-restful-api/model/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if validationErrors(w, r, err) {
		return
	}
	if notFoundError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")

		response := response.Response{
			Code:   http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE ENTITY",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, response, http.StatusUnprocessableEntity)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")

		response := response.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, response, http.StatusNotFound)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response := response.Response{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(w, response, http.StatusInternalServerError)
}

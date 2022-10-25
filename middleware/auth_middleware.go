package middleware

import (
	"context"
	"go-restful-api/helper"
	"go-restful-api/model/response"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

var UserInfo map[string]interface{}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("x-userinfo") != "" {
		userinfo := r.Header.Get("x-userinfo")
		userinfo = strings.ReplaceAll(userinfo, `\`, "")
		newContext := context.WithValue(r.Context(), "userinfo", userinfo)

		middleware.Handler.ServeHTTP(w, r.WithContext(newContext))
	} else if r.Header.Get("X-API-KEY") != "RAHASIA" {
		w.Header().Set("Content-Type", "application/json")

		response := response.Response{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, response, http.StatusUnauthorized)
	} else {
		newContext := context.WithValue(r.Context(), "key", "val")
		middleware.Handler.ServeHTTP(w, r.WithContext(newContext))
	}
}

func setUserHeader(userInfo string) {

}

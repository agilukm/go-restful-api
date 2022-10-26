package middleware

import (
	"context"
	"go-restful-api/helper"
	"go-restful-api/model/response"
	"go-restful-api/utils"
	"io"
	"log"
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
	} else if r.Header.Get("x-userinfo") == "" && r.Header.Get("Authorization") != "" {
		config, _ := utils.LoadConfig()

		client := &http.Client{}

		req, _ := http.NewRequest("GET", config.BaseURL+"/user/me", nil)
		req.Header.Add("Authorization", r.Header.Get("Authorization"))

		resp, _ := client.Do(req)

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			userinfo := string(bodyBytes)

			newContext := context.WithValue(r.Context(), "userinfo", userinfo)

			middleware.Handler.ServeHTTP(w, r.WithContext(newContext))
		}
	} else {
		w.Header().Set("Content-Type", "application/json")

		response := response.Response{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, response, http.StatusUnauthorized)
	}
}

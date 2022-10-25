package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/model/entity"
	"go-restful-api/repository"
	"go-restful-api/services"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sqlx.DB {
	db, err := sql.Open("mysql", "root:1@tcp(localhost:3306)/belajar_go")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sqlx.DB) http.Handler {
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := services.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(productController)

	return middleware.NewAuthMiddleware(router)
}

func truncateProduct(db *sqlx.DB) {
	db.Exec("TRUNCATE product")
}

func TestCreateProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget", "price" : 1000}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3030/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, float64(1000), responseBody["data"].(map[string]interface{})["price"])
}
func TestCreateProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "", "price" : 1000}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3030/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 422, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNPROCESSABLE ENTITY", responseBody["status"])
}
func TestFindProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateProduct(db)
	tx, _ := db.Begin()
	productRepository := repository.NewProductRepository()
	product := productRepository.Save(context.Background(), tx, entity.Product{
		Name:  "Oke",
		Price: 100,
	})
	tx.Commit()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "", "price" : 1000}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3030/api/products/"+strconv.Itoa(int(product.Id)), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, product.Id, int64(responseBody["data"].(map[string]interface{})["id"].(float64)))

}

func TestFindProductFailed(t *testing.T) {

}
func TestGetProductSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3030/api/products", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestDeleteProductFailed(t *testing.T) {

}
func TestDeleteProductSuccess(t *testing.T) {

}

func TestGetProductFailed(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}

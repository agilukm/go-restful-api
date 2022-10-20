package helper

import (
	"go-restful-api/model/entity"
	"go-restful-api/model/response"
)

func ToProductResponse(product entity.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: int64(product.Price),
	}
}

func ToProductResponses(products []entity.Product) []response.ProductResponse {
	var productResponses = []response.ProductResponse{}
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}

package services

import (
	"context"
	"go-restful-api/model/request"
	"go-restful-api/model/response"
)

type ProductService interface {
	Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) response.ProductResponse
	Browse(ctx context.Context) []response.ProductResponse
}

package services

import (
	"context"
	"database/sql"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
	"go-restful-api/model/request"
	"go-restful-api/model/response"
	"go-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service ProductServiceImpl) Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := entity.Product{
		Name:  request.Name,
		Price: int64(request.Price),
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = request.Name
	product.Price = int64(request.Price)

	service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := service.ProductRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service ProductServiceImpl) FindById(ctx context.Context, id int) response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Browse(ctx context.Context) []response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	helper.PanicIfError(err)

	return helper.ToProductResponses(products)
}

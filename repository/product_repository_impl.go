package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product {
	query := "insert into products(name, price) values (?, ?)"
	result, err := tx.ExecContext(ctx, query, product.Name, product.Price)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = id

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product {
	query := "update products set name=?, price=? where id = ?"

	_, err := tx.ExecContext(ctx, query, product.Name, product.Price, product.Id)

	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product entity.Product) {
	query := "delete from products where id = ?"

	_, err := tx.ExecContext(ctx, query, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Product, error) {
	query := "select id, name, price from products where id = ?"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)
	defer rows.Close()

	product := entity.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}

}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Product {
	query := "select id, name, price from products"

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var products []entity.Product

	for rows.Next() {
		product := entity.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	defer rows.Close()

	return products
}

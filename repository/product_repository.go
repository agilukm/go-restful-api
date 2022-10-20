package repository

import (
	"context"
	"database/sql"
	"go-restful-api/model/entity"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product
	Update(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product
	Delete(ctx context.Context, tx *sql.Tx, product entity.Product)
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Product
}

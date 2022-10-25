package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-restful-api/model/entity"
	"net/url"
)

type WorkspaceRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, workspace entity.Workspace) entity.Workspace
	Update(ctx context.Context, tx *sqlx.Tx, workspace entity.Workspace) entity.Workspace
	Delete(ctx context.Context, tx *sqlx.Tx, workspace entity.Workspace)
	FindById(ctx context.Context, tx *sqlx.Tx, id int) (entity.Workspace, error)
	FindAll(ctx context.Context, tx *sqlx.Tx, values url.Values) ([]entity.Workspace, interface{})
}

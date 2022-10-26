package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
	"go-restful-api/utils"
	"net/url"
)

type WorkspaceRepositoryImpl struct {
}

func NewWorkspaceRepository() WorkspaceRepository {
	return &WorkspaceRepositoryImpl{}
}

func (repository *WorkspaceRepositoryImpl) Save(ctx context.Context, tx *sqlx.Tx, Workspace entity.Workspace) entity.Workspace {
	query := "insert into workspaces(name, user_id, token, token_expired_at) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, Workspace.Name, Workspace.UserId, Workspace.Token, Workspace.TokenExpiredAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	Workspace.Id = id

	return Workspace
}

func (repository *WorkspaceRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, Workspace entity.Workspace) entity.Workspace {
	query := "update workspaces set name=?, user_id=? where id = ?"

	_, err := tx.ExecContext(ctx, query, Workspace.Name, Workspace.UserId, Workspace.Id)

	helper.PanicIfError(err)
	return Workspace
}

func (repository *WorkspaceRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, Workspace entity.Workspace) {
	query := "delete from workspaces where id = ?"

	_, err := tx.ExecContext(ctx, query, Workspace.Id)
	helper.PanicIfError(err)
}

func (repository *WorkspaceRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, id int) (entity.Workspace, error) {
	query := "select id, name, user_id from workspaces where id = ?"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)
	defer rows.Close()

	Workspace := entity.Workspace{}
	if rows.Next() {
		err := rows.Scan(&Workspace.Id, &Workspace.Name, &Workspace.UserId)
		helper.PanicIfError(err)
		return Workspace, nil
	} else {
		return Workspace, errors.New("Workspace not found")
	}

}

func (repository *WorkspaceRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx, values url.Values) ([]entity.Workspace, interface{}) {
	query := "select id, name, user_id, token, token_expired_at from workspaces"
	countQuery := "select count(id) from workspaces"

	userinfo := utils.GetUserinfo(ctx)

	level := userinfo.Level

	switch level {
	case "LECTURER":
		query = fmt.Sprintf("%s where user_id = %s", query, userinfo.Id)
		countQuery = fmt.Sprintf("%s where user_id = %s", countQuery, userinfo.Id)
	default:
	}

	var workspaces []entity.Workspace
	newStruct := new(entity.WorkspaceFilterable)

	query = helper.Filter(values, newStruct, query)
	query = helper.Sort(values, newStruct, query)

	countQuery = helper.Filter(values, newStruct, countQuery)
	countQuery = helper.Sort(values, newStruct, countQuery)

	query, _ = helper.Pager(values, query)

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	countRows := 0
	for rows.Next() {
		Workspace := entity.Workspace{}
		err := rows.Scan(&Workspace.Id, &Workspace.Name, &Workspace.UserId, &Workspace.Token, &Workspace.TokenExpiredAt)
		helper.PanicIfError(err)
		workspaces = append(workspaces, Workspace)
		countRows++
	}

	var total int

	tx.QueryRow(countQuery).Scan(&total)

	defer rows.Close()

	return workspaces, map[string]int{
		"total": total,
		"count": countRows,
	}
}

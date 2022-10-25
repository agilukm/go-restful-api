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
	"reflect"
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
	//limit, _ := strconv.Atoi(params.ByName("per_page"))

	userinfo := utils.GetUserinfo(ctx)
	fmt.Println(userinfo.Id)
	fmt.Println(userinfo.Name)
	fmt.Println(userinfo.Email)

	var workspaces []entity.Workspace
	newStruct := new(entity.WorkspaceFilterable)

	query = utils.Filter(values, newStruct, query)

	query = fmt.Sprintf("%s order by ID desc", query)

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	for rows.Next() {
		Workspace := entity.Workspace{}
		err := rows.Scan(&Workspace.Id, &Workspace.Name, &Workspace.UserId, &Workspace.Token, &Workspace.TokenExpiredAt)
		helper.PanicIfError(err)
		workspaces = append(workspaces, Workspace)
	}

	countQuery := "select count(id) from workspaces"
	var total int

	tx.QueryRow(countQuery).Scan(&total)

	defer rows.Close()

	return workspaces, map[string]int{
		"total_record": total,
		"total_page":   2,
	}
}

func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}
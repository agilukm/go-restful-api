package services

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
	"go-restful-api/model/request"
	"go-restful-api/repository"
	"net/url"
)

type WorkspaceServiceImpl struct {
	WorkspaceRepository repository.WorkspaceRepository
	DB                  *sqlx.DB
	Validate            *validator.Validate
}

func NewWorkspaceService(workspaceRepository repository.WorkspaceRepository, db *sqlx.DB, validate *validator.Validate) WorkspaceService {
	return &WorkspaceServiceImpl{
		WorkspaceRepository: workspaceRepository,
		DB:                  db,
		Validate:            validate,
	}
}

func (service WorkspaceServiceImpl) Create(ctx context.Context, request request.WorkspaceCreateRequest) entity.Workspace {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	workspace := entity.Workspace{
		Name:           request.Name,
		UserId:         int64(request.UserId),
		Token:          request.Token,
		TokenExpiredAt: request.TokenExpiredAt,
	}

	workspace = service.WorkspaceRepository.Save(ctx, tx, workspace)

	return workspace
}

func (service WorkspaceServiceImpl) Update(ctx context.Context, request request.WorkspaceUpdateRequest, id int) entity.Workspace {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	workspace, err := service.WorkspaceRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	workspace.Name = request.Name
	workspace.UserId = int64(request.UserId)

	service.WorkspaceRepository.Update(ctx, tx, workspace)

	return workspace
}

func (service WorkspaceServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	workspace, err := service.WorkspaceRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.WorkspaceRepository.Delete(ctx, tx, workspace)
}

func (service WorkspaceServiceImpl) FindById(ctx context.Context, id int) entity.Workspace {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	workspace, err := service.WorkspaceRepository.FindById(ctx, tx, id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return workspace
}

func (service WorkspaceServiceImpl) Browse(ctx context.Context, values url.Values) ([]entity.Workspace, interface{}) {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	workspaces, paginator := service.WorkspaceRepository.FindAll(ctx, tx, values)

	helper.PanicIfError(err)

	return workspaces, paginator
}

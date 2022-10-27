package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
	"go-restful-api/model/request"
	"go-restful-api/repository"
	"go-restful-api/utils"
	"gopkg.in/guregu/null.v3"
	"net/url"
	"strconv"
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
	userinfo := utils.GetUserinfo(ctx)
	UserId, _ := strconv.ParseInt(userinfo.Id, 0, 64)
	workspace := entity.Workspace{
		Name:           request.Name,
		UserId:         UserId,
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

	return workspaces, paginator
}

func (service WorkspaceServiceImpl) GenerateToken(ctx context.Context, request request.GenerateTokenRequest, id int) entity.Workspace {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	workspace, err := service.WorkspaceRepository.FindById(ctx, tx, id)

	expired_at := sql.NullString{
		String: request.TokenExpiredAt,
		Valid:  true,
	}

	workspace.TokenExpiredAt = null.String{expired_at}

	workspace.Token = null.NewString(utils.GenerateString(5), true)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.WorkspaceRepository.GenerateToken(ctx, tx, workspace)

	return workspace
}

func (service WorkspaceServiceImpl) Join(ctx context.Context, request request.JoinWorkspaceRequest) entity.WorkspaceMember {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	userinfo := utils.GetUserinfo(ctx)

	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	userId, err := strconv.Atoi(userinfo.Id)
	helper.PanicIfError(err)

	member := entity.WorkspaceMember{
		WorkspaceId: request.WorkspaceId,
		UserId:      int64(userId),
		Type:        request.Type,
	}
	repository.FindWithToken(ctx, tx, member, request.Token)
	member = service.WorkspaceRepository.Join(ctx, tx, member, request.Token)

	return member
}

func (service WorkspaceServiceImpl) RemoveMember(ctx context.Context, request request.RemoveWorkspaceRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.WorkspaceRepository.FindById(ctx, tx, int(request.WorkspaceId))

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	member := entity.WorkspaceMember{
		WorkspaceId: request.WorkspaceId,
		UserId:      request.UserId,
	}

	service.WorkspaceRepository.RemoveMember(ctx, tx, member)
}

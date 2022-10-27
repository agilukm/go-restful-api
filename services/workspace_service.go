package services

import (
	"context"
	"go-restful-api/model/entity"
	"go-restful-api/model/request"
	"net/url"
)

type WorkspaceService interface {
	Create(ctx context.Context, request request.WorkspaceCreateRequest) entity.Workspace
	Update(ctx context.Context, request request.WorkspaceUpdateRequest, id int) entity.Workspace
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) entity.Workspace
	Browse(ctx context.Context, values url.Values) ([]entity.Workspace, interface{})
	GenerateToken(ctx context.Context, request request.GenerateTokenRequest, id int) entity.Workspace
	Join(ctx context.Context, request request.JoinWorkspaceRequest) entity.WorkspaceMember
	RemoveMember(ctx context.Context, request request.RemoveWorkspaceRequest)
}

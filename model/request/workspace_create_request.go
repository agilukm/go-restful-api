package request

import (
	"gopkg.in/guregu/null.v3"
)

type WorkspaceCreateRequest struct {
	Name           string      `validate:"required,min=1,max=200"`
	UserId         int64       `json:"user_id"`
	Token          null.String `validate:"min=0"`
	TokenExpiredAt null.String `json:"token_expired_at"`
}

type GenerateTokenRequest struct {
	TokenExpiredAt string `json:"token_expired_at" validate:"required"`
}

type JoinWorkspaceRequest struct {
	WorkspaceId int64  `json:"workspace_id" validate:"required"`
	Token       string `json:"token" validate:"required"`
	Type        string `json:"type" validate:"required"`
}

type RemoveWorkspaceRequest struct {
	WorkspaceId int64 `json:"workspace_id" validate:"required"`
	UserId      int64 `json:"user_id" validate:"required"`
}

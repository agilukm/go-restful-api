package request

import (
	"gopkg.in/guregu/null.v3"
)

type WorkspaceCreateRequest struct {
	Name           string      `validate:"required,min=1,max=200"`
	UserId         int64       `validate:"required,min=0" json:"user_id"`
	Token          null.String `validate:"min=0"`
	TokenExpiredAt null.String `json:"token_expired_at"`
}

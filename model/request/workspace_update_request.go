package request

import (
	"gopkg.in/guregu/null.v3"
)

type WorkspaceUpdateRequest struct {
	Name           string      `validate:"required,min=1,max=200"`
	UserId         int64       `validate:"required,min=0,int"`
	Token          null.String `validate:"required,min=0"`
	TokenExpiredAt null.String `validate:"required"`
}

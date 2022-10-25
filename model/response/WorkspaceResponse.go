package response

import (
	"gopkg.in/guregu/null.v3"
)

type WorkspaceResponse struct {
	Id             int64       `json:"id"`
	Name           string      `json:"name"`
	UserId         int64       `json:"user_id"`
	Token          null.String `json:"token"`
	TokenExpiredAt null.String `json:"token_expired_at"`
}

package entity

import "gopkg.in/guregu/null.v3"

type Workspace struct {
	Id             int64       `json:"id"`
	Name           string      `json:"name"`
	UserId         int64       `json:"user_id"`
	Token          null.String `json:"token"`
	TokenExpiredAt null.String `json:"token_expired_at"`
}

type WorkspaceFilterable struct {
	id               string
	name             string
	user_id          string
	token            string
	token_expired_at string
}

package entity

import "gopkg.in/guregu/null.v3"

type Workspace struct {
	Id             int64       `json:"id" db:"id"`
	Name           string      `json:"name" db:"name"`
	UserId         int64       `json:"user_id" db:"user_id"`
	Token          null.String `json:"token"`
	TokenExpiredAt null.String `json:"token_expired_at" db:"token_expired_at"`
	CreatedAt      null.String `json:"created_at" db:"created_at"`
	UpdatedAt      null.String `json:"updated_at" db:"updated_at"`
	DeletedAt      null.String `json:"deleted_at" db:"deleted_at"`
}

type WorkspaceFilterable struct {
	id               string
	name             string
	user_id          string
	token            string
	token_expired_at string
}

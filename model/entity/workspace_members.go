package entity

type WorkspaceMember struct {
	Id          int64  `json:"id" db:"id"`
	WorkspaceId int64  `json:"workspace_id" db:"workspace_id"`
	UserId      int64  `json:"user_id" db:"user_id"`
	Type        string `json:"type" db:"type"`
}

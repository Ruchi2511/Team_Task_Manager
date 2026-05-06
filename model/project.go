package model

import "time"

type CreateProjectRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description"`
}
type Project struct {
	ID          string    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
type AddProjectMemberRequest struct {
	UserID string `json:"user_id" validate:"required,uuid"`
}

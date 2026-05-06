package model

import "time"

type CreateTaskRequest struct {
	ProjectID   string `json:"project_id" validate:"required,uuid"`
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description"`
	AssignedTo  string `json:"assigned_to" validate:"required,uuid"`
	Priority    string `json:"priority" validate:"required"`
	DueDate     string `json:"due_date" validate:"required"`
}
type Task struct {
	ID          string    `db:"id" json:"id"`
	ProjectID   string    `db:"project_id" json:"project_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	AssignedTo  string    `db:"assigned_to" json:"assigned_to"`
	AssignedBy  string    `db:"assigned_by" json:"assigned_by"`
	Status      string    `db:"status" json:"status"`
	Priority    string    `db:"priority" json:"priority"`
	DueDate     time.Time `db:"due_date" json:"due_date"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
type UpdateTaskStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

package model

type DashboardStats struct {
	TotalProjects   int `db:"total_projects" json:"total_projects"`
	TotalTasks      int `db:"total_tasks" json:"total_tasks"`
	CompletedTasks  int `db:"completed_tasks" json:"completed_tasks"`
	InProgressTasks int `db:"in_progress_tasks" json:"in_progress_tasks"`
	OverdueTasks    int `db:"overdue_tasks" json:"overdue_tasks"`
}

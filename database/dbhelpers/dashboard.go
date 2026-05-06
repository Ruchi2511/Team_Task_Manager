package dbhelpers

import (
	"Team_Task_Manager/database"
	"Team_Task_Manager/model"
)

func GetDashboardStats(role, userID string) (model.DashboardStats, error) {

	query := `
		SELECT
			COUNT(DISTINCT p.id) AS total_projects,
			COUNT(t.id) AS total_tasks,

			COUNT(t.id) FILTER (
				WHERE t.status = 'completed'
			) AS completed_tasks,

			COUNT(t.id) FILTER (
				WHERE t.status = 'in_progress'
			) AS in_progress_tasks,

			COUNT(t.id) FILTER (
				WHERE t.due_date < NOW()
				AND t.status != 'completed'
			) AS overdue_tasks

		FROM projects p

		LEFT JOIN tasks t
			ON t.project_id = p.id
			AND t.archived_at IS NULL

		WHERE p.archived_at IS NULL
	`

	if role != "admin" {
		query += `
			AND p.id IN (
				SELECT project_id
				FROM project_members
				WHERE user_id = $1
			)
		`
	}
	var stats model.DashboardStats
	var err error
	if role != "admin" {
		err = database.Asset.Get(&stats, query, userID)
	} else {
		err = database.Asset.Get(&stats, query)
	}
	if err != nil {
		return stats, err
	}
	return stats, nil
}

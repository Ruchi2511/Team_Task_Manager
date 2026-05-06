package dbhelpers

import (
	"Team_Task_Manager/database"
	"Team_Task_Manager/model"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

func CreateTask(tx *sqlx.Tx, projectID, title, description, assignedTo, assignedBy, priority string, dueDate time.Time) (string, error) {

	isProjectExist, err := IsProjectExist(projectID)
	if err != nil {
		return "", err
	}

	if !isProjectExist {
		return "", errors.New("project not found")
	}

	isUserExist, err := IsUserActive(assignedTo)
	if err != nil {
		return "", err
	}

	if !isUserExist {
		return "", errors.New("assigned user not found")
	}

	isMember, err := IsProjectMember(projectID, assignedTo)
	if err != nil {
		return "", err
	}

	if !isMember {
		return "", errors.New("user is not a member of this project")
	}

	query := `
		INSERT INTO tasks (
			project_id,
			title,
			description,
			assigned_to,
			assigned_by,
			priority,
			due_date
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
		RETURNING id
	`

	var taskID string

	err = tx.Get(&taskID, query, projectID, title, description, assignedTo, assignedBy, priority, dueDate)
	if err != nil {
		return "", err
	}

	return taskID, nil
}

func IsProjectMember(projectID, userID string) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM project_members
			WHERE project_id = $1
			AND user_id = $2
		)
	`

	var exists bool

	err := database.Asset.Get(&exists, query, projectID, userID)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func GetTasks(projectID, status, priority, role, userID string, limit, offset int) ([]model.Task, error) {

	query := `
		SELECT
			t.id,
			t.project_id,
			t.title,
			t.description,
			t.assigned_to,
			t.assigned_by,
			t.status,
			t.priority,
			t.due_date,
			t.created_at
		FROM tasks t
		WHERE t.archived_at IS NULL
		AND ($1 = '' OR t.project_id::text = $1)
		AND ($2 = '' OR t.status::text ILIKE '%'||$2||'%')
		AND ($3 = '' OR t.priority::text ILIKE '%'||$3||'%')
	`
	if role != "admin" {
		query += `
			AND t.assigned_to = $4
			ORDER BY t.created_at DESC
			LIMIT $5 OFFSET $6
		`
	} else {
		query += `
			ORDER BY t.created_at DESC
			LIMIT $4 OFFSET $5
		`
	}
	tasks := make([]model.Task, 0)
	var err error
	if role != "admin" {
		err = database.Asset.Select(&tasks, query, projectID, status, priority, userID, limit, offset)
	} else {
		err = database.Asset.Select(&tasks, query, projectID, status, priority, limit, offset)
	}
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func UpdateTaskStatus(taskID, status string) error {

	query := `
		UPDATE tasks
		SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2
		AND archived_at IS NULL
	`
	result, err := database.Asset.Exec(query, status, taskID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil
}
func IsTaskDueDateCrossed(taskID string) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM tasks
			WHERE id = $1
			AND due_date < NOW()
			AND archived_at IS NULL
		)
	`
	var exists bool
	err := database.Asset.Get(&exists, query, taskID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

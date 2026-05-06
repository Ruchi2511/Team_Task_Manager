package dbhelpers

import (
	"Team_Task_Manager/database"
	"Team_Task_Manager/model"
	"errors"

	"github.com/jmoiron/sqlx"
)

func CreateProject(tx *sqlx.Tx, title string, description string, createdBy string) (string, error) {

	isUserExist, err := IsUserActive(createdBy)

	if err != nil {
		return "", err
	}

	if !isUserExist {
		return "", errors.New("user not found")
	}

	query := `
		INSERT INTO projects (
			title,
			description,
			created_by
		)
		VALUES (
			$1,
			$2,
			$3
		)
		RETURNING id
	`

	var projectID string

	err = tx.Get(&projectID, query, title, description, createdBy)

	if err != nil {
		return "", err
	}

	return projectID, nil
}

func GetProjects(title, role, userID string, limit, offset int) ([]model.Project, error) {

	query := `
		SELECT
			p.id,
			p.title,
			p.description,
			p.created_by,
			p.created_at
		FROM projects p
		WHERE p.archived_at IS NULL
		AND ($1 = '' OR p.title ILIKE '%'||$1||'%')
	`

	if role != "admin" {
		query += `
			AND p.id IN (
				SELECT project_id
				FROM project_members
				WHERE user_id = $2
				AND archived_at IS NULL
			)
			ORDER BY p.created_at DESC
			LIMIT $3 OFFSET $4
		`
	} else {
		query += `
			ORDER BY p.created_at DESC
			LIMIT $2 OFFSET $3
		`
	}

	projects := make([]model.Project, 0)

	var err error

	if role != "admin" {
		err = database.Asset.Select(
			&projects,
			query,
			title,
			userID,
			limit,
			offset,
		)
	} else {
		err = database.Asset.Select(
			&projects,
			query,
			title,
			limit,
			offset,
		)
	}
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func AddProjectMember(tx *sqlx.Tx, projectID string, userID string) error {
	isProjectExist, err := IsProjectExist(projectID)
	if err != nil {
		return err
	}
	if !isProjectExist {
		return errors.New("project not found")
	}
	isUserExist, err := IsUserActive(userID)
	if err != nil {
		return err
	}
	if !isUserExist {
		return errors.New("user not found")
	}
	query := `
		INSERT INTO project_members (
			project_id,
			user_id
		)
		VALUES (
			$1,
			$2
		)
	`
	_, err = tx.Exec(query, projectID, userID)
	if err != nil {
		return err
	}
	return nil
}

func IsProjectExist(projectID string) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM projects
			WHERE id = $1
			AND archived_at IS NULL
		)
	`
	var exists bool
	err := database.Asset.Get(&exists, query, projectID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func IsUserActive(userID string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE id = $1
			AND archived_at IS NULL
		)
	`
	var exists bool
	err := database.Asset.Get(&exists, query, userID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

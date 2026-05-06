package dbhelpers

import (
	"Team_Task_Manager/database"
	"Team_Task_Manager/model"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func GetUserIDFromSession(sessionID string) (string, error) {

	query := `
		SELECT user_id
		FROM user_sessions
		WHERE id = $1
		AND archived_at IS NULL
		AND expires_at > NOW()
	`
	var userID string
	err := database.Asset.Get(&userID, query, sessionID)

	if err != nil {
		return "", err
	}

	return userID, nil
}
func IsUserExist(email string) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE email = $1
			AND archived_at IS NULL
		)
	`
	var exists bool
	err := database.Asset.Get(&exists, query, email)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func CreateUser(tx *sqlx.Tx, name string, email string, password string, role string) (string, error) {

	query := `
		INSERT INTO users (
			name,
			email,
			password,
			role
		)
		VALUES (
			$1,
			TRIM(LOWER($2)),
			$3,
			$4
		)
		RETURNING id
	`

	var userID string

	err := tx.Get(&userID, query, name, email, password, role)
	if err != nil {
		return "", err
	}
	return userID, nil
}
func CreateUserSession(tx *sqlx.Tx, userID string) (string, error) {

	query := `
		INSERT INTO user_sessions (
			user_id,
			refresh_token,
			expires_at
		)
		VALUES (
			$1,
			gen_random_uuid()::text,
			NOW() + INTERVAL '7 days'
		)
		RETURNING id
	`
	var sessionID string
	err := tx.Get(&sessionID, query, userID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}
func GetUserByEmail(tx *sqlx.Tx, email, password string) (string, string, error) {

	query := `
		SELECT 
			id,
			password,
			role
		FROM users 
		WHERE TRIM(LOWER(email)) = LOWER($1)
		AND archived_at IS NULL
	`

	var result model.UserExist

	err := tx.Get(&result, query, email)

	if err != nil {
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(result.Password),
		[]byte(password),
	); err != nil {
		return "", "", err
	}

	return result.ID, result.Role, nil
}
func ArchiveUserSession(sessionID string) error {
	query := `
		UPDATE user_sessions
		SET archived_at = NOW()
		WHERE id = $1
		AND archived_at IS NULL
	`
	_, err := database.Asset.Exec(query, sessionID)
	if err != nil {
		return err
	}
	return nil
}

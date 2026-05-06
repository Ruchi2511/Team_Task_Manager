package handlers

import (
	"Team_Task_Manager/database"
	"Team_Task_Manager/database/dbhelpers"
	"Team_Task_Manager/middleware"
	"Team_Task_Manager/model"
	"Team_Task_Manager/utils"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

var validate = validator.New()

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var body model.RegisterUserRequest
	var userID string
	var sessionID string
	defaultRole := "member"
	if err := utils.ParseBody(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "failed to parse request body", err)
		return
	}
	err := validate.Struct(&body)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "validation failed", err)
		return
	}

	isUserExist, err := dbhelpers.IsUserExist(body.Email)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to check existing user", err)

		return
	}

	if isUserExist {
		utils.RespondError(w, http.StatusConflict, "user already exists", errors.New("email already registered"))

		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)

	if err != nil {

		utils.RespondError(w, http.StatusInternalServerError, "failed to hash password", err)
		return
	}
	txErr := database.Tx(func(tx *sqlx.Tx) error {

		userID, err = dbhelpers.CreateUser(
			tx,
			body.Name,
			body.Email,
			hashedPassword,
			defaultRole,
		)
		if err != nil {
			return err
		}
		sessionID, err = dbhelpers.CreateUserSession(tx, userID)
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to create user", txErr)
		return
	}
	token, err := utils.GenerateJWT(
		userID,
		sessionID,
		defaultRole,
	)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to generate token", err)
		return
	}
	utils.RespondJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "user registered successfully",
		"data": map[string]interface{}{
			"user_id": userID,
			"role":    defaultRole,
		},
		"access_token": token,
	},
	)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {

	var body model.LoginRequest
	var userID string
	var userRole string
	var sessionID string
	var userName string

	if err := utils.ParseBody(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "failed to parse request body", err)
		return
	}
	err := validate.Struct(&body)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "validation failed", err)
		return
	}
	txErr := database.Tx(func(tx *sqlx.Tx) error {
		var err error
		userID, userRole, userName, err = dbhelpers.GetUserByEmail(
			tx,
			body.Email,
			body.Password,
		)
		if err != nil {
			return err
		}

		sessionID, err = dbhelpers.CreateUserSession(tx, userID)
		if err != nil {
			return err
		}

		return nil
	})

	if txErr != nil {
		utils.RespondError(w, http.StatusUnauthorized, "invalid email or password", txErr)
		return
	}

	token, err := utils.GenerateJWT(userID, sessionID, userRole)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to generate token", err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "login successful",
		"data": map[string]interface{}{
			"user_id": userID,
			"role":    userRole,
			"name":    userName,
		},
		"access_token": token,
	})
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	auth, ok := middleware.GetAuthContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	err := dbhelpers.ArchiveUserSession(auth.SessionID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to logout user", err)
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "logout successful",
	})
}
func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := dbhelpers.GetUsers()

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to fetch users", err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

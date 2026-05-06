package handlers

import (
	"Team_Task_Manager/database"
	"Team_Task_Manager/database/dbhelpers"
	"Team_Task_Manager/middleware"
	"Team_Task_Manager/model"
	"Team_Task_Manager/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	auth, ok := middleware.GetAuthContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	var body model.CreateTaskRequest
	if err := utils.ParseBody(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "failed to parse request body", err)
		return
	}
	err := validate.Struct(&body)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "validation failed", err)
		return
	}
	dueDate, err := time.Parse("2006-01-02", body.DueDate)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid due_date format", err)
		return
	}
	var taskID string
	txErr := database.Tx(func(tx *sqlx.Tx) error {
		var err error

		taskID, err = dbhelpers.CreateTask(
			tx,
			body.ProjectID,
			body.Title,
			body.Description,
			body.AssignedTo,
			auth.UserID,
			body.Priority,
			dueDate,
		)

		if err != nil {
			return err
		}

		return nil
	})

	if txErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to create task", txErr)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "task created successfully",
		"task_id": taskID,
	})
}
func GetTasks(w http.ResponseWriter, r *http.Request) {
	auth, ok := middleware.GetAuthContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	projectFilter := r.URL.Query().Get("project_id")
	statusFilter := r.URL.Query().Get("status")
	priorityFilter := r.URL.Query().Get("priority")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page := 1
	limit := 5
	if pageStr != "" {
		pageValue, err := strconv.Atoi(pageStr)
		if err != nil || pageValue <= 0 {
			utils.RespondError(w, http.StatusBadRequest, "invalid page", err)
			return
		}
		page = pageValue
	}
	if limitStr != "" {
		limitValue, err := strconv.Atoi(limitStr)

		if err != nil || limitValue <= 0 {
			utils.RespondError(w, http.StatusBadRequest, "invalid limit", err)
			return
		}
		limit = limitValue
	}
	offset := (page - 1) * limit
	tasks, err := dbhelpers.GetTasks(
		projectFilter,
		statusFilter,
		priorityFilter,
		auth.Role,
		auth.UserID,
		limit,
		offset,
	)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to fetch tasks", err)
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}
func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {

	taskID := chi.URLParam(r, "id")

	var body model.UpdateTaskStatusRequest

	if err := utils.ParseBody(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "failed to parse request body", err)
		return
	}

	err := validate.Struct(&body)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "validation failed", err)
		return
	}

	if body.Status == "completed" {
		isDueDateCrossed, err := dbhelpers.IsTaskDueDateCrossed(taskID)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "failed to validate due date", err)
			return
		}
		if isDueDateCrossed {
			utils.RespondError(w, http.StatusBadRequest, "cannot complete overdue task", nil)
			return
		}
	}
	err = dbhelpers.UpdateTaskStatus(taskID, body.Status)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to update task status", err)
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "task status updated successfully",
	})
}

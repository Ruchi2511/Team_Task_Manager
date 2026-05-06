package handlers

import (
	"Team_Task_Manager/database/dbhelpers"
	"Team_Task_Manager/middleware"
	"Team_Task_Manager/utils"
	"net/http"
)

func GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	auth, ok := middleware.GetAuthContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}
	stats, err := dbhelpers.GetDashboardStats(auth.Role, auth.UserID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to fetch dashboard stats", err)
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"stats": stats,
	})
}

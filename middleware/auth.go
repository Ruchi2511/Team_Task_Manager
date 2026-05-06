package middleware

import (
	"Team_Task_Manager/database/dbhelpers"
	"Team_Task_Manager/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/form3tech-oss/jwt-go"
)

type AuthContext struct {
	UserID    string
	SessionID string
	Role      string
}

type contextKey string

const authContextKey contextKey = "authContext"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))

		if authHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, "authorization header required", nil)
			return
		}

		parts := strings.Fields(authHeader)

		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondError(w, http.StatusUnauthorized, "invalid authorization header", nil)
			return
		}

		tokenStr := parts[1]

		secret := os.Getenv("JWT_SECRET")

		if secret == "" {
			utils.RespondError(w, http.StatusInternalServerError, "server configuration error", nil)
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, "invalid token", nil)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, "invalid token claims", nil)
			return
		}

		userID, ok := claims["user_id"].(string)

		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, "invalid token data", nil)
			return
		}

		sessionID, ok := claims["session_id"].(string)

		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, "invalid token data", nil)
			return
		}

		role, ok := claims["role"].(string)

		if !ok {
			utils.RespondError(
				w, http.StatusUnauthorized, "invalid token data", nil)
			return
		}

		dbUserID, err := dbhelpers.GetUserIDFromSession(sessionID)

		if err != nil || dbUserID != userID {
			utils.RespondError(w, http.StatusUnauthorized, "invalid session", nil)
			return
		}

		authContext := AuthContext{
			UserID:    userID,
			SessionID: sessionID,
			Role:      strings.ToLower(role),
		}

		ctx := context.WithValue(
			r.Context(),
			authContextKey,
			authContext,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequiredRoles(allowedRoles ...string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			auth, ok := GetAuthContext(r)

			if !ok {
				utils.RespondError(w, http.StatusUnauthorized, "unauthorized", nil)
				return
			}

			for _, role := range allowedRoles {

				if strings.ToLower(auth.Role) == strings.ToLower(role) {
					next.ServeHTTP(w, r)
					return
				}
			}

			utils.RespondError(w, http.StatusForbidden, "forbidden", nil)
		})
	}
}

func AdminOnly(next http.Handler) http.Handler {
	return RequiredRoles("admin")(next)
}

func GetAuthContext(r *http.Request) (AuthContext, bool) {
	auth, ok := r.Context().Value(authContextKey).(AuthContext)
	return auth, ok
}

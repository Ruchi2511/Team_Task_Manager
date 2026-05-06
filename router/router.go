package router

import (
	"Team_Task_Manager/handlers"
	"Team_Task_Manager/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() chi.Router {

	r := chi.NewRouter()

	r.Post("/register", handlers.RegisterUser)
	r.Post("/login", handlers.LoginUser)
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/logout", handlers.LogoutUser)
	})
	return r
}

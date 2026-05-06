package router

import (
	"Team_Task_Manager/handlers"
	"Team_Task_Manager/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.CORSMiddleware)

	r.Post("/register", handlers.RegisterUser)
	r.Post("/login", handlers.LoginUser)

	r.Group(func(r chi.Router) {

		r.Use(middleware.AuthMiddleware)

		r.Post("/logout", handlers.LogoutUser)

		r.Route("/projects", func(r chi.Router) {

			r.Get("/", handlers.GetProjects)

			r.Group(func(r chi.Router) {

				r.Use(middleware.RequiredRoles("admin"))

				r.Post("/", handlers.CreateProject)

				r.Post("/{id}/members", handlers.AddProjectMember)
			})
		})

		r.Route("/tasks", func(r chi.Router) {

			r.Get("/", handlers.GetTasks)

			r.Patch("/{id}/status", handlers.UpdateTaskStatus)

			r.Group(func(r chi.Router) {

				r.Use(middleware.RequiredRoles("admin"))

				r.Post("/", handlers.CreateTask)
			})
		})

		r.Route("/dashboard", func(r chi.Router) {

			r.Get("/stats", handlers.GetDashboardStats)
		})

		r.Route("/users", func(r chi.Router) {

			r.Get("/", handlers.GetUsers)
		})
	})

	return r
}

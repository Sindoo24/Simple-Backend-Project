package routes

import (
	"github.com/gofiber/fiber/v2"

	"BACKEND/internal/handler"
	"BACKEND/internal/middleware"
)

func Register(app *fiber.App, h *handler.UserHandler, authHandler *handler.AuthHandler, adminHandler *handler.AdminHandler, jwtSecret string) {
	
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())

	
	app.Post("/auth/signup", authHandler.Signup)
	app.Post("/auth/login", authHandler.Login)

	
	protected := app.Group("/users")
	protected.Use(middleware.Auth(jwtSecret))
	{
		protected.Get("/me", h.GetCurrentUser)
		protected.Post("/", h.Create)
		protected.Get("/:id", h.GetByID)
		protected.Get("/", h.List)
		protected.Put("/:id", h.Update)
		protected.Delete("/:id", h.Delete)
	}

	
	admin := app.Group("/admin")
	admin.Use(middleware.Auth(jwtSecret))
	admin.Use(middleware.RequireRole("admin"))
	{
		admin.Get("/users", adminHandler.GetAllUsers)
		admin.Get("/stats", adminHandler.GetStats)
	}
}

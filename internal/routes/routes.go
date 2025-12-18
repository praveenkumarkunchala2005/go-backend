package routes

import (
	"backend-task/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.CreateUser)
	app.Get("/users/:id", h.GetUserByID)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
	app.Get("/users", h.ListUsersPaginated)
	// app.Get("/users", h.ListUsers)
}

package routes

import (
	"github.com/fengjutian/azure-dragon-guard/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", handlers.GetUsers)
	api.Post("/users", handlers.CreateUser)
}

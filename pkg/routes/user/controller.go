package user

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	example := app.Group("/user")
	example.Post("/", createUserHandler)
	example.Get("/+", getUserHandler)
	example.Patch("/+", updateUserHandler)
	example.Delete("/+", deleteUserHandler)
}

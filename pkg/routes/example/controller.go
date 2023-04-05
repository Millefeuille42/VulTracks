package example

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	example := app.Group("/example")
	example.Get("/hello", helloHandler)
}

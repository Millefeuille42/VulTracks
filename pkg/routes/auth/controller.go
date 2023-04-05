package auth

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", loginHandler)
	auth.Get("/logout", logoutHandler)
	auth.Get("/whoami", whoAmIHandler)
}

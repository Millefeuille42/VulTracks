package settingsRoute

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	settings := app.Group("/settings")
	settings.Get("/", getSettingsHandler)
	settings.Put("/", updateSettingsHandler)
	settings.Get("/id3", getID3SettingsHandler)
	settings.Put("/id3", updateID3SettingsHandler)
	settings.Post("/restart", restartHandler)
}

package track

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	track := app.Group("/track")
	track.Get("/:id", getTrackHandler)
	//track.Get("/user/+", getUserTracksHandler)
	track.Post("/", createTrackHandler)
	track.Patch("/:id/id3", updateTrackId3Handler)
	//track.Patch("/:id", updateTrackHandler)
	track.Delete("/:id", deleteTrackHandler)
}

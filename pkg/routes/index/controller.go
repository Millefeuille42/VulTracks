package index

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	index := app.Group("/")
	index.Get("/", indexHandler)
	index.Get("/login", loginHandler)
	index.Get("/onboarding", onboardingHandler)
	index.Get("/edit-track/:id", editTrackHandler)
}

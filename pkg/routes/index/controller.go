package index

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	index := app.Group("/")
	index.Get("/", indexHandler)
	index.Get("/login", loginHandler)
	index.Get("/onboarding", onboardingHandler)
	index.Get("/restart", restartHandler)

	dashboard := index.Group("/dashboard")
	dashboard.Get("/tracks", dashboardTracksHandler)
	dashboard.Get("/tracks/edit/:id", dashboardTracksEdit)
	dashboard.Get("/folders", dashboardFoldersHandler)
	dashboard.Get("/settings", dashboardSettingsHandler)
}

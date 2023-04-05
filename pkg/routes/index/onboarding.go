package index

import (
	"VulTracks/pkg/globals"
	"github.com/gofiber/fiber/v2"
)

func onboardingHandler(c *fiber.Ctx) error {
	if !globals.FirstRun {
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	return c.Render("onboarding", fiber.Map{})
}

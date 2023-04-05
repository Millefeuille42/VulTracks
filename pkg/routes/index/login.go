package index

import (
	"VulTracks/pkg/globals"
	"github.com/gofiber/fiber/v2"
)

func loginHandler(c *fiber.Ctx) error {
	if globals.FirstRun {
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}

	if c.Locals("name") != nil && c.Locals("name") != "" {
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	return c.Render("login", fiber.Map{})
}

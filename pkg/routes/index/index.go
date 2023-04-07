package index

import (
	"github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
	return c.Redirect("/dashboard/tracks", fiber.StatusSeeOther)
}

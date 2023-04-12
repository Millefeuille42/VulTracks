package index

import (
	"VulTracks/pkg/utils/id3Utils"
	"VulTracks/pkg/utils/settings"
	"github.com/gofiber/fiber/v2"
)

func dashboardSettingsHandler(c *fiber.Ctx) error {
	return c.Render("settings", fiber.Map{
		"username":  c.Locals("name").(string),
		"settings":  settings.ToParameters("", nil),
		"id3Frames": id3Utils.ID3Frames,
		"sections":  getSections("settings"),
	})
}

package index

import (
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"VulTracks/pkg/utils/settings"
	"github.com/gofiber/fiber/v2"
)

func dashboardSettingsHandler(c *fiber.Ctx) error {
	return c.Render("settings", fiber.Map{
		"username":  c.Locals("name").(string),
		"settings":  utils.StructToMap(settings.Settings),
		"id3Frames": id3Utils.ID3Frames,
		"sections":  getSections("settings"),
	})
}

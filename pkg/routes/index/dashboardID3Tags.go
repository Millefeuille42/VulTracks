package index

import (
	"VulTracks/pkg/utils/id3Utils"
	"github.com/gofiber/fiber/v2"
)

func dashboardID3TagsHandler(c *fiber.Ctx) error {
	return c.Render("id3Tags", fiber.Map{
		"username":  c.Locals("name").(string),
		"id3Frames": id3Utils.ID3Frames,
		"sections":  getSections("id3Tags"),
	})
}

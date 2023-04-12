package settingsRoute

import (
	"VulTracks/pkg/utils/id3Utils"
	"github.com/gofiber/fiber/v2"
)

func getID3SettingsHandler(c *fiber.Ctx) error {
	return c.JSON(id3Utils.ID3Frames)
}

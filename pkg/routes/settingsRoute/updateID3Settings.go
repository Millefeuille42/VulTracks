package settingsRoute

import (
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"github.com/gofiber/fiber/v2"
)

func updateID3SettingsHandler(c *fiber.Ctx) error {
	newFrames := make([]id3Utils.ID3Frame, 0)

	if err := c.BodyParser(&newFrames); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	id3Utils.ID3Frames = newFrames

	err := id3Utils.RewriteID3Frames()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(id3Utils.ID3Frames)
}

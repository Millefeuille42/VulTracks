package settingsRoute

import (
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/settings"
	"github.com/gofiber/fiber/v2"
)

func updateSettingsHandler(c *fiber.Ctx) error {
	newSettings := make(map[string]string)

	if err := c.BodyParser(&newSettings); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	settings.Settings = newSettings

	err := settings.RewriteSettings()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(settings.Settings)
}

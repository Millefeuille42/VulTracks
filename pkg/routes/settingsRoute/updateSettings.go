package settingsRoute

import (
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/settings"
	"VulTracks/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func updateSettingsHandler(c *fiber.Ctx) error {
	newSettings := new(settings.SettingsStruct)

	if err := c.BodyParser(newSettings); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	errorsList := validator.ValidateStruct(*newSettings, false)
	if errorsList != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorsList,
		})
	}

	settings.Settings = *newSettings

	err := settings.RewriteSettings()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(settings.Settings)
}

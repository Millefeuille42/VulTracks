package settingsRoute

import (
	"VulTracks/pkg/utils/settings"
	"github.com/gofiber/fiber/v2"
)

func getSettingsHandler(c *fiber.Ctx) error {
	return c.JSON(settings.Settings)
}

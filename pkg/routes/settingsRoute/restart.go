package settingsRoute

import (
	"VulTracks/pkg/app"
	"VulTracks/pkg/store"
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func restartHandler(c *fiber.Ctx) error {
	session, err := store.Store.Sessions.Get(c)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = session.Destroy()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = app.StopApp(app.App)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

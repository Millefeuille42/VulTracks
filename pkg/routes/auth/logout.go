package auth

import (
	"VulTracks/pkg/store"
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func logoutHandler(c *fiber.Ctx) error {
	session, err := store.Store.Sessions.Get(c)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = session.Destroy()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}
	return c.Redirect("/login", fiber.StatusSeeOther)
}

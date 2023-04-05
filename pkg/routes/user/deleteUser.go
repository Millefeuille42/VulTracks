package user

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func deleteUserHandler(c *fiber.Ctx) error {
	user := new(models.UserModel)
	userParam := c.Params("+")

	_, err := user.GetUserByIdOrUsername(userParam)
	if err != nil {
		if err.Error() == "Not Found" {
			return utils.ReturnError(c, fiber.StatusNotFound, err)
		}

		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = user.DeleteUser()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

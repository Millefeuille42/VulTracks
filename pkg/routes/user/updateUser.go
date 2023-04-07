package user

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func updateUserHandler(c *fiber.Ctx) error {
	user := new(models.UserModel)

	if err := c.BodyParser(user); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	errorsList := validator.ValidateStruct(*user, false)
	if errorsList != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorsList,
		})
	}

	userParam := c.Params("+")
	oldUser := new(models.UserModel)
	_, err := oldUser.GetUserByIdOrUsername(userParam)
	if err != nil {
		if err.Error() == "Not Found" {
			return utils.ReturnError(c, fiber.StatusNotFound, err)
		}
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = oldUser.UpdateUser(user)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	oldUser.Password = ""
	return c.Status(fiber.StatusOK).JSON(oldUser)
}

package user

import "C"
import (
	"VulTracks/pkg/globals"
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func createUserHandler(c *fiber.Ctx) error {
	user := new(models.UserModel)

	if err := c.BodyParser(user); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	errorsList := validator.ValidateStruct(*user, true)
	if errorsList != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorsList,
		})
	}

	err := user.HashPassword()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	err = user.CreateUser()
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return utils.ReturnError(c, fiber.StatusConflict, err)
		}
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	globals.FirstRun = false
	return c.Status(fiber.StatusOK).JSON(user)
}

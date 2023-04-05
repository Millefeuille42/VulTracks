package auth

import (
	"VulTracks/pkg/interfaces"
	"VulTracks/pkg/models"
	"VulTracks/pkg/store"
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func loginHandler(c *fiber.Ctx) error {
	credentials := new(interfaces.AuthInterface)

	if err := c.BodyParser(credentials); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing data",
			"error":   err.Error(),
		})
	}

	user := new(models.UserModel)
	_, err := user.GetUserByUsername(credentials.Username)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusNotFound, fiber.ErrUnauthorized)
	}

	err = user.ComparePassword(credentials.Password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return utils.ReturnError(c, fiber.StatusNotFound, fiber.ErrUnauthorized)
		}
		return utils.ReturnError(c, fiber.StatusUnauthorized, err)
	}

	session, err := store.Store.Sessions.Get(c)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	session.Set("user", user.Username)
	session.Set("id", user.Id)

	err = session.Save()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

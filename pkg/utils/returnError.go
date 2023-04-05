package utils

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func ReturnError(c *fiber.Ctx, status int, err error) error {
	log.Println(err.Error())
	return c.Status(status).JSON(fiber.Map{
		"message": err.Error(),
	})
}

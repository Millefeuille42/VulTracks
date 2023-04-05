package auth

import (
	"github.com/gofiber/fiber/v2"
)

func whoAmIHandler(c *fiber.Ctx) error {
	return c.Status(200).SendString("Hello " + c.Locals("name").(string))
}

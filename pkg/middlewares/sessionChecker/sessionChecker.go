package sessionChecker

import (
	"VulTracks/pkg/globals"
	"VulTracks/pkg/store"
	"VulTracks/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
}

func DefaultFilter(c *fiber.Ctx) bool {
	excluded := []string{"/login", "/onboarding", "/static", "/auth/login"}
	for _, path := range excluded {
		if strings.HasPrefix(c.Path(), path) {
			return true
		}
	}
	return c.Path() == "/user" && c.Method() == fiber.MethodPost && globals.FirstRun
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Filter != nil && config.Filter(c) {
			fmt.Println("SessionChecker: Skipping middleware")
			return c.Next()
		}

		session, err := store.Store.Sessions.Get(c)
		if err != nil {
			return utils.ReturnError(c, fiber.StatusInternalServerError, err)
		}

		name := session.Get("user")
		if name == nil {
			return c.Redirect("/login", fiber.StatusSeeOther)
		}

		if c.Path() == "/login" {
			return c.Redirect("/", fiber.StatusSeeOther)
		}

		c.Locals("name", name)
		c.Locals("id", session.Get("id"))

		return c.Next()
	}
}

package index

import (
	"VulTracks/pkg/utils/settings"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func restartHandler(c *fiber.Ctx) error {
	hostname := strings.Split(c.Hostname(), ":")
	return c.Render("restart", fiber.Map{
		"accessURL": c.Protocol() + "://" + hostname[0] + ":" + settings.Settings.Port,
	})
}

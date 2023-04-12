package app

import (
	"VulTracks/pkg/globals"
	"VulTracks/pkg/utils/settings"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
	"time"
)

func NewApp() *fiber.App {
	engine := jet.New(globals.TemplateLocation, ".jet.html")

	app := fiber.New(fiber.Config{
		AppName:     "VulTracks",
		Views:       engine,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	return app
}

func StartApp(app *fiber.App) error {
	return app.Listen(fmt.Sprintf("%s:%s", settings.Settings["server.host"], settings.Settings["server.port"]))
}

func StopApp(app *fiber.App) error {
	return app.ShutdownWithTimeout(time.Duration(5) * time.Second)
}

var App *fiber.App

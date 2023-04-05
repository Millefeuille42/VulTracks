package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

// SetupCloseHandler Set up a handler for Ctrl+C and closing the bot
func SetupCloseHandler(app *fiber.App) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		_ = app.Shutdown()
		os.Exit(0)
	}()
}

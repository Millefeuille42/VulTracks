package folder

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	folder := app.Group("/folder")
	folder.Get("/user", getFoldersOfUserHandler)
	folder.Post("/:id/sync", syncFolderHandler)
	folder.Post("/", createFolderHandler)
	folder.Delete("/:id", deleteFolderHandler)
}

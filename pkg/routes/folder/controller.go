package folder

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	folder := app.Group("/folder")
	//folder.Get("/:id", getFolderHandler)
	//track.Get("/user/+", getUserFoldersHandler)
	folder.Post("/", createFolderHandler)
	//track.Patch("/:id", updateFolderHandler)
	//folder.Delete("/:id", deleteFolderHandler)
}

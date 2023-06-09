package folder

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func getFoldersOfUserHandler(c *fiber.Ctx) error {
	userId := c.Locals("id").(string)

	folders, err := models.GetCountPerFolderByUserId(userId)
	if err != nil {
		if err.Error() == "Not Found" {
			return c.Status(fiber.StatusOK).JSON([]models.FolderModel{})
		}
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(folders)
}

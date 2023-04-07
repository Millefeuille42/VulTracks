package folder

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func syncFolderHandler(c *fiber.Ctx) error {
	folder := new(models.FolderModel)
	trackID := c.Params("id")

	_, err := folder.GetFolderById(trackID)
	if err != nil {
		if err.Error() == "Not Found" {
			return utils.ReturnError(c, fiber.StatusNotFound, err)
		}
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	if folder.UserId != c.Locals("id").(string) {
		return utils.ReturnError(c, fiber.StatusUnauthorized, err)
	}

	folders, err := syncTracksOfFolder(folder.Path, folder.UserId, folder.Id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(folders)
}

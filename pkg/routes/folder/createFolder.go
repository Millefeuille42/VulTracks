package folder

import (
	"VulTracks/pkg/interfaces"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"os"
)

func createFolderHandler(c *fiber.Ctx) error {
	folderData := new(interfaces.FolderInterface)

	if err := c.BodyParser(folderData); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	errorsList := validator.ValidateStruct(*folderData, true)
	if errorsList != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorsList,
		})
	}

	info, err := os.Stat(folderData.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return utils.ReturnError(c, fiber.StatusNotFound, err)
		}
		return err
	}
	if !info.IsDir() {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	folders, err := syncTracksOfFolder(folderData.Path, c.Locals("id").(string), "0")
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(folders)
}

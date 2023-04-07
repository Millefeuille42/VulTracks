package index

import (
	"VulTracks/pkg/interfaces"
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/folderTree"
	"VulTracks/pkg/utils/id3Utils"
	"github.com/gofiber/fiber/v2"
)

func dashboardFoldersHandler(c *fiber.Ctx) error {
	id := c.Locals("id")

	folders, err := models.GetCountPerFolderByUserId(id.(string))
	if err != nil {
		if err.Error() != "Not Found" {
			return utils.ReturnError(c, fiber.StatusInternalServerError, err)
		}
		folders = []interfaces.CountPerFolderInterface{}
	}

	foldersTree := folderTree.BuildTree(folders, "")

	return c.Render("folders", fiber.Map{
		"username":     c.Locals("name").(string),
		"folders":      foldersTree,
		"foldersCount": len(folders),
		"ID3Frames":    id3Utils.ID3Frames,
		"sections":     getSections("folders"),
	})
}

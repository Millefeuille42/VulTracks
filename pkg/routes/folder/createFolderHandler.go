package folder

import (
	"VulTracks/pkg/interfaces"
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"VulTracks/pkg/validator"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"strings"
	"time"
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

	dirFiles, err := os.ReadDir(folderData.Path)
	if err != nil {
		return err
	}

	for _, file := range dirFiles {
		if file.IsDir() {
			continue
		}
	}

	folders := make([]models.FolderModel, 0)

	err = utils.RecursiveReadDir(folderData.Path, func(path string, files []os.DirEntry) error {
		filteredFiles := make([]os.DirEntry, 0)
		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(file.Name(), ".mp3") {
				continue
			}
			filteredFiles = append(filteredFiles, file)
		}
		if len(filteredFiles) == 0 {
			return nil
		}

		folder := new(models.FolderModel)
		folder.Path = path
		folder.LastScan = time.Now().String()
		folder.UserId = c.Locals("id").(string)
		err = folder.CreateFolder()
		if err != nil {
			if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return utils.ReturnError(c, fiber.StatusInternalServerError, err)
			}
		}
		folders = append(folders, *folder)

		for _, file := range filteredFiles {
			track := new(models.TrackModel)
			track.Path = path + "/" + file.Name()
			err = id3Utils.SetTagOfTrack(*track, "TIT2", strings.TrimRight(file.Name(), ".mp3"), true)
			if err != nil {
				log.Println(err)
				continue
			}
			track.UserId = c.Locals("id").(string)
			track.Name = file.Name()
			track.FolderId = sql.NullString{
				String: folder.Id,
				Valid:  true,
			}
			err = track.CreateTrack()
			if err != nil {
				if strings.Contains(err.Error(), "UNIQUE constraint failed") {
					continue
				}
				return utils.ReturnError(c, fiber.StatusInternalServerError, err)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(folders)
}

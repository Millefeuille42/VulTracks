package track

import (
	"VulTracks/pkg/interfaces"
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/validator"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mikkyang/id3-go"
	"github.com/mikkyang/id3-go/v2"
	"strings"
)

func createTrackHandler(c *fiber.Ctx) error {
	trackData := new(interfaces.TrackInterface)

	if err := c.BodyParser(trackData); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	errorsList := validator.ValidateStruct(*trackData, true)
	if errorsList != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorsList,
		})
	}

	if !strings.Contains(trackData.Path, ".mp3") {
		return utils.ReturnError(c, fiber.StatusBadRequest, errors.New("file is not mp3"))
	}

	file, err := id3.Open(trackData.Path)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return utils.ReturnError(c, fiber.StatusNotFound, err)
		}
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}
	defer file.Close()

	trackTitle := file.Frame("TIT2")
	if trackTitle == nil {
		frame := v2.V23FrameTypeMap["TIT2"]
		textFrame := v2.NewTextFrame(frame, trackData.TrackNameFallback)
		file.AddFrames(textFrame)
	}

	track := new(models.TrackModel)
	track.Path = trackData.Path
	track.UserId = c.Locals("id").(string)
	track.Name = trackData.TrackNameFallback
	err = track.CreateTrack()
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return utils.ReturnError(c, fiber.StatusConflict, err)
		}
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(track)
}

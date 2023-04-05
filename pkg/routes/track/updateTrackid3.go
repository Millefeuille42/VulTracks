package track

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"github.com/gofiber/fiber/v2"
)

func updateTrackId3Handler(c *fiber.Ctx) error {
	track := new(models.TrackModel)
	trackID := c.Params("id")

	_, err := track.GetTrackById(trackID)
	if err != nil {
		if err.Error() == "Not Found" {
			return utils.ReturnError(c, fiber.StatusNotFound, err)
		}

		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	if track.UserId != c.Locals("id").(string) {
		return utils.ReturnError(c, fiber.StatusUnauthorized, err)
	}

	newTrackData := map[string]string{}
	if err = c.BodyParser(&newTrackData); err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err)
	}

	oldTrackData := id3Utils.GetTagsOfTrack(*track)
	for key, value := range oldTrackData {
		if newTrackData[key] == "" {
			newTrackData[key] = value
		}
	}
	err = id3Utils.SetTagsOfTrack(*track, newTrackData)

	return c.SendStatus(fiber.StatusOK)
}

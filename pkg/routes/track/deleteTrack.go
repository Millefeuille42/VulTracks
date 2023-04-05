package track

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func deleteTrackHandler(c *fiber.Ctx) error {
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

	err = track.DeleteTrack()
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

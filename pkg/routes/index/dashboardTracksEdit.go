package index

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"github.com/gofiber/fiber/v2"
)

func dashboardTracksEdit(c *fiber.Ctx) error {
	id := c.Params("id")
	track := new(models.TrackModel)

	_, err := track.GetTrackById(id)
	if err != nil {
		if err.Error() == "Not Found" {
			return utils.ReturnError(c, fiber.StatusNotFound, err)
		}
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}
	trackData := id3Utils.GetTagsOfTrack(*track)

	return c.Render("editTrack", fiber.Map{
		"track":     trackData,
		"ID3Frames": id3Utils.ID3Frames,
	})
}

package index

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"VulTracks/pkg/utils/settings"
	"github.com/gofiber/fiber/v2"
)

func dashboardTracksHandler(c *fiber.Ctx) error {
	id := c.Locals("id")

	tracks, err := models.GetTracksByUserId(id.(string))
	if err != nil {
		if err.Error() != "Not Found" {
			return utils.ReturnError(c, fiber.StatusInternalServerError, err)
		}
		tracks = []models.TrackModel{}
	}
	tracksData := make([]map[string]string, 0)
	for _, track := range tracks {
		trackData := id3Utils.GetTagsOfTrack(track)
		tracksData = append(tracksData, trackData)
	}

	return c.Render("tracks", fiber.Map{
		"username":    c.Locals("name").(string),
		"tracks":      tracksData,
		"tracksCount": len(tracks),
		"ID3Frames":   id3Utils.ID3Frames,
		"sections":    getSections("tracks"),
		"heading":     settings.Settings.Heading,
	})
}

package track

import (
	"VulTracks/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/mikkyang/id3-go"
)

func getTrackHandler(c *fiber.Ctx) error {
	//track := new(models.TrackModel)
	//id := c.Params("id")

	//_, err := track.GetTrackById(id)
	//if err != nil {
	//	if err.Error() == "Not Found" {
	//		return utils.ReturnError(c, fiber.StatusNotFound, err)
	//	}
	//	return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	//}

	file, err := id3.Open("./data/sample.mp3")
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}
	defer file.Close()
	frames := file.Frame("TIT2")
	if frames == nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).Send([]byte(file.Frame("TIT2").String()))
}

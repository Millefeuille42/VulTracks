package index

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"VulTracks/pkg/utils/settings"
	"github.com/gofiber/fiber/v2"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseDashboardTracksQuery(c *fiber.Ctx) (int, int, string, string) {
	limit := c.QueryInt("limit")
	order := c.Query("order")
	sortBy := c.Query("sortBy")

	if order == "" {
		order = settings.Settings["dashboard.tracks.default_order"]
	}

	if sortBy == "" {
		sortBy = settings.Settings["dashboard.tracks.default_sort"]
	}

	if c.Query("limit") == "" {
		l64, err := strconv.ParseInt(settings.Settings["dashboard.tracks.default_per_page"], 10, 64)
		if err != nil {
			log.Println(err)
			limit = -1
		} else {
			limit = int(l64)
		}
	}

	page := c.QueryInt("page")
	if page < 0 {
		page = 0
	}
	offset := page * limit

	return limit, offset, sortBy, order
}

func dashboardTracksHandler(c *fiber.Ctx) error {
	id := c.Locals("id")
	limit, offset, sortBy, order := parseDashboardTracksQuery(c)

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

	sort.SliceStable(tracksData, func(i, j int) bool {
		if tracksData[i][sortBy] == "" {
			return false
		}
		if tracksData[j][sortBy] == "" {
			return true
		}
		if order == "DESC" {
			return tracksData[i][sortBy] > tracksData[j][sortBy]
		}
		return tracksData[i][sortBy] < tracksData[j][sortBy]
	})

	numberPages := 0
	if limit > 0 {
		numberPages = len(tracksData) / limit
		if len(tracksData)%limit != 0 {
			numberPages += 1
		}
	}

	if limit != -1 {
		if offset == -1 {
			offset = 0
		}
		if offset > len(tracksData) {
			tracksData = []map[string]string{}
		} else if offset+limit > len(tracksData) {
			tracksData = tracksData[offset:]
		} else {
			tracksData = tracksData[offset : offset+limit]
		}
	}

	return c.Render("tracks", fiber.Map{
		"username":    c.Locals("name").(string),
		"tracks":      tracksData,
		"tracksCount": len(tracks),
		"ID3Frames":   id3Utils.ID3Frames,
		"sections":    getSections("tracks"),
		"heading":     settings.Settings["dashboard.tracks.heading"],
		"sortOrder":   strings.ToLower(order),
		"sortField":   sortBy,
		"numberPages": numberPages,
		"currentPage": c.QueryInt("page"),
	})
}

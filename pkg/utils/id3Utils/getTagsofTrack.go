package id3Utils

import (
	"VulTracks/pkg/models"
	"github.com/mikkyang/id3-go"
	"strconv"
	"strings"
	"unicode"
)

func GetTagsOfTrack(track models.TrackModel) map[string]string {
	trackData := make(map[string]string)
	trackData["id"] = track.Id

	trackFile, err := id3.Open(track.Path)
	hasError := err != nil
	if !hasError {
		defer trackFile.Close()
	}
	trackData["hasError"] = strconv.FormatBool(hasError)
	for _, frame := range ID3Frames {
		if hasError {
			if frame.Tag == "TIT2" {
				trackData[frame.Tag] = track.Name
				continue
			}
			trackData[frame.Tag] = "ERROR PATH"
			continue
		}
		frameValue := trackFile.Frame(frame.Tag)
		if frameValue == nil {
			trackData[frame.Tag] = ""
			continue
		}
		trackData[frame.Tag] = strings.TrimFunc(frameValue.String(), func(r rune) bool {
			return !unicode.IsGraphic(r)
		})
	}
	return trackData
}

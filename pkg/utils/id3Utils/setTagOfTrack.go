package id3Utils

import (
	"VulTracks/pkg/models"
	"github.com/mikkyang/id3-go"
	v2 "github.com/mikkyang/id3-go/v2"
)

func SetTagOfTrack(track models.TrackModel, tagName, tagValue string, replace bool) error {
	trackFile, err := id3.Open(track.Path)
	if err != nil {
		return err
	}
	defer trackFile.Close()

	frameType := v2.V23FrameTypeMap[tagName]
	textFrame := v2.NewTextFrame(frameType, tagValue)
	if replace {
		trackFile.DeleteFrames(tagName)
	}
	trackFile.AddFrames(textFrame)

	return nil
}

package id3Utils

import (
	"VulTracks/pkg/models"
	"github.com/mikkyang/id3-go"
	"github.com/mikkyang/id3-go/v2"
)

func SetTagsOfTrack(track models.TrackModel, trackData map[string]string) error {
	trackFile, err := id3.Open(track.Path)
	if err != nil {
		return err
	}
	defer trackFile.Close()

	for _, frame := range ID3Frames {
		if trackData[frame.Tag] == "" {
			continue
		}
		frameType := v2.V23FrameTypeMap[frame.Tag]
		textFrame := v2.NewTextFrame(frameType, trackData[frame.Tag])
		trackFile.DeleteFrames(frame.Tag)
		trackFile.AddFrames(textFrame)
	}

	return nil
}

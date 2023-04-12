package id3Utils

import (
	"VulTracks/pkg/globals"
	"github.com/goccy/go-json"
	"os"
)

type ID3Frame struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

var (
	ID3Frames []ID3Frame
)

func RefreshID3Frames() error {
	data, err := os.ReadFile(globals.ConfigLocation + "/ID3Frames.json")
	if err != nil {
		return err
	}
	if ID3Frames == nil {
		ID3Frames = make([]ID3Frame, 0)
	}
	err = json.Unmarshal(data, &ID3Frames)
	if err != nil {
		return err
	}
	return nil
}

func RewriteID3Frames() error {
	data, err := json.MarshalIndent(ID3Frames, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(globals.ConfigLocation+"/ID3Frames.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

package settings

import (
	"VulTracks/pkg/globals"
	"github.com/goccy/go-json"
	"os"
)

type SettingsStruct struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

var Settings = SettingsStruct{}

func RefreshSettings() error {
	data, err := os.ReadFile(globals.ConfigLocation + "/config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &Settings)
	if err != nil {
		return err
	}
	return nil
}

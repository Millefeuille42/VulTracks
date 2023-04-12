package settings

import (
	"VulTracks/pkg/globals"
	"github.com/goccy/go-json"
	"os"
)

type SettingsStruct struct {
	Port    string `json:"port" validate:"required,numeric"`
	Host    string `json:"host" validate:"required,ip"`
	Heading string `json:"heading" validate:"required"`
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

func RewriteSettings() error {
	data, err := json.MarshalIndent(Settings, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(globals.ConfigLocation+"/config.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

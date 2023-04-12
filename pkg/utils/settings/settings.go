package settings

import (
	"VulTracks/pkg/globals"
	"github.com/goccy/go-json"
	"os"
)

type Parameter struct {
	Name          string      `json:"name" validate:"required"`
	Type          string      `json:"type" validate:"required"`
	Description   string      `json:"description,omitempty"`
	CausesRestart bool        `json:"causes_restart,omitempty"`
	DefaultValue  string      `json:"default_value,omitempty"`
	Fields        []Parameter `json:"fields,omitempty"`
}

var Defaults = make([]Parameter, 0)
var Settings = make(map[string]string)

func defaultsToSettings(parent string, list []Parameter, settings map[string]string) map[string]string {
	for _, param := range list {
		if param.Type == "section" {
			settings = defaultsToSettings(parent+param.Name+".", param.Fields, settings)
		} else {
			settings[parent+param.Name] = param.DefaultValue
		}
	}
	return settings
}

func ParseDefaults() error {
	data, err := os.ReadFile(globals.ConfigLocation + "/defaults.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &Defaults)
	if err != nil {
		return err
	}

	Settings = defaultsToSettings("", Defaults, Settings)
	return nil
}

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

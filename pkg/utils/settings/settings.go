package settings

import (
	"VulTracks/pkg/globals"
	"github.com/goccy/go-json"
	"os"
	"strings"
)

type Parameter struct {
	FullName      string      `json:"full_name,omitempty"`
	Name          string      `json:"name" validate:"required"`
	Type          string      `json:"type" validate:"required"`
	Description   string      `json:"description,omitempty"`
	CausesRestart bool        `json:"causes_restart,omitempty"`
	DefaultValue  string      `json:"default_value,omitempty"`
	Choices       []string    `json:"choices,omitempty"`
	Fields        []Parameter `json:"fields,omitempty"`
}

var Defaults = make([]Parameter, 0)
var Settings = make(map[string]string)

func ToParameters(parent string, list []Parameter) []Parameter {
	ret := make([]Parameter, 0)

	if list == nil {
		err := ParseDefaults(false)
		if err != nil {
			return nil
		}
		list = Defaults
	}
	for _, param := range list {
		if param.Type == "section" {
			param.Name = strings.Replace(param.Name, "_", " ", -1)
			ret = append(ret, param)
			ret = append(ret, ToParameters(parent+param.Name+"-", param.Fields)...)
		} else {
			param.DefaultValue = Settings[parent+param.Name]
			param.FullName = parent + param.Name
			param.Name = strings.Replace(param.Name, "_", " ", -1)
			ret = append(ret, param)
		}
	}
	return ret
}

func defaultsToSettings(parent string, list []Parameter, settings map[string]string) map[string]string {
	for _, param := range list {
		if param.Type == "section" {
			settings = defaultsToSettings(parent+param.Name+"-", param.Fields, settings)
		} else {
			settings[parent+param.Name] = param.DefaultValue
		}
	}
	return settings
}

func ParseDefaults(set bool) error {
	data, err := os.ReadFile(globals.ConfigLocation + "/defaults.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &Defaults)
	if err != nil {
		return err
	}

	if set {
		Settings = defaultsToSettings("", Defaults, Settings)
	}
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

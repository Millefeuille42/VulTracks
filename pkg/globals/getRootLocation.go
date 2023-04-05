package globals

import "os"

func getRootLocation() string {
	if Dev {
		return "./data"
	}

	location, err := os.UserConfigDir()
	if err == nil && location != "" {
		return location + "/VulTracks"
	}

	location, err = os.UserHomeDir()
	if err == nil && location != "" {
		return location + "/.config" + "/VulTracks"
	}

	return ""
}

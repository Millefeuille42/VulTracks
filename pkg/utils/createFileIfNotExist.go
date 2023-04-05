package utils

import "os"

// CreateFileIfNotExist Check if file exists, if not create it
func CreateFileIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(path)
		}
	}
	return err
}

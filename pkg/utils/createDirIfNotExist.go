package utils

import "os"

// CreateDirIfNotExist Check if dir exists, if not create it
func CreateDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModePerm)
		}
	}
	return err
}

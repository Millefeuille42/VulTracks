package utils

import (
	"fmt"
	"os"
	"strings"
)

func RecursiveReadDir(path string, dirFunc func(string, []os.DirEntry) error) error {
	strings.TrimRight(path, "/")

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	err = dirFunc(path, files)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		newPath := fmt.Sprintf("%s/%s", path, file.Name())
		err = RecursiveReadDir(newPath, dirFunc)
		if err != nil {
			return err
		}
	}
	return nil
}

package folder

import (
	"VulTracks/pkg/models"
	"VulTracks/pkg/utils"
	"VulTracks/pkg/utils/id3Utils"
	"database/sql"
	"log"
	"os"
	"strings"
	"time"
)

func getFilteredFiles(files []os.DirEntry) []os.DirEntry {
	filteredFiles := make([]os.DirEntry, 0)
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".mp3") {
			continue
		}
		filteredFiles = append(filteredFiles, file)
	}
	return filteredFiles
}

func createTracksOfFolder(folder *models.FolderModel, files []os.DirEntry) error {
	for _, file := range files {
		track := new(models.TrackModel)
		track.Path = folder.Path + "/" + file.Name()
		err := id3Utils.SetTagOfTrack(*track, "TIT2", strings.TrimRight(file.Name(), ".mp3"), true)
		if err != nil {
			log.Println(err)
			continue
		}
		track.UserId = folder.UserId
		track.Name = file.Name()
		track.FolderId = sql.NullString{
			String: folder.Id,
			Valid:  true,
		}
		err = track.CreateTrack()
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				continue
			}
			return err
		}
	}
	return nil
}

func syncTracksOfFolder(folderPath, userId, folderId string) ([]models.FolderModel, error) {
	folders := make([]models.FolderModel, 0)

	err := utils.RecursiveReadDir(folderPath, func(path string, files []os.DirEntry) error {
		filteredFiles := getFilteredFiles(files)
		if len(filteredFiles) == 0 {
			return nil
		}

		folder := new(models.FolderModel)
		folder.Path = path
		folder.LastScan = time.Now().String()
		folder.UserId = userId
		err := folder.CreateFolder()
		if err != nil {
			if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return err
			}
			if folderId == "0" {
				return nil
			}
			_, err = folder.GetFolderById(folderId)
			if err != nil {
				return err
			}
			folder.LastScan = time.Now().String()
			err = folder.UpdateFolder()
			if err != nil {
				return err
			}
		} else {
			folders = append(folders, *folder)
		}

		return createTracksOfFolder(folder, filteredFiles)
	})

	return folders, err
}

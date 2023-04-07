package models

import (
	"VulTracks/pkg/database"
	"VulTracks/pkg/interfaces"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"path/filepath"
)

type FolderModel struct {
	Id       string         `json:"id"`
	Path     string         `json:"path" validate:"required"`
	LastScan string         `json:"last_scan" validate:"required"`
	UserId   string         `json:"user_id" validate:"required"`
	ParentId sql.NullString `json:"parent_id"`
}

func (folder *FolderModel) CreateTable() error {
	_, err := database.Database.Exec(`
		CREATE TABLE IF NOT EXISTS folders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL,
			last_scan TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			parent_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY(parent_id) REFERENCES folders(id) ON DELETE CASCADE,
        	UNIQUE(path, user_id)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func appendFoldersToList(list []FolderModel, rows *sql.Rows) ([]FolderModel, error) {
	var folder FolderModel
	err := rows.Scan(&folder.Id, &folder.Path, &folder.LastScan, &folder.UserId, &folder.ParentId)
	if err != nil {
		return nil, err
	}
	list = append(list, folder)
	return list, nil
}

func getFoldersListFromRows(rows *sql.Rows) ([]FolderModel, error) {
	var list []FolderModel
	var err error
	list, err = appendFoldersToList(list, rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list, err = appendFoldersToList(list, rows)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func (folder *FolderModel) getFolderByQuery(query squirrel.SelectBuilder) error {
	rows, err := database.SelectHelper(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = rows.Scan(&folder.Id, &folder.Path, &folder.LastScan, &folder.UserId, &folder.ParentId)
	if err != nil {
		return err
	}
	return nil
}

func GetFolders() ([]FolderModel, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("*").
			From("folders"),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	folders, err := getFoldersListFromRows(rows)
	return folders, err
}

func (folder *FolderModel) GetFolderById(id string) (*FolderModel, error) {
	return folder, folder.getFolderByQuery(
		squirrel.
			Select("*").
			From("folders").
			Where(squirrel.Eq{"id": id}),
	)
}

func (folder *FolderModel) GetFolderByPathAndUserId(path, userId string) (*FolderModel, error) {
	return folder, folder.getFolderByQuery(
		squirrel.
			Select("*").
			From("folders").
			Where(squirrel.And{squirrel.Eq{"path": path}, squirrel.Eq{"user_id": userId}}),
	)
}

func GetCountPerFolderByUserId(userId string) ([]interfaces.CountPerFolderInterface, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("folders.id", "folders.path", "folders.last_scan", "folders.parent_id", "COUNT(tracks.id) AS num_tracks").
			From("folders").
			LeftJoin("tracks ON folders.id = tracks.folder_id").
			Where(squirrel.Eq{"folders.user_id": userId}).
			GroupBy("folders.id"),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]interfaces.CountPerFolderInterface, 0)

	var folder interfaces.CountPerFolderInterface
	err = rows.Scan(&folder.Id, &folder.Path, &folder.LastScan, &folder.ParentId, &folder.Count)
	if err != nil {
		return nil, err
	}
	folder.Name = filepath.Base(folder.Path)
	list = append(list, folder)
	for rows.Next() {
		err = rows.Scan(&folder.Id, &folder.Path, &folder.LastScan, &folder.ParentId, &folder.Count)
		if err != nil {
			return nil, err
		}
		folder.Name = filepath.Base(folder.Path)
		list = append(list, folder)
	}
	return list, nil
}

func GetFoldersByUserId(userId string) ([]FolderModel, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("*").
			From("folders").
			Where(squirrel.Eq{"user_id": userId}),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getFoldersListFromRows(rows)
}

func (folder *FolderModel) CreateFolder() error {
	result, err := squirrel.
		Insert("folders").
		Columns("path", "last_scan", "user_id", "parent_id").
		Values(folder.Path, folder.LastScan, folder.UserId, folder.ParentId).
		RunWith(database.Database).Exec()
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	folder.Id = fmt.Sprintf("%d", id)
	return nil
}

func (folder *FolderModel) DeleteFolder() error {
	_, err := squirrel.
		Delete("folders").
		Where(squirrel.Eq{"id": folder.Id}).
		RunWith(database.Database).Exec()
	return err
}

func (folder *FolderModel) UpdateFolder() error {
	folderQuery := squirrel.Update("folders")

	if folder.Path != "" {
		folderQuery = folderQuery.Set("path", folder.Path)
	}

	if folder.UserId != "" {
		folderQuery = folderQuery.Set("user_id", folder.UserId)
	}

	if folder.LastScan != "" {
		folderQuery = folderQuery.Set("last_scan", folder.LastScan)
	}

	if folder.ParentId.String != "" {
		folderQuery = folderQuery.Set("parent_id", folder.ParentId)
	}

	_, err := folderQuery.
		Where(squirrel.Eq{"id": folder.Id}).
		RunWith(database.Database).Exec()
	return err
}

func (folder *FolderModel) GetNumberOfFolderByUserId(userId string) (int, error) {
	rows, err := database.SelectHelper(
		squirrel.Select("COUNT(*)").
			From("folders").
			Where(squirrel.Eq{"user_id": userId}),
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

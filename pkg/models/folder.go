package models

import (
	"VulTracks/pkg/database"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type FolderModel struct {
	Id       string `json:"id"`
	Path     string `json:"path" validate:"required"`
	LastScan string `json:"last_scan" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
}

func (folder *FolderModel) CreateTable() error {
	_, err := database.Database.Exec(`
		CREATE TABLE IF NOT EXISTS folders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL UNIQUE,
			last_scan TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func appendFoldersToList(list []FolderModel, rows *sql.Rows) ([]FolderModel, error) {
	var folder FolderModel
	err := rows.Scan(&folder.Id, &folder.Path, &folder.LastScan, &folder.UserId)
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

	err = rows.Scan(&folder.Id, &folder.Path, &folder.LastScan, &folder.UserId)
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
		Columns("path", "last_scan", "user_id").
		Values(folder.Path, folder.LastScan, folder.UserId).
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

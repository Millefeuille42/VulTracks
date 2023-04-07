package interfaces

import "database/sql"

type FolderInterface struct {
	Path string `json:"path" validate:"required"`
}

type CountPerFolderInterface struct {
	Id       string         `json:"folder"`
	Name     string         `json:"name"`
	Path     string         `json:"path"`
	LastScan string         `json:"last_scan"`
	ParentId sql.NullString `json:"parent_id"`
	Count    int            `json:"count"`
}

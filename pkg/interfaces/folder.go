package interfaces

type FolderInterface struct {
	Path string `json:"path" validate:"required"`
}

type CountPerFolderInterface struct {
	Id       string `json:"folder"`
	Path     string `json:"path"`
	LastScan string `json:"last_scan"`
	Count    string `json:"count"`
}

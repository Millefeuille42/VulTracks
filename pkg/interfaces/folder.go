package interfaces

type FolderInterface struct {
	Path string `json:"path" validate:"required"`
}

package interfaces

type AuthInterface struct {
	Password string `json:"password" validate:"required,min=8"`
	Username string `json:"username" validate:"required,min=3,max=20"`
}

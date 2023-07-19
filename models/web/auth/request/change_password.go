package request

type ChangePasswordRequest struct {
	Password         string `json:"password" form:"password" validate:"required,min=8,eqfield=RepeatedPassword"`
	RepeatedPassword string `json:"repeated_password" form:"repeated_password" validate:"required,min=8,eqfield=Password"`
}

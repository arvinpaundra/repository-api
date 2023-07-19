package request

type ForgotPasswordRequest struct {
	Password    string `json:"password" form:"password" validate:"required,min=8"`
	Base64Token string `json:"token" form:"token" validate:"required"`
}

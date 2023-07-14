package request

type SendEmailForgotPasswordRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

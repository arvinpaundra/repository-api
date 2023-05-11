package request

type ConfirmRequest struct {
	Status string `json:"status" form:"status" validate:"required"`
}

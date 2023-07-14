package request

type CreateInternshipReportRequest struct {
	DepartementId string `json:"departement_id" form:"departement_id" validate:"required"`
	CategoryId    string `json:"category_id" form:"category_id" validate:"required"`
	Title         string `json:"title" form:"title" validate:"required"`
	DateValidated string `json:"date_validated" form:"date_validated" validate:"required"`
	Author        string `json:"author" form:"author" validate:"required"`
	Mentor        string `json:"mentor" form:"mentor" validate:"required"`
}

package request

type UpdateResearchReportRequest struct {
	CollectionId  string   `json:"collection_id" form:"collection_id" validate:"required"`
	DepartementId string   `json:"departement_id" form:"departement_id" validate:"required"`
	CategoryId    string   `json:"category_id" form:"category_id" validate:"required"`
	Title         string   `json:"title" form:"title" validate:"required"`
	Abstract      string   `json:"abstract" form:"abstract" validate:"required"`
	Authors       []string `json:"authors" form:"authors" validate:"required"`
	DateValidated string   `json:"date_validated" form:"date_validated" validate:"required"`
	Status        string   `json:"status" form:"status" validate:"required"`
}

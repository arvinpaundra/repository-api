package request

type CreateFinalProjectReportRequest struct {
	DepartementId  string `json:"departement_id" form:"departement_id" validate:"required"`
	CategoryId     string `json:"category_id" form:"category_id" validate:"required"`
	Title          string `json:"title" form:"title" validate:"required"`
	Abstract       string `json:"abstract" form:"abstract" validate:"required"`
	DateValidated  string `json:"date_validated" form:"date_validated" validate:"required"`
	Improvement    string `json:"improvement" form:"improvement" validate:"required"`
	RelatedTitle   string `json:"related_title" form:"related_title"`
	UpdateDesc     string `json:"update_desc" form:"update_desc"`
	Author         string `json:"author" form:"author" validate:"required"`
	FirstMentor    string `json:"first_mentor" form:"first_mentor" validate:"required"`
	SecondMentor   string `json:"second_mentor" form:"second_mentor" validate:"required"`
	FirstExaminer  string `json:"first_examiner" form:"first_examiner" validate:"required"`
	SecondExaminer string `json:"second_examiner" form:"second_examiner" validate:"required"`
}

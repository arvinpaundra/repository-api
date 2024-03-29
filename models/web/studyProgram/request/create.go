package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreateStudyProgramRequest struct {
	DepatementId string `json:"departement_id" form:"departement_id" validate:"required"`
	Name         string `json:"name" form:"name" validate:"required"`
	CoverColor   string `json:"cover_color" form:"cover_color" validate:"required,max=50"`
}

func (req *CreateStudyProgramRequest) ToDomainStudyProgram() domain.StudyProgram {
	return domain.StudyProgram{
		DepartementId: req.DepatementId,
		Name:          req.Name,
		CoverColor:    req.CoverColor,
	}
}

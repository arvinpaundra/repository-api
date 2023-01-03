package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateDepartementRequest struct {
	StudyProgramId string `json:"study_program_id" form:"study_program_id" validate:"required"`
	Name           string `json:"name" form:"name" validate:"required"`
}

func (req *UpdateDepartementRequest) ToDomainDepartement() domain.Departement {
	return domain.Departement{
		StudyProgramId: req.StudyProgramId,
		Name:           req.Name,
	}
}

package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreateDepartementRequest struct {
	StudyProgramId string `json:"study_program_id" form:"study_program_id" validate:"required"`
	Name           string `json:"name" form:"name" validate:"required"`
	Code           string `json:"code" form:"code" validate:"required"`
}

func (req *CreateDepartementRequest) ToDomainDepartement() domain.Departement {
	return domain.Departement{
		StudyProgramId: req.StudyProgramId,
		Name:           req.Name,
		Code:           req.Code,
	}
}

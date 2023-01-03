package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type DepartementResponse struct {
	ID               string    `json:"id"`
	StudyProgramId   string    `json:"study_program_id"`
	Name             string    `json:"name"`
	StudyProgramName string    `json:"study_program"`
	Code             string    `json:"code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func ToDepartementResponse(departementDomain domain.Departement) DepartementResponse {
	return DepartementResponse{
		ID:               departementDomain.ID,
		StudyProgramId:   departementDomain.StudyProgramId,
		Name:             departementDomain.Name,
		StudyProgramName: departementDomain.StudyProgram.Name,
		Code:             departementDomain.Code,
		CreatedAt:        departementDomain.CreatedAt,
		UpdatedAt:        departementDomain.UpdatedAt,
	}
}

func ToDepartementsResponse(departementDomain []domain.Departement) []DepartementResponse {
	var departements []DepartementResponse

	for _, departement := range departementDomain {
		departements = append(departements, ToDepartementResponse(departement))
	}

	return departements
}

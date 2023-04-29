package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type StudyProgramResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Departement string    `json:"departement"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToStudyProgramResponse(studyProgramDomain domain.StudyProgram) StudyProgramResponse {
	return StudyProgramResponse{
		ID:          studyProgramDomain.ID,
		Name:        studyProgramDomain.Name,
		Departement: studyProgramDomain.Departement.Name,
		CreatedAt:   studyProgramDomain.CreatedAt,
		UpdatedAt:   studyProgramDomain.UpdatedAt,
	}
}

func ToStudyProgramsResponse(studyProgramDomain []domain.StudyProgram) []StudyProgramResponse {
	var studyPrograms []StudyProgramResponse

	for _, studyProgram := range studyProgramDomain {
		studyPrograms = append(studyPrograms, ToStudyProgramResponse(studyProgram))
	}

	return studyPrograms
}

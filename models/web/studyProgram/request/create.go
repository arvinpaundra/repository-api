package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreateStudyProgramRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *CreateStudyProgramRequest) ToDomainStudyProgram() domain.StudyProgram {
	return domain.StudyProgram{
		Name: req.Name,
	}
}

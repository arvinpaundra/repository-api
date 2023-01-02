package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateStudyProgramRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *UpdateStudyProgramRequest) ToDomainStudyProgram() domain.StudyProgram {
	return domain.StudyProgram{
		Name: req.Name,
	}
}

package study_program

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/studyProgram/request"
	"github.com/arvinpaundra/repository-api/models/web/studyProgram/response"
)

type StudyProgramService interface {
	Create(ctx context.Context, studyProgram request.CreateStudyProgramRequest) error
	Update(ctx context.Context, studyProgram request.UpdateStudyProgramRequest, studyProgramId string) error
	Delete(ctx context.Context, studyProgramId string) error
	FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.StudyProgramResponse, int, int, error)
	FindById(ctx context.Context, studyProgramId string) (response.StudyProgramResponse, error)
}

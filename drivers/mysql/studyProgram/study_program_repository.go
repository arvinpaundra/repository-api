package study_program

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type StudyProgramRepository interface {
	Save(ctx context.Context, studyProgram domain.StudyProgram) error
	Update(ctx context.Context, studyProgram domain.StudyProgram, studyProgramId string) error
	Delete(ctx context.Context, studyProgramId string) error
	FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.StudyProgram, int, error)
	FindById(ctx context.Context, studyProgramId string) (domain.StudyProgram, error)
}

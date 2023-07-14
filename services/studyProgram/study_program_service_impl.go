package study_program

import (
	"context"
	"math"

	"github.com/arvinpaundra/repository-api/drivers/mysql/departement"
	studyProgram "github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram"
	"github.com/arvinpaundra/repository-api/models/web/studyProgram/request"
	"github.com/arvinpaundra/repository-api/models/web/studyProgram/response"
	"github.com/google/uuid"
)

type StudyProgramServiceImpl struct {
	studyProgramRepository studyProgram.StudyProgramRepository
	departementRepository  departement.DepartementRepository
}

func NewStudyProgramService(
	studyProgramRepository studyProgram.StudyProgramRepository,
	departementRepository departement.DepartementRepository,
) StudyProgramService {
	return StudyProgramServiceImpl{
		studyProgramRepository: studyProgramRepository,
		departementRepository:  departementRepository,
	}
}

func (service StudyProgramServiceImpl) Create(ctx context.Context, studyProgram request.CreateStudyProgramRequest) error {
	if _, err := service.departementRepository.FindById(ctx, studyProgram.DepatementId); err != nil {
		return err
	}

	studyProgramDomain := studyProgram.ToDomainStudyProgram()

	studyProgramDomain.ID = uuid.NewString()

	err := service.studyProgramRepository.Save(ctx, studyProgramDomain)

	if err != nil {
		return err
	}

	return nil
}

func (service StudyProgramServiceImpl) Update(ctx context.Context, studyProgram request.UpdateStudyProgramRequest, studyProgramId string) error {
	if _, err := service.departementRepository.FindById(ctx, studyProgram.DepatementId); err != nil {
		return err
	}

	if _, err := service.studyProgramRepository.FindById(ctx, studyProgramId); err != nil {
		return err
	}

	err := service.studyProgramRepository.Update(ctx, studyProgram.ToDomainStudyProgram(), studyProgramId)

	if err != nil {
		return err
	}

	return nil
}

func (service StudyProgramServiceImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.StudyProgramResponse, int, int, error) {
	studyPrograms, totalRows, err := service.studyProgramRepository.FindAll(ctx, keyword, limit, offset)

	if err != nil {
		return []response.StudyProgramResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToStudyProgramsResponse(studyPrograms), totalRows, int(totalPages), nil
}

func (service StudyProgramServiceImpl) FindById(ctx context.Context, studyProgramId string) (response.StudyProgramResponse, error) {
	studyProgram, err := service.studyProgramRepository.FindById(ctx, studyProgramId)

	if err != nil {
		return response.StudyProgramResponse{}, err
	}

	return response.ToStudyProgramResponse(studyProgram), nil
}

func (service StudyProgramServiceImpl) FindByDepartementId(ctx context.Context, departementId string) ([]response.StudyProgramResponse, error) {
	if _, err := service.departementRepository.FindById(ctx, departementId); err != nil {
		return []response.StudyProgramResponse{}, err
	}

	studyPrograms, err := service.studyProgramRepository.FindByDepartementId(ctx, departementId)

	if err != nil {
		return []response.StudyProgramResponse{}, err
	}

	return response.ToStudyProgramsResponse(studyPrograms), nil
}

func (service StudyProgramServiceImpl) Delete(ctx context.Context, studyProgramId string) error {
	panic("not implemented")
}

package study_program_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/studyProgram/request"
	studyProgram "github.com/arvinpaundra/repository-api/services/studyProgram"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	studyProgramRepository mocks.StudyProgramRepository
	studyProgramService    studyProgram.StudyProgramService

	studyProgramDomain    domain.StudyProgram
	createStudyProgramDTO request.CreateStudyProgramRequest
	updateStudyProgramDTO request.UpdateStudyProgramRequest

	ctx       context.Context
	totalRows int
)

func TestMain(m *testing.M) {
	studyProgramService = studyProgram.NewStudyProgramService(&studyProgramRepository)

	studyProgramDomain = domain.StudyProgram{
		ID:   uuid.NewString(),
		Name: "test",
	}

	createStudyProgramDTO = request.CreateStudyProgramRequest{
		Name: studyProgramDomain.Name,
	}

	updateStudyProgramDTO = request.UpdateStudyProgramRequest{
		Name: studyProgramDomain.Name,
	}

	ctx = context.Background()
	totalRows = 1

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		studyProgramRepository.Mock.On("Save", ctx, mock.Anything).Return(nil).Once()

		err := studyProgramService.Create(ctx, createStudyProgramDTO)

		assert.NoError(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		studyProgramRepository.Mock.On("Save", ctx, mock.Anything).Return(errors.New("error occurred")).Once()

		err := studyProgramService.Create(ctx, createStudyProgramDTO)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		studyProgramRepository.Mock.On("Update", ctx, updateStudyProgramDTO.ToDomainStudyProgram(), studyProgramDomain.ID).Return(nil).Once()

		err := studyProgramService.Update(ctx, updateStudyProgramDTO, studyProgramDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Failed | Study program not found", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(domain.StudyProgram{}, utils.ErrStudyProgramNotFound).Once()

		err := studyProgramService.Update(ctx, updateStudyProgramDTO, studyProgramDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		studyProgramRepository.Mock.On("Update", ctx, updateStudyProgramDTO.ToDomainStudyProgram(), studyProgramDomain.ID).Return(errors.New("error occurred")).Once()

		err := studyProgramService.Update(ctx, updateStudyProgramDTO, studyProgramDomain.ID)

		assert.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.StudyProgram{studyProgramDomain}, totalRows, nil).Once()

		results, actualTotalRows, actualTotalPages, err := studyProgramService.FindAll(ctx, "", 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.NotEmpty(t, actualTotalRows)
		assert.NotEmpty(t, actualTotalPages)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.StudyProgram{}, 0, errors.New("error occurred")).Once()

		results, actualTotalRows, actualTotalPages, err := studyProgramService.FindAll(ctx, "", 10, 0)

		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Empty(t, actualTotalRows)
		assert.Empty(t, actualTotalPages)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		result, err := studyProgramService.FindById(ctx, studyProgramDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Failed | Study program not found", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(domain.StudyProgram{}, utils.ErrStudyProgramNotFound).Once()

		result, err := studyProgramService.FindById(ctx, studyProgramDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

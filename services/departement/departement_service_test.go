package departement_test

import (
	"context"
	"errors"
	"testing"

	deptMock "github.com/arvinpaundra/repository-api/drivers/mysql/departement/mocks"
	stuProMock "github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/departement/request"
	"github.com/arvinpaundra/repository-api/services/departement"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	departementRepository  deptMock.DepartementRepository
	studyProgramRepository stuProMock.StudyProgramRepository
	departementService     departement.DepartementService

	departementDomain    domain.Departement
	studyProgramDomain   domain.StudyProgram
	createDepartementDTO request.CreateDepartementRequest
	updateDepartementDTO request.UpdateDepartementRequest

	ctx       context.Context
	totalRows int
)

func TestMain(m *testing.M) {
	departementService = departement.NewDepartementService(&departementRepository, &studyProgramRepository)

	studyProgramDomain = domain.StudyProgram{
		ID:   uuid.NewString(),
		Name: "test",
	}

	departementDomain = domain.Departement{
		ID:             uuid.NewString(),
		StudyProgramId: studyProgramDomain.ID,
		Name:           "test",
		Code:           "test",
	}

	createDepartementDTO = request.CreateDepartementRequest{
		StudyProgramId: studyProgramDomain.ID,
		Name:           departementDomain.Name,
		Code:           departementDomain.Code,
	}

	updateDepartementDTO = request.UpdateDepartementRequest{
		StudyProgramId: studyProgramDomain.ID,
		Name:           departementDomain.Name,
	}

	ctx = context.Background()
	totalRows = 1

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		departementRepository.Mock.On("Save", ctx, mock.Anything).Return(nil).Once()

		err := departementService.Create(ctx, createDepartementDTO)

		assert.NoError(t, err)
	})

	t.Run("Failed | Study program not found", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(domain.StudyProgram{}, utils.ErrStudyProgramNotFound).Once()

		err := departementService.Create(ctx, createDepartementDTO)

		assert.Error(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		departementRepository.Mock.On("Save", ctx, mock.Anything).Return(errors.New("error occurred")).Once()

		err := departementService.Create(ctx, createDepartementDTO)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		departementRepository.Mock.On("FindById", ctx, departementDomain.ID).Return(departementDomain, nil).Once()

		departementRepository.Mock.On("Update", ctx, updateDepartementDTO.ToDomainDepartement(), departementDomain.ID).Return(nil).Once()

		err := departementService.Update(ctx, updateDepartementDTO, departementDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Failed | Study program not found", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(domain.StudyProgram{}, utils.ErrStudyProgramNotFound).Once()

		err := departementService.Update(ctx, updateDepartementDTO, departementDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Departement not found", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		departementRepository.Mock.On("FindById", ctx, departementDomain.ID).Return(domain.Departement{}, utils.ErrDepartementNotFound).Once()

		err := departementService.Update(ctx, updateDepartementDTO, departementDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		departementRepository.Mock.On("FindById", ctx, departementDomain.ID).Return(departementDomain, nil).Once()

		departementRepository.Mock.On("Update", ctx, updateDepartementDTO.ToDomainDepartement(), departementDomain.ID).Return(errors.New("error occurred")).Once()

		err := departementService.Update(ctx, updateDepartementDTO, departementDomain.ID)

		assert.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		departementRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.Departement{departementDomain}, totalRows, nil).Once()

		results, actualTotalRows, actualTotalPages, err := departementService.FindAll(ctx, "", 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.NotEmpty(t, actualTotalRows)
		assert.NotEmpty(t, actualTotalPages)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		departementRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.Departement{}, 0, errors.New("error occurred")).Once()

		results, actualTotalRows, actualTotalPages, err := departementService.FindAll(ctx, "", 10, 0)

		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Empty(t, actualTotalRows)
		assert.Empty(t, actualTotalPages)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		departementRepository.Mock.On("FindById", ctx, departementDomain.ID).Return(departementDomain, nil).Once()

		result, err := departementService.FindById(ctx, departementDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Failed | Departement not found", func(t *testing.T) {
		departementRepository.Mock.On("FindById", ctx, departementDomain.ID).Return(domain.Departement{}, utils.ErrDepartementNotFound).Once()

		result, err := departementService.FindById(ctx, departementDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

func TestFindByProgramStudyId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		departementRepository.Mock.On("FindByProgramStudyId", ctx, studyProgramDomain.ID).Return([]domain.Departement{departementDomain}, nil).Once()

		results, err := departementService.FindByProgramStudyId(ctx, studyProgramDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, results)
	})

	t.Run("Failed | Study program not found", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(domain.StudyProgram{}, utils.ErrStudyProgramNotFound).Once()

		results, err := departementService.FindByProgramStudyId(ctx, studyProgramDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, results)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		studyProgramRepository.Mock.On("FindById", ctx, studyProgramDomain.ID).Return(studyProgramDomain, nil).Once()

		departementRepository.Mock.On("FindByProgramStudyId", ctx, studyProgramDomain.ID).Return([]domain.Departement{}, errors.New("error occurred")).Once()

		results, err := departementService.FindByProgramStudyId(ctx, studyProgramDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, results)
	})
}

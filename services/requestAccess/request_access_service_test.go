package request_access_test

import (
	"context"
	"errors"
	"testing"

	pemustakaMock "github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka/mocks"
	reqAccessMock "github.com/arvinpaundra/repository-api/drivers/mysql/requestAccess/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/requestAccess/request"
	requestAccess "github.com/arvinpaundra/repository-api/services/requestAccess"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	requestAccessRepository reqAccessMock.RequestAccessRepository
	pemustakaRepository     pemustakaMock.PemustakaRepository
	requestAccessService    requestAccess.RequestAccessService

	requestAccessDomain domain.RequestAccess
	pemustakaDomain     domain.Pemustaka

	updateRequestAccessDTO request.UpdateRequestAccessRequest

	ctx    context.Context
	tx     *gorm.DB
	limit  int
	offset int
)

func TestMain(m *testing.M) {
	requestAccessService = requestAccess.NewRequestAccessService(&requestAccessRepository, &pemustakaRepository, tx)

	pemustakaDomain = domain.Pemustaka{
		ID:                      uuid.NewString(),
		UserId:                  uuid.NewString(),
		StudyProgramId:          uuid.NewString(),
		DepartementId:           uuid.NewString(),
		RoleId:                  uuid.NewString(),
		MemberCode:              "test",
		Fullname:                "test",
		IdentityNumber:          "test",
		YearGen:                 "test",
		Gender:                  "test",
		Telp:                    "test",
		BirthDate:               "test",
		Address:                 "test",
		IsCollectedFinalProject: "0",
		IsActive:                "0",
		Avatar:                  "test",
	}

	requestAccessDomain = domain.RequestAccess{
		ID:              uuid.NewString(),
		PemustakaId:     pemustakaDomain.ID,
		SupportEvidence: "test",
		Status:          "test",
	}

	updateRequestAccessDTO = request.UpdateRequestAccessRequest{
		Status: "test",
	}

	ctx = context.Background()
	limit = 10
	offset = 0

	m.Run()
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		requestAccessRepository.Mock.On("FindAll", ctx, "", "", limit, offset).Return([]domain.RequestAccess{requestAccessDomain}, 1, nil).Once()

		requestAccesses, totalRows, totalPages, err := requestAccessService.FindAll(ctx, "", "", limit, offset)

		assert.NoError(t, err)
		assert.NotEmpty(t, requestAccesses)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		requestAccessRepository.Mock.On("FindAll", ctx, "", "", limit, offset).Return([]domain.RequestAccess{}, 0, errors.New("error occurred")).Once()

		requestAccesses, totalRows, totalPages, err := requestAccessService.FindAll(ctx, "", "", limit, offset)

		assert.Error(t, err)
		assert.Empty(t, requestAccesses)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		requestAccessRepository.Mock.On("FindById", ctx, requestAccessDomain.ID).Return(requestAccessDomain, nil).Once()

		requestAccess, err := requestAccessService.FindById(ctx, requestAccessDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, requestAccess)
	})

	t.Run("Failed | Request access not found", func(t *testing.T) {
		requestAccessRepository.Mock.On("FindById", ctx, requestAccessDomain.ID).Return(domain.RequestAccess{}, utils.ErrRequestAccessNotFound).Once()

		requestAccess, err := requestAccessService.FindById(ctx, requestAccessDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, requestAccess)
	})
}

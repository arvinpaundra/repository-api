package pemustaka_test

import (
	"context"
	"errors"
	"testing"

	deptMock "github.com/arvinpaundra/repository-api/drivers/mysql/departement/mocks"
	pemustakaMock "github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka/mocks"
	requestAccessMock "github.com/arvinpaundra/repository-api/drivers/mysql/requestAccess/mocks"
	roleMock "github.com/arvinpaundra/repository-api/drivers/mysql/role/mocks"
	stuPrdMock "github.com/arvinpaundra/repository-api/drivers/mysql/studyProgram/mocks"
	userMock "github.com/arvinpaundra/repository-api/drivers/mysql/user/mocks"
	cloudinaryMock "github.com/arvinpaundra/repository-api/helper/cloudinary/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/request"
	"github.com/arvinpaundra/repository-api/services/pemustaka"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	userRepository          userMock.UserRepository
	pemustakaRepository     pemustakaMock.PemustakaRepository
	stuProdRepository       stuPrdMock.StudyProgramRepository
	departementRepository   deptMock.DepartementRepository
	roleRepository          roleMock.RoleRepository
	requestAccessRepository requestAccessMock.RequestAccessRepository
	cloudinary              cloudinaryMock.Cloudinary
	pemustakaService        pemustaka.PemustakaService

	userDomain         domain.User
	pemustakaDomain    domain.Pemustaka
	studyProgramDomain domain.StudyProgram
	departementDomain  domain.Departement
	roleDomain         domain.Role

	registerPemustakaDTO  request.RegisterPemustakaRequest
	loginPemustakaDTO     request.LoginPemustakaRequest
	updatePemustakaDTO    request.UpdatePemustakaRequest
	pemustakaRequestQuery request.PemustakaRequestQuery

	ctx    context.Context
	tx     *gorm.DB
	limit  int
	offset int
)

func TestMain(m *testing.M) {
	pemustakaService = pemustaka.NewPemustakaService(
		&userRepository,
		&pemustakaRepository,
		&stuProdRepository,
		&departementRepository,
		&roleRepository,
		&requestAccessRepository,
		&cloudinary,
		tx,
	)

	userDomain = domain.User{
		ID:       uuid.NewString(),
		Email:    "test@mail.com",
		Password: "12345678",
	}

	studyProgramDomain = domain.StudyProgram{
		ID:   uuid.NewString(),
		Name: "test",
	}

	departementDomain = domain.Departement{
		ID:   uuid.NewString(),
		Name: "test",
		Code: "test",
	}

	roleDomain = domain.Role{
		ID:         uuid.NewString(),
		Role:       "test",
		Visibility: "all",
	}

	pemustakaDomain = domain.Pemustaka{
		ID:                      uuid.NewString(),
		UserId:                  userDomain.ID,
		StudyProgramId:          studyProgramDomain.ID,
		DepartementId:           departementDomain.ID,
		RoleId:                  roleDomain.ID,
		MemberCode:              "test",
		Fullname:                "test",
		IdentityNumber:          "test",
		YearGen:                 "test",
		Gender:                  "test",
		Telp:                    "test",
		BirthDate:               "test",
		Address:                 "test",
		IsCollectedFinalProject: "1",
		IsActive:                "1",
		Avatar:                  "test",
	}

	registerPemustakaDTO = request.RegisterPemustakaRequest{
		Email:          userDomain.Email,
		Password:       userDomain.Password,
		StudyProgramId: studyProgramDomain.ID,
		DepartementId:  departementDomain.ID,
		RoleId:         roleDomain.ID,
		Fullname:       pemustakaDomain.Fullname,
		IdentityNumber: pemustakaDomain.IdentityNumber,
		YearGen:        pemustakaDomain.YearGen,
	}

	loginPemustakaDTO = request.LoginPemustakaRequest{
		Email:    userDomain.Email,
		Password: userDomain.Password,
	}

	updatePemustakaDTO = request.UpdatePemustakaRequest{
		StudyProgramId: studyProgramDomain.ID,
		DepartementId:  departementDomain.ID,
		Fullname:       pemustakaDomain.Fullname,
		YearGen:        pemustakaDomain.YearGen,
		Gender:         pemustakaDomain.Gender,
		Telp:           pemustakaDomain.Telp,
		BirthDate:      pemustakaDomain.BirthDate,
		Address:        pemustakaDomain.Address,
	}

	pemustakaRequestQuery = request.PemustakaRequestQuery{
		Keyword:                 "",
		RoleId:                  roleDomain.ID,
		DepartementId:           "",
		IsCollectedFinalProject: "",
		YearGen:                 "",
	}

	ctx = context.Background()
	limit = 10
	offset = 0

	m.Run()
}

func TestLogin(t *testing.T) {
	t.Run("Failed | User email not found", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", ctx, loginPemustakaDTO.Email).Return(domain.User{}, utils.ErrEmailNotFound).Once()

		token, err := pemustakaService.Login(ctx, loginPemustakaDTO)

		assert.Error(t, err)
		assert.Empty(t, token)
	})

	t.Run("Failed | Pemustaka not found", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", ctx, loginPemustakaDTO.Email).Return(userDomain, nil).Once()

		pemustakaRepository.Mock.On("FindByUserId", ctx, userDomain.ID).Return(pemustakaDomain, utils.ErrPemustakaNotFound).Once()

		token, err := pemustakaService.Login(ctx, loginPemustakaDTO)

		assert.Error(t, err)
		assert.Empty(t, token)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindAll", ctx, pemustakaRequestQuery, limit, offset).Return([]domain.Pemustaka{pemustakaDomain}, 1, nil).Once()

		results, totalRows, totalPages, err := pemustakaService.FindAll(ctx, pemustakaRequestQuery, limit, offset)

		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindAll", ctx, pemustakaRequestQuery, limit, offset).Return([]domain.Pemustaka{}, 0, errors.New("error occurred")).Once()

		results, totalRows, totalPages, err := pemustakaService.FindAll(ctx, pemustakaRequestQuery, limit, offset)

		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		result, err := pemustakaService.FindById(ctx, pemustakaDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Success", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(domain.Pemustaka{}, utils.ErrPemustakaNotFound).Once()

		result, err := pemustakaService.FindById(ctx, pemustakaDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

package author_test

import (
	"context"
	"errors"
	"testing"

	authorMock "github.com/arvinpaundra/repository-api/drivers/mysql/author/mocks"
	pemustakaMock "github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka/mocks"
	repositoryMock "github.com/arvinpaundra/repository-api/drivers/mysql/repository/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/services/author"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	authorRepository    authorMock.AuthorRepository
	repository          repositoryMock.Repository
	pemustakaRepository pemustakaMock.PemustakaRepository
	authorService       author.AuthorService

	pemustakaDomain  domain.Pemustaka
	repositoryDomain domain.Repository
	authorDomain     domain.Author

	ctx context.Context
)

func TestMain(m *testing.M) {
	authorService = author.NewAuthorService(&authorRepository, &repository, &pemustakaRepository)

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
		IsCollectedFinalProject: "1",
		IsActive:                "1",
		Avatar:                  "test",
	}

	repositoryDomain := domain.Repository{
		ID:            uuid.NewString(),
		CollectionId:  uuid.NewString(),
		Title:         "test",
		Abstract:      "test",
		Improvement:   "0",
		RelatedTitle:  "",
		UpdateDesc:    "",
		DateValidated: "",
		Status:        "pending",
	}

	authorDomain = domain.Author{
		ID:           uuid.NewString(),
		RepositoryId: repositoryDomain.ID,
		PemustakaId:  pemustakaDomain.ID,
	}

	ctx = context.Background()

	m.Run()
}

func TestDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		authorRepository.Mock.On("Delete", ctx, repositoryDomain.ID, pemustakaDomain.ID).Return(nil).Once()

		err := authorService.Delete(ctx, repositoryDomain.ID, pemustakaDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Failed | Repository not found", func(t *testing.T) {
		repository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(domain.Repository{}, utils.ErrRepositoryNotFound).Once()

		err := authorService.Delete(ctx, repositoryDomain.ID, pemustakaDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Pemustaka not found", func(t *testing.T) {
		repository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(domain.Pemustaka{}, utils.ErrPemustakaNotFound).Once()

		err := authorService.Delete(ctx, repositoryDomain.ID, pemustakaDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Faailed | Error occurred", func(t *testing.T) {
		repository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		authorRepository.Mock.On("Delete", ctx, repositoryDomain.ID, pemustakaDomain.ID).Return(errors.New("error occurred")).Once()

		err := authorService.Delete(ctx, repositoryDomain.ID, pemustakaDomain.ID)

		assert.Error(t, err)
	})
}

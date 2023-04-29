package repository_test

import (
	"context"
	"errors"
	"testing"

	authorMock "github.com/arvinpaundra/repository-api/drivers/mysql/author/mocks"
	collectionMock "github.com/arvinpaundra/repository-api/drivers/mysql/collection/mocks"
	contributorMock "github.com/arvinpaundra/repository-api/drivers/mysql/contributor/mocks"
	departementMock "github.com/arvinpaundra/repository-api/drivers/mysql/departement/mocks"
	documentMock "github.com/arvinpaundra/repository-api/drivers/mysql/document/mocks"
	pemustakaMock "github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka/mocks"
	repositoryMock "github.com/arvinpaundra/repository-api/drivers/mysql/repository/mocks"
	cloudinaryMock "github.com/arvinpaundra/repository-api/helper/cloudinary/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/repository/request"
	"github.com/arvinpaundra/repository-api/services/repository"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	repoRepository        repositoryMock.Repository
	collectionRepository  collectionMock.CollectionRepository
	departementRepository departementMock.DepartementRepository
	pemustakaRepository   pemustakaMock.PemustakaRepository
	authorRepository      authorMock.AuthorRepository
	contributorRepository contributorMock.ContributorRepository
	documentRepository    documentMock.DocumentRepository
	cloudinary            cloudinaryMock.Cloudinary
	repositoryService     repository.RepositoryService

	repositoryDomain  domain.Repository
	collectionDomain  domain.Collection
	departementDomain domain.Departement
	categoryDomain    domain.Category
	pemustakaDomain   domain.Pemustaka
	authorDomain      domain.Author
	contributorDomain domain.Contributor
	documentDomain    domain.Document

	query request.RepositoryRequestQuery

	ctx context.Context
	tx  *gorm.DB
)

func TestMain(m *testing.M) {
	repositoryService = repository.NewRepositoryService(
		&collectionRepository,
		&departementRepository,
		&pemustakaRepository,
		&authorRepository,
		&contributorRepository,
		&repoRepository,
		&documentRepository,
		&cloudinary,
		tx,
	)

	collectionDomain = domain.Collection{
		ID:   uuid.NewString(),
		Name: "test",
	}

	departementDomain = domain.Departement{
		ID:   uuid.NewString(),
		Name: "test",
		Code: "test",
	}

	categoryDomain = domain.Category{
		ID:   uuid.NewString(),
		Name: "test",
	}

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

	repositoryDomain = domain.Repository{
		ID:            uuid.NewString(),
		CollectionId:  collectionDomain.ID,
		DepartementId: departementDomain.ID,
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
	}

	contributorDomain = domain.Contributor{
		ID:            uuid.NewString(),
		RepositoryId:  repositoryDomain.ID,
		PemustakaId:   pemustakaDomain.ID,
		ContributedAs: "test",
	}

	documentDomain = domain.Document{
		ID:                  uuid.NewString(),
		RepositoryId:        repositoryDomain.ID,
		ValidityPage:        "test",
		CoverAndListContent: "test",
		ChpOne:              "test",
		ChpTwo:              "test",
		ChpThree:            "test",
		ChpFour:             "test",
		ChpFive:             "test",
		Bibliography:        "test",
	}

	query = request.RepositoryRequestQuery{
		Keyword:       "",
		CollectionId:  "",
		DepartementId: "",
		Improvement:   "",
		Status:        "",
		Sort:          "created_at DESC",
	}

	ctx = context.Background()

	m.Run()
}

func TestDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpOne).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpTwo).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpThree).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpFour).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpFive).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.Bibliography).Return(nil).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Failed | Repository not found", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(domain.Repository{}, utils.ErrRepositoryNotFound).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Document not found", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(domain.Document{}, errors.New("Document not found")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete from db", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(errors.New("delete error")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete validity_page", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(errors.New("error delete validity page")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete cover and list content", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(errors.New("error delete cover and list content")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete chp one", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(errors.New("error delete cover and list content")).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpOne).Return(errors.New("error delete chp one")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete chp two", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpOne).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpTwo).Return(errors.New("error delete chp two")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete chp three", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpOne).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpThree).Return(errors.New("error delete chp three")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete chp four", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpOne).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpThree).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpFour).Return(errors.New("error delete chp four")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete chp five", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpOne).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpThree).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpFour).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpFive).Return(errors.New("error delete chp five")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error delete chp bibliography", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repoRepository.Mock.On("Delete", ctx, repositoryDomain.ID).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ValidityPage).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.CoverAndListContent).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpOne).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpThree).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpFour).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.ChpFive).Return(nil).Once()

		cloudinary.Mock.On("Delete", ctx, documentDomain.Bibliography).Return(errors.New("error delete chp bibliography")).Once()

		err := repositoryService.Delete(ctx, repositoryDomain.ID)

		assert.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repoRepository.Mock.On("FindAll", ctx, query, 10, 0).Return([]domain.Repository{repositoryDomain}, 1, nil).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindAll(ctx, query, 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, repositories)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Success | Error find all", func(t *testing.T) {
		repoRepository.Mock.On("FindAll", ctx, query, 10, 0).Return([]domain.Repository{}, 0, errors.New("error occurred")).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindAll(ctx, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		authorRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return([]domain.Author{authorDomain}, nil).Once()

		contributorRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return([]domain.Contributor{contributorDomain}, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(documentDomain, nil).Once()

		repository, err := repositoryService.FindById(ctx, repositoryDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, repository)
	})

	t.Run("Failed | Repository not found", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(domain.Repository{}, utils.ErrRepositoryNotFound).Once()

		repository, err := repositoryService.FindById(ctx, repositoryDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, repository)
	})

	t.Run("Failed | Authors not found", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		authorRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return([]domain.Author{}, errors.New("authors not found")).Once()

		repository, err := repositoryService.FindById(ctx, repositoryDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, repository)
	})

	t.Run("Failed | Contributors not found", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		authorRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return([]domain.Author{authorDomain}, nil).Once()

		contributorRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return([]domain.Contributor{}, errors.New("contributors not found")).Once()

		repository, err := repositoryService.FindById(ctx, repositoryDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, repository)
	})

	t.Run("Failed | Documents not found", func(t *testing.T) {
		repoRepository.Mock.On("FindById", ctx, repositoryDomain.ID).Return(repositoryDomain, nil).Once()

		authorRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return([]domain.Author{authorDomain}, nil).Once()

		contributorRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return([]domain.Contributor{contributorDomain}, nil).Once()

		documentRepository.Mock.On("FindByRepositoryId", ctx, repositoryDomain.ID).Return(domain.Document{}, errors.New("documents not found")).Once()

		repository, err := repositoryService.FindById(ctx, repositoryDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, repository)
	})
}

func TestFindByAuthorId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		repoRepository.Mock.On("FindByAuthorId", ctx, pemustakaDomain.ID, query, 10, 0).Return([]domain.Repository{repositoryDomain}, 1, nil).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByAuthorId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, repositories)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Failed | Author not found", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(domain.Pemustaka{}, utils.ErrPemustakaNotFound).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByAuthorId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})

	t.Run("Failed | Error find by author_id", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		repoRepository.Mock.On("FindByAuthorId", ctx, pemustakaDomain.ID, query, 10, 0).Return([]domain.Repository{}, 0, errors.New("error occurred")).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByAuthorId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

func TestFindByMentorId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		repoRepository.Mock.On("FindByMentorId", ctx, pemustakaDomain.ID, query, 10, 0).Return([]domain.Repository{repositoryDomain}, 1, nil).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByMentorId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, repositories)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Failed | Mentor not found", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(domain.Pemustaka{}, utils.ErrPemustakaNotFound).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByMentorId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})

	t.Run("Failed | Error find by mentor_id", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		repoRepository.Mock.On("FindByMentorId", ctx, pemustakaDomain.ID, query, 10, 0).Return([]domain.Repository{}, 0, errors.New("error occurred")).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByMentorId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

func TestFindByExaminerId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		repoRepository.Mock.On("FindByExaminerId", ctx, pemustakaDomain.ID, query, 10, 0).Return([]domain.Repository{repositoryDomain}, 1, nil).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByExaminerId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, repositories)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Failed | Examiner not found", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(domain.Pemustaka{}, utils.ErrPemustakaNotFound).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByExaminerId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})

	t.Run("Failed | Error find by examiner_id", func(t *testing.T) {
		pemustakaRepository.Mock.On("FindById", ctx, pemustakaDomain.ID).Return(pemustakaDomain, nil).Once()

		repoRepository.Mock.On("FindByExaminerId", ctx, pemustakaDomain.ID, query, 10, 0).Return([]domain.Repository{}, 0, errors.New("error occurred")).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByExaminerId(ctx, pemustakaDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

func TestFindByCollectionId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(collectionDomain, nil).Once()

		repoRepository.Mock.On("FindByCollectionId", ctx, collectionDomain.ID, query, 10, 0).Return([]domain.Repository{repositoryDomain}, 1, nil).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByCollectionId(ctx, collectionDomain.ID, query, 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, repositories)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Failed | Collection not found", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(domain.Collection{}, utils.ErrCollectionNotFound).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByCollectionId(ctx, collectionDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})

	t.Run("Failed | Error find by collection_id", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(collectionDomain, nil).Once()

		repoRepository.Mock.On("FindByCollectionId", ctx, collectionDomain.ID, query, 10, 0).Return([]domain.Repository{}, 0, errors.New("error occurred")).Once()

		repositories, totalRows, totalPages, err := repositoryService.FindByCollectionId(ctx, collectionDomain.ID, query, 10, 0)

		assert.Error(t, err)
		assert.Empty(t, repositories)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

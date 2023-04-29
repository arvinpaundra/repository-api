package collection_test

import (
	"context"
	"errors"
	"testing"

	collectionMock "github.com/arvinpaundra/repository-api/drivers/mysql/collection/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/collection/request"
	"github.com/arvinpaundra/repository-api/services/collection"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	collectionRepository collectionMock.CollectionRepository
	collectionService    collection.CollectionService

	collectionDomain    domain.Collection
	createCollectionDTO request.CreateCollectionRequest
	updateCollectionDTO request.UpdateCollectionRequest

	ctx context.Context
)

func TestMain(m *testing.M) {
	collectionService = collection.NewCollectionService(&collectionRepository)

	collectionDomain = domain.Collection{
		ID:   uuid.NewString(),
		Name: "test",
	}

	createCollectionDTO = request.CreateCollectionRequest{
		Name: collectionDomain.Name,
	}

	updateCollectionDTO = request.UpdateCollectionRequest{
		Name: collectionDomain.Name,
	}

	ctx = context.Background()

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		collectionRepository.Mock.On("Save", ctx, mock.Anything).Return(nil).Once()

		err := collectionService.Create(ctx, createCollectionDTO)

		assert.NoError(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		collectionRepository.Mock.On("Save", ctx, mock.Anything).Return(errors.New("error occurred")).Once()

		err := collectionService.Create(ctx, createCollectionDTO)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(collectionDomain, nil).Once()

		collectionRepository.Mock.On("Update", ctx, mock.Anything, collectionDomain.ID).Return(nil).Once()

		err := collectionService.Update(ctx, updateCollectionDTO, collectionDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Failed | Collection not found", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(domain.Collection{}, utils.ErrCollectionNotFound).Once()

		err := collectionService.Update(ctx, updateCollectionDTO, collectionDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(collectionDomain, nil).Once()

		collectionRepository.Mock.On("Update", ctx, mock.Anything, collectionDomain.ID).Return(errors.New("error occurred")).Once()

		err := collectionService.Update(ctx, updateCollectionDTO, collectionDomain.ID)

		assert.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		collectionRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.Collection{collectionDomain}, 1, nil).Once()

		collections, totalRows, totalPages, err := collectionService.FindAll(ctx, "", "all", 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, collections)
		assert.NotEmpty(t, totalRows)
		assert.NotEmpty(t, totalPages)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		collectionRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.Collection{}, 0, errors.New("error occurred")).Once()

		collections, totalRows, totalPages, err := collectionService.FindAll(ctx, "", "all", 10, 0)

		assert.Error(t, err)
		assert.Empty(t, collections)
		assert.Empty(t, totalRows)
		assert.Empty(t, totalPages)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(collectionDomain, nil).Once()

		collection, err := collectionService.FindById(ctx, collectionDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, collection)
	})

	t.Run("Failed | Collection not found", func(t *testing.T) {
		collectionRepository.Mock.On("FindById", ctx, collectionDomain.ID).Return(domain.Collection{}, utils.ErrCollectionNotFound).Once()

		collection, err := collectionService.FindById(ctx, collectionDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, collection)
	})
}

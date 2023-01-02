package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arvinpaundra/repository-api/drivers/mysql/category/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/category/request"
	"github.com/arvinpaundra/repository-api/models/web/category/response"
	"github.com/arvinpaundra/repository-api/services/category"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository mocks.CategoryRepository
	categoryService    category.CategoryService

	categoryDomain      domain.Category
	createCategoryDTO   request.CreateCategoryRequest
	updateCategoryDTO   request.UpdateCategoryRequest
	categoryResponseDTO response.CategoryResponse

	ctx       context.Context
	totalRows int64
)

func TestMain(m *testing.M) {
	categoryService = category.NewCategoryService(&categoryRepository)

	categoryDomain = domain.Category{
		ID:   uuid.NewString(),
		Name: "test",
	}

	categoryResponseDTO = response.CategoryResponse{
		ID:   categoryDomain.ID,
		Name: categoryDomain.Name,
	}

	createCategoryDTO = request.CreateCategoryRequest{
		Name: "test",
	}

	updateCategoryDTO = request.UpdateCategoryRequest{
		Name: "test",
	}

	ctx = context.Background()
	totalRows = 1

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		categoryRepository.Mock.On("Save", ctx, mock.Anything).Return(nil).Once()

		err := categoryService.Create(ctx, createCategoryDTO)

		assert.NoError(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		categoryRepository.Mock.On("Save", ctx, mock.Anything).Return(errors.New("error occurred")).Once()

		err := categoryService.Create(ctx, createCategoryDTO)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", ctx, categoryDomain.ID).Return(categoryDomain, nil).Once()

		categoryRepository.Mock.On("Update", ctx, updateCategoryDTO.ToDomainCategory(), categoryDomain.ID).Return(nil).Once()

		err := categoryService.Update(ctx, updateCategoryDTO, categoryDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Failed | Category not found", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", ctx, categoryDomain.ID).Return(domain.Category{}, utils.ErrCategoryNotFound).Once()

		err := categoryService.Update(ctx, updateCategoryDTO, categoryDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", ctx, categoryDomain.ID).Return(categoryDomain, nil).Once()

		categoryRepository.Mock.On("Update", ctx, updateCategoryDTO.ToDomainCategory(), categoryDomain.ID).Return(errors.New("error occurred")).Once()

		err := categoryService.Update(ctx, updateCategoryDTO, categoryDomain.ID)

		assert.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		categoryRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.Category{categoryDomain}, totalRows, nil).Once()

		results, actualTotalRows, actualTotalPages, err := categoryService.FindAll(ctx, "", 10, 0)

		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.NotEmpty(t, actualTotalRows)
		assert.NotEmpty(t, actualTotalPages)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		categoryRepository.Mock.On("FindAll", ctx, "", 10, 0).Return([]domain.Category{}, int64(0), errors.New("error occurred")).Once()

		results, actualTotalRows, actualTotalPages, err := categoryService.FindAll(ctx, "", 10, 0)

		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Empty(t, actualTotalRows)
		assert.Empty(t, actualTotalPages)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", ctx, categoryDomain.ID).Return(categoryDomain, nil).Once()

		result, err := categoryService.FindById(ctx, categoryDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Failed | Category not found", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", ctx, categoryDomain.ID).Return(domain.Category{}, utils.ErrCategoryNotFound).Once()

		result, err := categoryService.FindById(ctx, categoryDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

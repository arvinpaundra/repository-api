package category

import (
	"context"
	"math"

	"github.com/arvinpaundra/repository-api/drivers/mysql/category"
	"github.com/arvinpaundra/repository-api/models/web/category/request"
	"github.com/arvinpaundra/repository-api/models/web/category/response"
	"github.com/google/uuid"
)

type CategoryServiceImpl struct {
	categoryRepository category.CategoryRepository
}

func NewCategoryService(categoryRepository category.CategoryRepository) CategoryService {
	return CategoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}

func (service CategoryServiceImpl) Create(ctx context.Context, category request.CreateCategoryRequest) error {
	domainCategory := category.ToDomainCategory()

	domainCategory.ID = uuid.NewString()

	err := service.categoryRepository.Save(ctx, domainCategory)

	if err != nil {
		return err
	}

	return nil
}

func (service CategoryServiceImpl) Update(ctx context.Context, category request.UpdateCategoryRequest, categoryId string) error {
	if _, err := service.categoryRepository.FindById(ctx, categoryId); err != nil {
		return err
	}

	err := service.categoryRepository.Update(ctx, category.ToDomainCategory(), categoryId)

	if err != nil {
		return err
	}

	return nil
}

func (service CategoryServiceImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.CategoryResponse, int, int, error) {
	categories, totalRows, err := service.categoryRepository.FindAll(ctx, keyword, limit, offset)

	if err != nil {
		return []response.CategoryResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToCategoriesResponse(categories), int(totalRows), int(totalPages), nil
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId string) (response.CategoryResponse, error) {
	category, err := service.categoryRepository.FindById(ctx, categoryId)

	if err != nil {
		return response.CategoryResponse{}, err
	}

	return response.ToCategoryResponse(category), nil
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId string) error {
	panic("not implemented")
}

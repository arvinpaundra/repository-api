package service

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/category/request"
	"github.com/arvinpaundra/repository-api/models/web/category/response"
)

type CategoryService interface {
	Create(ctx context.Context, category request.CreateCategoryRequest) error
	Update(ctx context.Context, category request.UpdateCategoryRequest, categoryId string) error
	Delete(ctx context.Context, categoryId string) error
	FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.CategoryResponse, int, int, error)
	FindById(ctx context.Context, categoryId string) (response.CategoryResponse, error)
}

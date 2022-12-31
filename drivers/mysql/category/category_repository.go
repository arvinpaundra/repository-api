package category

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, category domain.Category) error
	Update(ctx context.Context, category domain.Category, categoryId string) error
	Delete(ctx context.Context, categoryId string) error
	FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.Category, int64, error)
	FindById(ctx context.Context, categoryId string) (domain.Category, error)
}

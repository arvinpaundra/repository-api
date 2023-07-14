package collection

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type CollectionRepository interface {
	Save(ctx context.Context, collection domain.Collection) error
	Update(ctx context.Context, collection domain.Collection, collectionId string) error
	Delete(ctx context.Context, collectionId string) error
	FindAll(ctx context.Context, keyword string, visibility string, limit int, offset int) ([]domain.Collection, int, error)
	FindById(ctx context.Context, collectionId string) (domain.Collection, error)
}

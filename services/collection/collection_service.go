package collection

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/collection/request"
	"github.com/arvinpaundra/repository-api/models/web/collection/response"
)

type CollectionService interface {
	Create(ctx context.Context, collection request.CreateCollectionRequest) error
	Update(ctx context.Context, collection request.UpdateCollectionRequest, collectionId string) error
	Delete(ctx context.Context, collectionId string) error
	FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.CollectionResponse, int, int, error)
	FindById(ctx context.Context, collectionId string) (response.CollectionResponse, error)
}

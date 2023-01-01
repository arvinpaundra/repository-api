package service

import (
	"context"
	"math"

	"github.com/arvinpaundra/repository-api/drivers/mysql/collection"
	"github.com/arvinpaundra/repository-api/models/web/collection/request"
	"github.com/arvinpaundra/repository-api/models/web/collection/response"
	"github.com/google/uuid"
)

type CollectionServiceImpl struct {
	collectionRepository collection.CollectionRepository
}

func NewCollectionService(collectionRepository collection.CollectionRepository) CollectionService {
	return CollectionServiceImpl{
		collectionRepository: collectionRepository,
	}
}

func (service CollectionServiceImpl) Create(ctx context.Context, collection request.CreateCollectionRequest) error {
	collectionDomain := collection.ToDomainCollection()

	collectionDomain.ID = uuid.NewString()

	err := service.collectionRepository.Save(ctx, collectionDomain)

	if err != nil {
		return err
	}

	return nil
}

func (service CollectionServiceImpl) Update(ctx context.Context, collection request.UpdateCollectionRequest, collectionId string) error {
	if _, err := service.collectionRepository.FindById(ctx, collectionId); err != nil {
		return err
	}

	err := service.collectionRepository.Update(ctx, collection.ToDomainCollection(), collectionId)

	if err != nil {
		return err
	}

	return nil
}

func (service CollectionServiceImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.CollectionResponse, int, int, error) {
	collections, totalRows, err := service.collectionRepository.FindAll(ctx, keyword, limit, offset)

	if err != nil {
		return []response.CollectionResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToCollectionsResponse(collections), int(totalRows), int(totalPages), nil
}

func (service CollectionServiceImpl) FindById(ctx context.Context, collectionId string) (response.CollectionResponse, error) {
	collection, err := service.collectionRepository.FindById(ctx, collectionId)

	if err != nil {
		return response.CollectionResponse{}, err
	}

	return response.ToCollectionResponse(collection), nil
}

func (service CollectionServiceImpl) Delete(ctx context.Context, collectionId string) error {
	panic("not implemented")
}

package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type CollectionResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func ToCollectionResponse(domainCollection domain.Collection) CollectionResponse {
	return CollectionResponse{
		ID:        domainCollection.ID,
		Name:      domainCollection.Name,
		CreatedAt: domainCollection.CreatedAt,
		UpdateAt:  domainCollection.UpdatedAt,
	}
}

func ToCollectionsResponse(domainCollection []domain.Collection) []CollectionResponse {
	var collections []CollectionResponse

	for _, collection := range domainCollection {
		collections = append(collections, ToCollectionResponse(collection))
	}

	return collections
}

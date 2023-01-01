package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreateCollectionRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *CreateCollectionRequest) ToDomainCollection() domain.Collection {
	return domain.Collection{
		Name: req.Name,
	}
}

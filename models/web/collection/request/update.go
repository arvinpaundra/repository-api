package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateCollectionRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *UpdateCollectionRequest) ToDomainCollection() domain.Collection {
	return domain.Collection{
		Name: req.Name,
	}
}

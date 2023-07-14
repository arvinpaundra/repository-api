package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateCollectionRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Visibility string `json:"visibility" form:"visibility" validate:"required"`
}

func (req *UpdateCollectionRequest) ToDomainCollection() domain.Collection {
	return domain.Collection{
		Name:       req.Name,
		Visibility: req.Visibility,
	}
}

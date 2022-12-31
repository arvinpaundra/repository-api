package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (req *UpdateCategoryRequest) ToDomainCategory() domain.Category {
	return domain.Category{
		Name: req.Name,
	}
}

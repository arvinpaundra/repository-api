package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreateCategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required,max=255"`
}

func (req *CreateCategoryRequest) ToDomainCategory() domain.Category {
	return domain.Category{
		Name: req.Name,
	}
}

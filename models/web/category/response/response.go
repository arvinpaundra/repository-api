package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type CategoryResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func ToCategoryResponse(domainCategory domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:        domainCategory.ID,
		Name:      domainCategory.Name,
		CreatedAt: domainCategory.CreatedAt,
		UpdateAt:  domainCategory.UpdatedAt,
	}
}

func ToCategoriesResponse(domainCategory []domain.Category) []CategoryResponse {
	var categories []CategoryResponse

	for _, category := range domainCategory {
		categories = append(categories, ToCategoryResponse(category))
	}

	return categories
}

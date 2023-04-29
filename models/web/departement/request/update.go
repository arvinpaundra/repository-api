package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateDepartementRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *UpdateDepartementRequest) ToDomainDepartement() domain.Departement {
	return domain.Departement{
		Name: req.Name,
	}
}

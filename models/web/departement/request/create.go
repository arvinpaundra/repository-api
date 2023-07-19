package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreateDepartementRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
	Code string `json:"code" form:"code" validate:"required,max=5,alpha"`
}

func (req *CreateDepartementRequest) ToDomainDepartement() domain.Departement {
	return domain.Departement{
		Name: req.Name,
		Code: req.Code,
	}
}

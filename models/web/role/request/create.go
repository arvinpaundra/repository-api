package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreateRoleRequest struct {
	Role       string `json:"role" form:"role" validate:"required"`
	Visibility string `json:"visibility" form:"visibility" validate:"required"`
}

func (req *CreateRoleRequest) ToDomainRole() domain.Role {
	return domain.Role{
		Role:       req.Role,
		Visibility: req.Visibility,
	}
}

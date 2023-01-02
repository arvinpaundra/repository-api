package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateRoleRequest struct {
	Role       string `json:"role" form:"role" validate:"required"`
	Visibility string `json:"visibility" form:"visibility" validate:"required"`
}

func (req *UpdateRoleRequest) ToDomainRole() domain.Role {
	return domain.Role{
		Role:       req.Role,
		Visibility: req.Visibility,
	}
}

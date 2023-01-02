package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type RoleResponse struct {
	ID         string    `json:"id"`
	Role       string    `json:"role"`
	Visibility string    `json:"visibility"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToRoleResponse(domainRole domain.Role) RoleResponse {
	return RoleResponse{
		ID:         domainRole.ID,
		Role:       domainRole.Role,
		Visibility: domainRole.Visibility,
		CreatedAt:  domainRole.CreatedAt,
		UpdatedAt:  domainRole.UpdatedAt,
	}
}

func ToRolesResponse(domainRole []domain.Role) []RoleResponse {
	var roles []RoleResponse

	for _, role := range domainRole {
		roles = append(roles, ToRoleResponse(role))
	}

	return roles
}

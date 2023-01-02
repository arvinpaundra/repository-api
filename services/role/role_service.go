package role

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/role/request"
	"github.com/arvinpaundra/repository-api/models/web/role/response"
)

type RoleService interface {
	Create(ctx context.Context, role request.CreateRoleRequest) error
	Update(ctx context.Context, role request.UpdateRoleRequest, roleId string) error
	Delete(ctx context.Context, roleId string) error
	FindAll(ctx context.Context, visibility string) ([]response.RoleResponse, error)
	FindById(ctx context.Context, roleId string) (response.RoleResponse, error)
}

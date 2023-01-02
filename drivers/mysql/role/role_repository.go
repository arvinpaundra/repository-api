package role

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type RoleRepository interface {
	Save(ctx context.Context, role domain.Role) error
	Update(ctx context.Context, role domain.Role, roleId string) error
	Delete(ctx context.Context, roleId string) error
	FindAll(ctx context.Context, visibiltity string) ([]domain.Role, error)
	FindById(ctx context.Context, roleId string) (domain.Role, error)
}

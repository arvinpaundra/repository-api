package departement

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type DepartementRepository interface {
	Save(ctx context.Context, departement domain.Departement) error
	Update(ctx context.Context, departement domain.Departement, departementId string) error
	Delete(ctx context.Context, departementId string) error
	FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.Departement, int, error)
	FindById(ctx context.Context, departementId string) (domain.Departement, error)
}

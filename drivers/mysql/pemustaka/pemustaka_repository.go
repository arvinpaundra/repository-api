package pemustaka

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/request"
	"gorm.io/gorm"
)

type PemustakaRepository interface {
	Save(ctx context.Context, tx *gorm.DB, pemustaka domain.Pemustaka) error
	Update(ctx context.Context, pemustaka domain.Pemustaka, pemustakaId string) error
	FindAll(ctx context.Context, query request.PemustakaRequestQuery, limit int, offset int) ([]domain.Pemustaka, int, error)
	FindById(ctx context.Context, pemustakaId string) (domain.Pemustaka, error)
	FindByUserId(ctx context.Context, userId string) (domain.Pemustaka, error)
	GetTotalPemustakaByDepartementId(ctx context.Context, departementId string) (int, error)
}

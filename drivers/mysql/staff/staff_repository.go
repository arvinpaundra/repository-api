package staff

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/staff/request"
	"gorm.io/gorm"
)

type StaffRepository interface {
	Save(ctx context.Context, tx *gorm.DB, staff domain.Staff) error
	Update(ctx context.Context, tx *gorm.DB, staffId string, staff domain.Staff) error
	FindAll(ctx context.Context, query request.StaffRequestQuery, limit int, offset int) ([]domain.Staff, int, error)
	FindById(ctx context.Context, staffId string) (domain.Staff, error)
	FindByUserId(ctx context.Context, userId string) (domain.Staff, error)
}

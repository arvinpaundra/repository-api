package request_access

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type RequestAccessRepository interface {
	Save(ctx context.Context, tx *gorm.DB, requestAccess domain.RequestAccess) error
	Update(ctx context.Context, tx *gorm.DB, requestAccess domain.RequestAccess, requestAccessId string) error
	FindAll(ctx context.Context, keyword string, status string, limit int, offset int) ([]domain.RequestAccess, int, error)
	FindById(ctx context.Context, requestAccessId string) (domain.RequestAccess, error)
	Total(ctx context.Context, status string) (int, error)
}

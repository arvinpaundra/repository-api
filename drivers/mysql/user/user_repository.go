package user

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user domain.User) error
	Update(ctx context.Context, user domain.User, email string) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindById(ctx context.Context, userId string) (domain.User, error)
}

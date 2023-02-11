package user

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) UserRepository {
	return UserRepositoryImpl{
		conn: conn,
	}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User) error {
	err := tx.WithContext(ctx).Model(&domain.User{}).Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository UserRepositoryImpl) Update(ctx context.Context, user domain.User, email string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.User{}).Where("email = ?", email).Updates(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var rec domain.User

	err := repository.conn.WithContext(ctx).Model(&domain.User{}).Where("email = ?", email).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, utils.ErrUserNotFound
		}

		return domain.User{}, err
	}

	return rec, nil
}

package role

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) RoleRepository {
	return RoleRepositoryImpl{
		conn: conn,
	}
}

func (repository RoleRepositoryImpl) Save(ctx context.Context, role domain.Role) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Role{}).Create(&role).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository RoleRepositoryImpl) Update(ctx context.Context, role domain.Role, roleId string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Role{}).Where("id = ?", roleId).Updates(&role).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository RoleRepositoryImpl) FindAll(ctx context.Context, visibility string) ([]domain.Role, error) {
	var rec []domain.Role

	err := repository.conn.WithContext(ctx).Model(&domain.Role{}).Where("visibility LIKE ?", "%"+visibility+"%").Order("role ASC").Find(&rec).Error

	if err != nil {
		return []domain.Role{}, err
	}

	return rec, nil
}

func (repository RoleRepositoryImpl) FindById(ctx context.Context, roleId string) (domain.Role, error) {
	var rec domain.Role

	err := repository.conn.WithContext(ctx).Model(&domain.Role{}).Where("id = ?", roleId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Role{}, utils.ErrRoleNotFound
		}

		return domain.Role{}, err
	}

	return rec, nil
}

func (repository RoleRepositoryImpl) Delete(ctx context.Context, roleId string) error {
	panic("not implemented")
}

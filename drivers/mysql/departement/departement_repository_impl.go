package departement

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
)

type DepartementRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) DepartementRepository {
	return DepartementRepositoryImpl{
		conn: conn,
	}
}

func (repository DepartementRepositoryImpl) Save(ctx context.Context, departement domain.Departement) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Departement{}).Create(&departement).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository DepartementRepositoryImpl) Update(ctx context.Context, departement domain.Departement, departementId string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Departement{}).Where("id = ?", departementId).Updates(&departement).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository DepartementRepositoryImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.Departement, int, error) {
	var err error

	var totalRows int64
	err = repository.conn.WithContext(ctx).Model(&domain.Departement{}).
		Where("name LIKE ?", "%"+keyword+"%").Count(&totalRows).Error
	if err != nil {
		return []domain.Departement{}, 0, err
	}

	var rec []domain.Departement
	err = repository.conn.WithContext(ctx).Model(&domain.Departement{}).
		Where("name LIKE ?", "%"+keyword+"%").Limit(limit).Offset(offset).
		Order("name ASC").Find(&rec).Error
	if err != nil {
		return []domain.Departement{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repository DepartementRepositoryImpl) FindById(ctx context.Context, departementId string) (domain.Departement, error) {
	var rec domain.Departement

	err := repository.conn.WithContext(ctx).Model(&domain.Departement{}).
		Where("id = ?", departementId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Departement{}, utils.ErrDepartementNotFound
		}

		return domain.Departement{}, err
	}

	return rec, nil
}

func (repository DepartementRepositoryImpl) Delete(ctx context.Context, departementId string) error {
	panic("not implemented")
}

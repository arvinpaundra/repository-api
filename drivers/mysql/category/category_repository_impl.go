package category

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	conn *gorm.DB
}

func NewRepositorySQL(conn *gorm.DB) CategoryRepository {
	return CategoryRepositoryImpl{
		conn: conn,
	}
}

func (repository CategoryRepositoryImpl) Save(ctx context.Context, category domain.Category) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Category{}).Create(&category).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category, categoryId string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Category{}).Where("id = ?", categoryId).Updates(&category).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, categoryId string) error {
	panic("not implemented")
}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.Category, int64, error) {
	var err error

	var totalRows int64
	err = repository.conn.WithContext(ctx).Model(&domain.Category{}).Where("name LIKE ?", "%"+keyword+"%").Count(&totalRows).Error
	if err != nil {
		return []domain.Category{}, 0, err
	}

	var rec []domain.Category
	err = repository.conn.WithContext(ctx).Model(&domain.Category{}).Where("name LIKE ?", "%"+keyword+"%").Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Category{}, 0, err
	}

	return rec, totalRows, nil
}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, categoryId string) (domain.Category, error) {
	var rec domain.Category

	err := repository.conn.WithContext(ctx).Model(&domain.Category{}).Where("id = ?", categoryId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Category{}, utils.ErrCategoryNotFound
		}

		return domain.Category{}, err
	}

	return rec, nil
}

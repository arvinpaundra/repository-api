package collection

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
)

type CollectionRepositoryImpl struct {
	conn *gorm.DB
}

func NewRepositorySQL(conn *gorm.DB) CollectionRepository {
	return CollectionRepositoryImpl{
		conn: conn,
	}
}

func (repository CollectionRepositoryImpl) Save(ctx context.Context, collection domain.Collection) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Collection{}).Create(&collection).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository CollectionRepositoryImpl) Update(ctx context.Context, collection domain.Collection, collectionId string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Collection{}).Where("id = ?", collectionId).Updates(&collection).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository CollectionRepositoryImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.Collection, int, error) {
	var err error

	var totalRows int64
	err = repository.conn.WithContext(ctx).Model(&domain.Collection{}).Where("name LIKE ?", "%"+keyword+"%").Count(&totalRows).Error
	if err != nil {
		return []domain.Collection{}, 0, err
	}

	var rec []domain.Collection
	err = repository.conn.WithContext(ctx).Model(&domain.Collection{}).Where("name LIKE ?", "%"+keyword+"%").Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Collection{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repository CollectionRepositoryImpl) FindById(ctx context.Context, collectionId string) (domain.Collection, error) {
	var rec domain.Collection

	err := repository.conn.WithContext(ctx).Model(&domain.Collection{}).Where("id = ?", collectionId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Collection{}, utils.ErrCollectionNotFound
		}

		return domain.Collection{}, err
	}

	return rec, nil
}

func (repository CollectionRepositoryImpl) Delete(ctx context.Context, collectionId string) error {
	panic("implement me")
}

package document

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type DocumentRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) DocumentRepository {
	return DocumentRepositoryImpl{
		conn: conn,
	}
}

func (repository DocumentRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, document domain.Document) error {
	err := tx.WithContext(ctx).Model(&domain.Document{}).Create(&document).Error

	if err != nil {
		return err
	}

	return err
}

func (repository DocumentRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, document domain.Document, repositoryId string) error {
	err := tx.WithContext(ctx).Model(&domain.Document{}).Where("repository_id = ?", repositoryId).Updates(&document).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository DocumentRepositoryImpl) FindByRepositoryId(ctx context.Context, repositoryId string) (domain.Document, error) {
	var rec domain.Document

	err := repository.conn.WithContext(ctx).Model(&domain.Document{}).
		Where("repository_id = ?", repositoryId).First(&rec).Error

	if err != nil {
		return domain.Document{}, err
	}

	return rec, nil
}

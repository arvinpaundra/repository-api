package author

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) AuthorRepository {
	return AuthorRepositoryImpl{
		conn: conn,
	}
}

func (repository AuthorRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, author []domain.Author) error {
	err := tx.WithContext(ctx).Model(&domain.Author{}).Create(&author).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository AuthorRepositoryImpl) Delete(ctx context.Context, repositoryId string, pemustakaId string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Author{}).Unscoped().
		Where("repository_id = ? AND pemustaka_id = ?", repositoryId, pemustakaId).Delete(&domain.Author{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository AuthorRepositoryImpl) FindByRepositoryId(ctx context.Context, repositoryId string) ([]domain.Author, error) {
	var rec []domain.Author

	err := repository.conn.WithContext(ctx).Model(&domain.Author{}).Preload("Pemustaka.User").
		Where("repository_id = ?", repositoryId).Find(&rec).Error

	if err != nil {
		return []domain.Author{}, err
	}

	return rec, nil
}

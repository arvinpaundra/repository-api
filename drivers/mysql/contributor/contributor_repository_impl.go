package contributor

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type ContributorRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) ContributorRepository {
	return ContributorRepositoryImpl{
		conn: conn,
	}
}

func (repository ContributorRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, contributor domain.Contributor) error {
	err := tx.WithContext(ctx).Model(&domain.Contributor{}).Create(&contributor).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository ContributorRepositoryImpl) Delete(ctx context.Context, repositoryId string, pemustakaId string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.Contributor{}).
		Where("repository_id = ? AND pemustaka_id = ?", repositoryId, pemustakaId).
		Delete(&domain.Contributor{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository ContributorRepositoryImpl) FindByRepositoryId(ctx context.Context, repositoryId string) ([]domain.Contributor, error) {
	var rec []domain.Contributor

	err := repository.conn.WithContext(ctx).Model(&domain.Contributor{}).Preload("Pemustaka").
		Where("repository_id = ?", repositoryId).Order("contributed_as ASC").
		Find(&rec).Error

	if err != nil {
		return []domain.Contributor{}, err
	}

	return rec, nil
}

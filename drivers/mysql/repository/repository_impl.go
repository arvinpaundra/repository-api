package repository

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/repository/request"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) Repository {
	return RepositoryImpl{
		conn: conn,
	}
}

func (repo RepositoryImpl) Save(ctx context.Context, tx *gorm.DB, repository domain.Repository) error {
	err := tx.WithContext(ctx).Model(&domain.Repository{}).Create(&repository).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo RepositoryImpl) Update(ctx context.Context, tx *gorm.DB, repository domain.Repository, repositoryId string) error {
	err := tx.WithContext(ctx).Model(&domain.Repository{}).Where("id = ?", repositoryId).Updates(&repository).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo RepositoryImpl) Delete(ctx context.Context, repositoryId string) error {
	err := repo.conn.WithContext(ctx).Unscoped().Model(&domain.Repository{}).Where("id = ?", repositoryId).Delete(&domain.Repository{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo RepositoryImpl) FindAll(ctx context.Context, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	var err error

	var totalRows int64
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Where(
			"title LIKE ? AND collection_id LIKE ? AND departement_id LIKE ? AND date_validated LIKE ? AND improvement LIKE ? AND status LIKE ? AND category_id LIKE ?",
			"%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%"+query.CategoryId+"%",
		).Count(&totalRows).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	var rec []domain.Repository
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Preload("Collection").Preload("Departement").Preload("Authors.Pemustaka").Preload("Category").
		Where(
			"title LIKE ? AND collection_id LIKE ? AND departement_id LIKE ? AND date_validated LIKE ? AND improvement LIKE ? AND status LIKE ? AND category_id LIKE ?",
			"%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%"+query.CategoryId+"%",
		).Order(query.Sort).Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repo RepositoryImpl) FindById(ctx context.Context, repositoryId string) (domain.Repository, error) {
	var rec domain.Repository

	err := repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Preload("Collection").Preload("Departement").Preload("Category").
		Where("id = ?", repositoryId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Repository{}, utils.ErrRepositoryNotFound
		}

		return domain.Repository{}, err
	}

	return rec, nil
}

func (repo RepositoryImpl) FindByAuthorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	var err error

	var totalRows int64
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Joins("INNER JOIN authors ON repositories.id = authors.repository_id").
		Where(
			"authors.pemustaka_id = ? AND repositories.title LIKE ? AND repositories.collection_id LIKE ? AND repositories.departement_id LIKE ? AND date_validated LIKE ? AND repositories.improvement LIKE ? AND repositories.status LIKE ? AND repositories.category_id LIKE ?",
			pemustakaId, "%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%"+query.CategoryId+"%",
		).Count(&totalRows).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	var rec []domain.Repository
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Preload("Collection").Preload("Departement").Preload("Authors.Pemustaka").Preload("Category").
		Joins("INNER JOIN authors ON repositories.id = authors.repository_id").
		Where(
			"authors.pemustaka_id = ? AND repositories.title LIKE ? AND repositories.collection_id LIKE ? AND repositories.departement_id LIKE ? AND date_validated LIKE ? AND repositories.improvement LIKE ? AND repositories.status LIKE ? AND repositories.category_id LIKE ?",
			pemustakaId, "%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%"+query.CategoryId+"%",
		).Order(query.Sort).Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repo RepositoryImpl) FindByMentorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	var err error

	var totalRows int64
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Joins("INNER JOIN contributors ON repositories.id = contributors.repository_id").
		Where(
			"contributors.pemustaka_id = ? AND repositories.title LIKE ? AND repositories.collection_id LIKE ? AND repositories.departement_id LIKE ? AND date_validated LIKE ? AND repositories.improvement LIKE ? AND repositories.status LIKE ? AND contributors.contributed_as LIKE ? AND repositories.category_id LIKE ?",
			pemustakaId, "%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%Pembimbing%", "%"+query.CategoryId+"%",
		).Count(&totalRows).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	var rec []domain.Repository
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Preload("Collection").Preload("Departement").Preload("Authors.Pemustaka").Preload("Category").
		Joins("INNER JOIN contributors ON repositories.id = contributors.repository_id").
		Where(
			"contributors.pemustaka_id = ? AND repositories.title LIKE ? AND repositories.collection_id LIKE ? AND repositories.departement_id LIKE ? AND date_validated LIKE ? AND repositories.improvement LIKE ? AND repositories.status LIKE ? AND contributors.contributed_as LIKE ? AND repositories.category_id LIKE ?",
			pemustakaId, "%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%Pembimbing%", "%"+query.CategoryId+"%",
		).Order(query.Sort).Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repo RepositoryImpl) FindByExaminerId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	var err error

	var totalRows int64
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Joins("INNER JOIN contributors ON repositories.id = contributors.repository_id").
		Where(
			"contributors.pemustaka_id = ? AND repositories.title LIKE ? AND repositories.collection_id LIKE ? AND repositories.departement_id LIKE ? AND date_validated LIKE ? AND repositories.improvement LIKE ? AND repositories.status LIKE ? AND contributors.contributed_as LIKE ? AND repositories.category_id LIKE ?",
			pemustakaId, "%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%Penguji%", "%"+query.CategoryId+"%",
		).Count(&totalRows).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	var rec []domain.Repository
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Preload("Collection").Preload("Departement").Preload("Authors.Pemustaka").Preload("Category").
		Joins("INNER JOIN contributors ON repositories.id = contributors.repository_id").
		Where(
			"contributors.pemustaka_id = ? AND repositories.title LIKE ? AND repositories.collection_id LIKE ? AND repositories.departement_id LIKE ? AND date_validated LIKE ? AND repositories.improvement LIKE ? AND repositories.status LIKE ? AND contributors.contributed_as LIKE ? AND repositories.category_id LIKE ?",
			pemustakaId, "%"+query.Keyword+"%", "%"+query.CollectionId+"%", "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "%"+query.Status+"%", "%Penguji%", "%"+query.CategoryId+"%",
		).Order(query.Sort).Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repo RepositoryImpl) FindByCollectionId(ctx context.Context, collectionId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	var err error

	var totalRows int64
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Where(
			"title LIKE ? AND collection_id = ? AND departement_id LIKE ? AND date_validated LIKE ? AND improvement LIKE ? AND status = ? AND category_id LIKE ?",
			"%"+query.Keyword+"%", collectionId, "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "approved", "%"+query.CategoryId+"%",
		).Count(&totalRows).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	var rec []domain.Repository
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Preload("Collection").Preload("Departement").Preload("Authors.Pemustaka").Preload("Category").
		Where(
			"title LIKE ? AND collection_id = ? AND departement_id LIKE ? AND date_validated LIKE ? AND improvement LIKE ? AND status = ? AND category_id LIKE ?",
			"%"+query.Keyword+"%", collectionId, "%"+query.DepartementId+"%", query.Year+"%", "%"+query.Improvement+"%", "approved", "%"+query.CategoryId+"%",
		).Order(query.Sort).Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repo RepositoryImpl) FindByDepartementId(ctx context.Context, departementId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	var err error

	var totalRows int64
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Where(
			"title LIKE ? AND collection_id LIKE ? AND departement_id = ? AND date_validated LIKE ? AND improvement LIKE ? AND status = ? AND category_id LIKE ?",
			"%"+query.Keyword+"%", "%"+query.CollectionId+"%", departementId, query.Year+"%", "%"+query.Improvement+"%", "approved", "%"+query.CategoryId+"%",
		).Count(&totalRows).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	var rec []domain.Repository
	err = repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Preload("Collection").Preload("Departement").Preload("Authors.Pemustaka").Preload("Category").
		Where(
			"title LIKE ? AND collection_id LIKE ? AND departement_id = ? AND date_validated LIKE ? AND improvement LIKE ? AND status = ? AND category_id LIKE ?",
			"%"+query.Keyword+"%", "%"+query.CollectionId+"%", departementId, query.Year+"%", "%"+query.Improvement+"%", "approved", "%"+query.CategoryId+"%",
		).Order(query.Sort).Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Repository{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repo RepositoryImpl) GetTotal(ctx context.Context, status string) (int, error) {
	var total int64

	err := repo.conn.WithContext(ctx).Model(&domain.Repository{}).
		Where("status LIKE ?", "%"+status+"%").Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), nil
}

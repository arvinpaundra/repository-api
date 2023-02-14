package request_access

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RequestAccessRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) RequestAccessRepository {
	return RequestAccessRepositoryImpl{
		conn: conn,
	}
}

func (repository RequestAccessRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, requestAccess domain.RequestAccess) error {
	err := tx.WithContext(ctx).Model(domain.RequestAccess{}).Create(&requestAccess).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository RequestAccessRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, requestAccess domain.RequestAccess, requestAccessId string) error {
	err := tx.WithContext(ctx).Model(domain.RequestAccess{}).Where("id = ?", requestAccessId).Updates(&requestAccess).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository RequestAccessRepositoryImpl) FindAll(ctx context.Context, keyword string, status string, limit int, offset int) ([]domain.RequestAccess, int, error) {
	var err error

	var totalRows int64
	err = repository.conn.WithContext(ctx).Model(domain.RequestAccess{}).Preload(clause.Associations).
		Joins("INNER JOIN pemustakas ON request_accesses.pemustaka_id = pemustakas.id").
		Where("pemustakas.fullname LIKE ? AND request_accesses.status LIKE ?", "%"+keyword+"%", "%"+status+"%").
		Count(&totalRows).Error
	if err != nil {
		return []domain.RequestAccess{}, 0, err
	}

	var rec []domain.RequestAccess
	err = repository.conn.WithContext(ctx).Model(domain.RequestAccess{}).
		Preload(clause.Associations).Preload("Pemustaka.StudyProgram").Preload("Pemustaka.Departement").Preload("Pemustaka.Role").
		Joins("INNER JOIN pemustakas ON request_accesses.pemustaka_id = pemustakas.id").
		Where("pemustakas.fullname LIKE ? AND request_accesses.status LIKE ?", "%"+keyword+"%", "%"+status+"%").
		Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.RequestAccess{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repository RequestAccessRepositoryImpl) FindById(ctx context.Context, requestAccessId string) (domain.RequestAccess, error) {
	var rec domain.RequestAccess

	err := repository.conn.WithContext(ctx).Model(domain.RequestAccess{}).
		Preload(clause.Associations).Preload("Pemustaka.StudyProgram").Preload("Pemustaka.Departement").Preload("Pemustaka.Role").
		Joins("INNER JOIN pemustakas ON request_accesses.pemustaka_id = pemustakas.id").
		Where("request_accesses.id = ?", requestAccessId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.RequestAccess{}, utils.ErrRequestAccessNotFound
		}

		return domain.RequestAccess{}, err
	}

	return rec, nil
}

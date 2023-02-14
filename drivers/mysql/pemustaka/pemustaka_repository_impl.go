package pemustaka

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/request"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PemustakaRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) PemustakaRepository {
	return PemustakaRepositoryImpl{
		conn: conn,
	}
}

func (repository PemustakaRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, pemustaka domain.Pemustaka) error {
	err := tx.WithContext(ctx).Model(&domain.Pemustaka{}).Create(&pemustaka).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository PemustakaRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, pemustaka domain.Pemustaka, pemustakaId string) error {
	err := tx.WithContext(ctx).Model(&domain.Pemustaka{}).Where("id = ?", pemustakaId).Updates(&pemustaka).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository PemustakaRepositoryImpl) FindAll(ctx context.Context, query request.PemustakaRequestQuery, limit int, offset int) ([]domain.Pemustaka, int, error) {
	var err error

	var totalRows int64
	err = repository.conn.WithContext(ctx).Model(&domain.Pemustaka{}).
		Where("fullname LIKE ? AND role_id = ? AND departement_id LIKE ? AND is_collected_final_project LIKE ? AND year_gen LIKE ?", "%"+query.Keyword+"%", query.RoleId, "%"+query.DepartementId+"%", "%"+query.IsCollectedFinalProject+"%", "%"+query.YearGen+"%").
		Count(&totalRows).Error
	if err != nil {
		return []domain.Pemustaka{}, 0, err
	}

	var rec []domain.Pemustaka
	err = repository.conn.WithContext(ctx).Model(&domain.Pemustaka{}).Preload(clause.Associations).
		Where("fullname LIKE ? AND role_id = ? AND departement_id LIKE ? AND is_collected_final_project LIKE ? AND year_gen LIKE ?", "%"+query.Keyword+"%", query.RoleId, "%"+query.DepartementId+"%", "%"+query.IsCollectedFinalProject+"%", "%"+query.YearGen+"%").
		Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return []domain.Pemustaka{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repository PemustakaRepositoryImpl) FindById(ctx context.Context, pemustakaId string) (domain.Pemustaka, error) {
	var rec domain.Pemustaka

	err := repository.conn.WithContext(ctx).Model(&domain.Pemustaka{}).Preload(clause.Associations).
		Where("id = ?", pemustakaId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Pemustaka{}, utils.ErrPemustakaNotFound
		}

		return domain.Pemustaka{}, err
	}

	return rec, nil
}

func (repository PemustakaRepositoryImpl) FindByUserId(ctx context.Context, userId string) (domain.Pemustaka, error) {
	var rec domain.Pemustaka

	err := repository.conn.WithContext(ctx).Model(&domain.Pemustaka{}).Preload(clause.Associations).
		Where("user_id = ?", userId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Pemustaka{}, utils.ErrUserNotFound
		}

		return domain.Pemustaka{}, err
	}

	return rec, nil
}

func (repository PemustakaRepositoryImpl) GetTotalPemustakaByDepartementId(ctx context.Context, departementId string) (int, error) {
	var totalRows int64

	err := repository.conn.WithContext(ctx).Model(&domain.Pemustaka{}).Where("departement_id = ?", departementId).Count(&totalRows).Error

	if err != nil {
		return 0, err
	}

	return int(totalRows), nil
}

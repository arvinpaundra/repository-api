package staff

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/staff/request"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StaffRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) StaffRepository {
	return StaffRepositoryImpl{
		conn: conn,
	}
}

func (repository StaffRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, staff domain.Staff) error {
	err := tx.WithContext(ctx).Model(&domain.Staff{}).Create(&staff).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository StaffRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, staffId string, staff domain.Staff) error {
	err := tx.WithContext(ctx).Model(&domain.Staff{}).Where("id = ?", staffId).Updates(&staff).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository StaffRepositoryImpl) FindAll(ctx context.Context, query request.StaffRequestQuery, limit int, offset int) ([]domain.Staff, int, error) {
	var err error

	var totalRows int64
	err = repository.conn.WithContext(ctx).Model(&domain.Staff{}).
		Where("fullname LIKE ? AND role_id LIKE ?", "%"+query.Keyword+"%", "%"+query.RoleId+"%").Count(&totalRows).Error
	if err != nil {
		return []domain.Staff{}, 0, err
	}

	var rec []domain.Staff
	err = repository.conn.WithContext(ctx).Model(&domain.Staff{}).Preload(clause.Associations).
		Where("fullname LIKE ? AND role_id LIKE ?", "%"+query.Keyword+"%", "%"+query.RoleId+"%").
		Limit(limit).Offset(offset).Order("fullname ASC").
		Find(&rec).Error
	if err != nil {
		return []domain.Staff{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repository StaffRepositoryImpl) FindById(ctx context.Context, staffId string) (domain.Staff, error) {
	var rec domain.Staff

	err := repository.conn.WithContext(ctx).Model(&domain.Staff{}).Preload(clause.Associations).
		Where("id = ?", staffId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Staff{}, utils.ErrStaffNotFound
		}

		return domain.Staff{}, err
	}

	return rec, nil
}

func (repository StaffRepositoryImpl) FindByUserId(ctx context.Context, userId string) (domain.Staff, error) {
	var rec domain.Staff

	err := repository.conn.WithContext(ctx).Model(&domain.Staff{}).Preload(clause.Associations).
		Where("user_id = ?", userId).First(&rec).Error

	if err != nil {
		return domain.Staff{}, utils.ErrStaffNotFound
	}

	return rec, nil
}

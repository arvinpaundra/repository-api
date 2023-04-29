package study_program

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/utils"
	"gorm.io/gorm"
)

type StudyProgramRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) StudyProgramRepository {
	return StudyProgramRepositoryImpl{
		conn: conn,
	}
}

func (repository StudyProgramRepositoryImpl) Save(ctx context.Context, studyProgram domain.StudyProgram) error {
	err := repository.conn.WithContext(ctx).Model(&domain.StudyProgram{}).Create(&studyProgram).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository StudyProgramRepositoryImpl) Update(ctx context.Context, studyProgram domain.StudyProgram, studyProgramId string) error {
	err := repository.conn.WithContext(ctx).Model(&domain.StudyProgram{}).Where("id = ?", studyProgramId).Updates(&studyProgram).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository StudyProgramRepositoryImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.StudyProgram, int, error) {
	var err error

	var totalRows int64
	err = repository.conn.WithContext(ctx).Model(&domain.StudyProgram{}).
		Where("name LIKE ?", "%"+keyword+"%").Count(&totalRows).Error
	if err != nil {
		return []domain.StudyProgram{}, 0, err
	}

	var rec []domain.StudyProgram
	err = repository.conn.WithContext(ctx).Model(&domain.StudyProgram{}).Preload("Departement").
		Where("name LIKE ?", "%"+keyword+"%").Limit(limit).Offset(offset).
		Order("name ASC").Find(&rec).Error
	if err != nil {
		return []domain.StudyProgram{}, 0, err
	}

	return rec, int(totalRows), nil
}

func (repository StudyProgramRepositoryImpl) FindById(ctx context.Context, studyProgramId string) (domain.StudyProgram, error) {
	var rec domain.StudyProgram

	err := repository.conn.WithContext(ctx).Model(&domain.StudyProgram{}).Preload("Departement").
		Where("id = ?", studyProgramId).First(&rec).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.StudyProgram{}, utils.ErrStudyProgramNotFound
		}

		return domain.StudyProgram{}, err
	}

	return rec, nil
}

func (repository StudyProgramRepositoryImpl) FindByDepartementId(ctx context.Context, departementId string) ([]domain.StudyProgram, error) {
	var rec []domain.StudyProgram

	err := repository.conn.WithContext(ctx).Model(&domain.StudyProgram{}).Preload("Departement").
		Where("departement_id = ?", departementId).Find(&rec).Error
	if err != nil {
		return []domain.StudyProgram{}, err
	}

	return rec, nil
}

func (repository StudyProgramRepositoryImpl) Delete(ctx context.Context, studyProgramId string) error {
	panic("not implemented")
}

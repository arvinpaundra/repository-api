package report

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type ReportRepositoryImpl struct {
	conn *gorm.DB
}

func NewSQLRepository(conn *gorm.DB) ReportRepository {
	return ReportRepositoryImpl{
		conn: conn,
	}
}

func (repository ReportRepositoryImpl) RecapCollectedReport(ctx context.Context, roleId, yearGen, collectionId string) ([]domain.Report, error) {
	var rec []domain.Report

	err := repository.conn.WithContext(ctx).Raw(`
		SELECT study_programs.name AS study_program, COUNT(pemustakas.id) AS total_pemustakas, COALESCE(pemustaka_count.count, 0) AS pemustaka_count 
		FROM study_programs
		LEFT JOIN pemustakas ON pemustakas.study_program_id = study_programs.id
			AND pemustakas.role_id LIKE ?
		LEFT JOIN (
			SELECT pemustakas.study_program_id, COUNT(pemustakas.id) AS count 
			FROM pemustakas 
			JOIN authors ON authors.pemustaka_id = pemustakas.id 
			JOIN repositories ON repositories.id = authors.repository_id 
			JOIN collections ON collections.id = repositories.collection_id 
			WHERE pemustakas.year_gen LIKE ? AND collections.id LIKE ? AND pemustakas.role_id LIKE ?
			GROUP BY pemustakas.study_program_id
		) AS pemustaka_count ON pemustaka_count.study_program_id = study_programs.id
		GROUP BY study_programs.id, pemustaka_count.count;
	`, "%"+roleId+"%", "%"+yearGen+"%", "%"+collectionId+"%", "%"+roleId+"%",
	).Scan(&rec).Error

	if err != nil {
		return []domain.Report{}, err
	}

	return rec, nil
}

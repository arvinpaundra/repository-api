package report

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type ReportRepository interface {
	RecapCollectedReport(ctx context.Context, roleId, yearGen, collectionId string) ([]domain.Report, error)
}

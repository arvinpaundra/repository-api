package report

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/report/request"
	"github.com/arvinpaundra/repository-api/models/web/report/response"
)

type ReportService interface {
	SuratKeteranganPenyerahanLaporan(ctx context.Context, req request.SuratKeteranganPenyerahanLaporanRequest) ([]byte, error)
	RecapCollectedReport(ctx context.Context, yearGen string, collectionId string) ([]response.RecapCollectedReportResponse, error)
	DownloadRecapCollectedReport(ctx context.Context, query request.QueryRecapCollectedReport) ([]byte, error)
}

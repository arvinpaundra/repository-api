package report

import (
	"bytes"
	"context"
	"html/template"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/drivers/mysql/collection"
	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/drivers/mysql/report"
	"github.com/arvinpaundra/repository-api/drivers/mysql/repository"
	"github.com/arvinpaundra/repository-api/drivers/mysql/staff"
	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/domain"
	requestReport "github.com/arvinpaundra/repository-api/models/web/report/request"
	"github.com/arvinpaundra/repository-api/models/web/report/response"
	requestStaff "github.com/arvinpaundra/repository-api/models/web/staff/request"
	"github.com/arvinpaundra/repository-api/templates"
	"github.com/arvinpaundra/repository-api/utils"
)

type ReportServiceImpl struct {
	pemustakaRepository  pemustaka.PemustakaRepository
	staffRepository      staff.StaffRepository
	collectionRepository collection.CollectionRepository
	reportRepository     report.ReportRepository
	repository           repository.Repository
}

func NewReportService(
	pemustakaRepository pemustaka.PemustakaRepository,
	collectionRepository collection.CollectionRepository,
	staffRepository staff.StaffRepository,
	reportRepository report.ReportRepository,
	repository repository.Repository,
) ReportService {
	return ReportServiceImpl{
		pemustakaRepository:  pemustakaRepository,
		collectionRepository: collectionRepository,
		staffRepository:      staffRepository,
		reportRepository:     reportRepository,
		repository:           repository,
	}
}

func (service ReportServiceImpl) SuratKeteranganPenyerahanLaporan(ctx context.Context, req requestReport.SuratKeteranganPenyerahanLaporanRequest) ([]byte, error) {
	wkhtmltopdf.SetPath(configs.GetConfig("WKHTMLTOPDF_PATH"))

	pemustaka, err := service.pemustakaRepository.FindById(ctx, req.PemustakaId)

	if err != nil {
		return nil, err
	}

	if req.CollectionId == configs.GetConfig("ID_FINAL_PROJECT") && pemustaka.IsCollectedFinalProject != "1" {
		return nil, utils.ErrNotCollectedFinalProject
	}

	if req.CollectionId == configs.GetConfig("ID_INTERNSHIP_REPORT") && pemustaka.IsCollectedInternshipReport != "1" {
		return nil, utils.ErrNotCollectedInternshipReport
	}

	collection, err := service.collectionRepository.FindById(ctx, req.CollectionId)

	if err != nil {
		return nil, err
	}

	repository, err := service.repository.FindById(ctx, req.RepositoryId)

	if err != nil {
		return nil, err
	}

	staff, _, err := service.staffRepository.FindAll(ctx, requestStaff.StaffRequestQuery{RoleId: configs.GetConfig("ID_ROLE_KEPALA_PERPUS")}, 1, 0)

	if err != nil {
		return nil, err
	}

	if len(staff) == 0 {
		return nil, utils.ErrHeadOfLibraryNotFound
	}

	tmpl, err := template.New("").Parse(templates.SuratKeteranganPenyerahanLaporan)

	if err != nil {
		return nil, err
	}

	data := map[string]string{
		"fullname":       pemustaka.Fullname,
		"identityNumber": pemustaka.IdentityNumber,
		"programStudy":   pemustaka.StudyProgram.Name,
		"collection":     collection.Name,
		"title":          repository.Title,
		"dateIssued":     helper.FormatDate(time.Now()),
		"headOfLibrary":  staff[0].Fullname,
		"nip":            staff[0].Nip,
	}

	var buff bytes.Buffer

	if err := tmpl.Execute(&buff, data); err != nil {
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewBuffer(buff.Bytes()))

	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	pdfg.MarginBottom.Set(10)
	pdfg.MarginLeft.Set(15)
	pdfg.MarginTop.Set(10)
	pdfg.MarginRight.Set(15)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	if err := pdfg.Create(); err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

func (service ReportServiceImpl) RecapCollectedReport(ctx context.Context, yearGen string, collectionId string) ([]response.RecapCollectedReportResponse, error) {

	collection, err := service.collectionRepository.FindById(ctx, collectionId)

	if err != nil {
		return nil, err
	}

	var reports []domain.Report

	if collection.Visibility == "Semua" {
		reports, err = service.reportRepository.RecapCollectedReport(ctx, "", yearGen, collectionId)
	} else if collection.Visibility == "Dosen" {
		reports, err = service.reportRepository.RecapCollectedReport(ctx, configs.GetConfig("ID_ROLE_DOSEN"), yearGen, collectionId)
	} else {
		reports, err = service.reportRepository.RecapCollectedReport(ctx, configs.GetConfig("ID_ROLE_MAHASISWA"), yearGen, collectionId)
	}

	if err != nil {
		return []response.RecapCollectedReportResponse{}, err
	}

	return response.ToRecapCollectedReportArrayResponse(reports), nil
}

func (service ReportServiceImpl) DownloadRecapCollectedReport(ctx context.Context, query requestReport.QueryRecapCollectedReport) ([]byte, error) {
	wkhtmltopdf.SetPath(configs.GetConfig("WKHTMLTOPDF_PATH"))

	collection, err := service.collectionRepository.FindById(ctx, query.CollectionId)

	if err != nil {
		return nil, err
	}

	var reports []domain.Report

	if collection.Visibility == "Semua" {
		reports, err = service.reportRepository.RecapCollectedReport(ctx, "", query.YearGen, query.CollectionId)
	} else if collection.Visibility == "Dosen" {
		reports, err = service.reportRepository.RecapCollectedReport(ctx, configs.GetConfig("ID_ROLE_DOSEN"), query.YearGen, query.CollectionId)
	} else {
		reports, err = service.reportRepository.RecapCollectedReport(ctx, configs.GetConfig("ID_ROLE_MAHASISWA"), query.YearGen, query.CollectionId)
	}

	if err != nil {
		return nil, err
	}

	staff, _, err := service.staffRepository.FindAll(ctx, requestStaff.StaffRequestQuery{RoleId: configs.GetConfig("ID_ROLE_KEPALA_PERPUS")}, 1, 0)

	if err != nil {
		return nil, err
	}

	if len(staff) == 0 {
		return nil, utils.ErrHeadOfLibraryNotFound
	}

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"add": add,
		"sub": sub,
	}).Parse(templates.RekapPenyerahanLaporan)

	if err != nil {
		return nil, err
	}

	totalCollected := 0
	totalNotCollected := 0

	for _, report := range reports {
		totalCollected += int(report.PemustakaCount)
		totalNotCollected += int(report.TotalPemustakas - report.PemustakaCount)
	}

	data := map[string]interface{}{
		"collection":        collection.Name,
		"yearGen":           query.YearGen,
		"reports":           reports,
		"totalCollected":    totalCollected,
		"totalNotCollected": totalNotCollected,
		"dateIssued":        helper.FormatDate(time.Now()),
		"headOfLibrary":     staff[0].Fullname,
		"nip":               staff[0].Nip,
	}

	var buff bytes.Buffer

	if err := tmpl.Execute(&buff, data); err != nil {
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewBuffer(buff.Bytes()))

	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	pdfg.MarginBottom.Set(10)
	pdfg.MarginLeft.Set(15)
	pdfg.MarginTop.Set(10)
	pdfg.MarginRight.Set(15)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	if err := pdfg.Create(); err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int64) int {
	return int(a - b)
}

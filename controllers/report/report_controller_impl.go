package report

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/report/request"
	"github.com/arvinpaundra/repository-api/services/report"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type ReportControllerImpl struct {
	reportService report.ReportService
}

func NewReportController(reportService report.ReportService) ReportController {
	return ReportControllerImpl{
		reportService: reportService,
	}
}

func (ctrl ReportControllerImpl) HandlerGetSuratKeteranganPenyerahanLaporan(c echo.Context) error {
	var req request.SuratKeteranganPenyerahanLaporanRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	pdf, err := ctrl.reportService.SuratKeteranganPenyerahanLaporan(c.Request().Context(), req)

	if err != nil {
		switch err {
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrNotCollectedFinalProject:
			return c.JSON(http.StatusUnprocessableEntity, helper.UnprocessableContentResponse(err.Error()))
		case utils.ErrNotCollectedInternshipReport:
			return c.JSON(http.StatusUnprocessableEntity, helper.UnprocessableContentResponse(err.Error()))
		case utils.ErrRepositoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrHeadOfLibraryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.Blob(http.StatusOK, "application/pdf", pdf)
}

func (ctrl ReportControllerImpl) HandlerRecapCollectedReport(c echo.Context) error {
	collectionId := c.QueryParam("collection_id")
	yearGen := c.QueryParam("year_gen")

	reports, err := ctrl.reportService.RecapCollectedReport(c.Request().Context(), yearGen, collectionId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(reports))
}

func (ctrl ReportControllerImpl) HandlerDownloadRecapCollectedReport(c echo.Context) error {
	collectionId := c.QueryParam("collection_id")
	if collectionId == "" {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.collection_id": "Bagian ini wajib diisi",
		}))
	}

	yearGen := c.QueryParam("year_gen")

	query := request.QueryRecapCollectedReport{
		CollectionId: collectionId,
		YearGen:      yearGen,
	}

	pdf, err := ctrl.reportService.DownloadRecapCollectedReport(c.Request().Context(), query)

	if err != nil {
		switch err {
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrHeadOfLibraryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.Blob(http.StatusOK, "application/pdf", pdf)
}

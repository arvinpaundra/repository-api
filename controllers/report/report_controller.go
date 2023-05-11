package report

import "github.com/labstack/echo/v4"

type ReportController interface {
	HandlerGetSuratKeteranganPenyerahanLaporan(c echo.Context) error
	HandlerRecapCollectedReport(c echo.Context) error
	HandlerDownloadRecapCollectedReport(c echo.Context) error
}

package dashboard

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/services/dashboard"
	"github.com/labstack/echo/v4"
)

type DashboardControllerImpl struct {
	dashboardService dashboard.DashboardService
}

func NewDashboardController(dashboardService dashboard.DashboardService) DashboardController {
	return DashboardControllerImpl{
		dashboardService: dashboardService,
	}
}

func (ctrl DashboardControllerImpl) HandlerOverview(c echo.Context) error {
	totalPemustaka, totalRepository, totalRequestAccessPending, err := ctrl.dashboardService.Overview(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(map[string]interface{}{
		"total_pemustaka":              totalPemustaka,
		"total_repository":             totalRepository,
		"total_request_access_pending": totalRequestAccessPending,
	}))
}

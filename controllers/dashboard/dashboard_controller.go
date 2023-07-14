package dashboard

import "github.com/labstack/echo/v4"

type DashboardController interface {
	HandlerOverview(c echo.Context) error
}

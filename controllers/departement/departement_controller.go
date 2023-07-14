package departement

import "github.com/labstack/echo/v4"

type DepartementController interface {
	HandlerCreateDepartement(c echo.Context) error
	HandlerUpdateDepartement(c echo.Context) error
	HandlerDeleteDepartement(c echo.Context) error
	HandlerFindAllDepartements(c echo.Context) error
	HandlerFindDepartementById(c echo.Context) error
}

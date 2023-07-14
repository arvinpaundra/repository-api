package pemustaka

import "github.com/labstack/echo/v4"

type PemustakaController interface {
	HandleCreatePemustaka(c echo.Context) error
	HandlerRegister(c echo.Context) error
	HandlerLogin(c echo.Context) error
	HandlerUpdatePemustaka(c echo.Context) error
	HandlerFindAllPemustaka(c echo.Context) error
	HandlerFindPemustakaById(c echo.Context) error
}

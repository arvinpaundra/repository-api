package request_access

import "github.com/labstack/echo/v4"

type RequestAccessController interface {
	HandlerUpdateRequestAccess(c echo.Context) error
	HandlerFindAllRequestAccesses(c echo.Context) error
	HandlerFindRequestAccessById(c echo.Context) error
	HandlerGetTotal(c echo.Context) error
}

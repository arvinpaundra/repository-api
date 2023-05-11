package staff

import "github.com/labstack/echo/v4"

type StaffController interface {
	HandlerRegister(c echo.Context) error
	HandlerLogin(c echo.Context) error
	HandlerUpdateStaff(c echo.Context) error
	HandlerFindAllStaffs(c echo.Context) error
	HandlerFindStaffById(c echo.Context) error
	HandlerUploadSignature(c echo.Context) error
}

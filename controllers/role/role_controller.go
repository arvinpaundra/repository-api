package role

import "github.com/labstack/echo/v4"

type RoleController interface {
	HandlerCreateRole(c echo.Context) error
	HandlerUpdateRole(c echo.Context) error
	HandlerDeleteRole(c echo.Context) error
	HandlerFindAllRoles(c echo.Context) error
	HandlerFindRoleById(c echo.Context) error
}

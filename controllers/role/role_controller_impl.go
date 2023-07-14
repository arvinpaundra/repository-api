package role

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/role/request"
	"github.com/arvinpaundra/repository-api/services/role"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type RoleControllerImpl struct {
	roleService role.RoleService
}

func NewRoleController(roleService role.RoleService) RoleController {
	return RoleControllerImpl{
		roleService: roleService,
	}
}

func (ctrl RoleControllerImpl) HandlerCreateRole(c echo.Context) error {
	var req request.CreateRoleRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.roleService.Create(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl RoleControllerImpl) HandlerUpdateRole(c echo.Context) error {
	roleId := c.Param("roleId")

	var req request.UpdateRoleRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.roleService.Update(c.Request().Context(), req, roleId)

	if err != nil {
		switch err {
		case utils.ErrRoleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl RoleControllerImpl) HandlerFindAllRoles(c echo.Context) error {
	visibility := c.QueryParam("visibility")

	roles, err := ctrl.roleService.FindAll(c.Request().Context(), visibility)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(roles, nil))
}

func (ctrl RoleControllerImpl) HandlerFindRoleById(c echo.Context) error {
	roleId := c.Param("roleId")

	role, err := ctrl.roleService.FindById(c.Request().Context(), roleId)

	if err != nil {
		switch err {
		case utils.ErrRoleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(role))
}

func (ctrl RoleControllerImpl) HandlerDeleteRole(c echo.Context) error {
	panic("not implemented")
}

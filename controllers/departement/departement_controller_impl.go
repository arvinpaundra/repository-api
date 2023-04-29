package departement

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/departement/request"
	"github.com/arvinpaundra/repository-api/services/departement"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type DepartementControllerImpl struct {
	departementService departement.DepartementService
}

func NewDepartementController(departementService departement.DepartementService) DepartementController {
	return DepartementControllerImpl{
		departementService: departementService,
	}
}

func (ctrl DepartementControllerImpl) HandlerCreateDepartement(c echo.Context) error {
	var req request.CreateDepartementRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.departementService.Create(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl DepartementControllerImpl) HandlerUpdateDepartement(c echo.Context) error {
	departementId := c.Param("departementId")
	var req request.UpdateDepartementRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.departementService.Update(c.Request().Context(), req, departementId)

	if err != nil {
		switch err {
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl DepartementControllerImpl) HandlerFindAllDepartements(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	departements, totalRows, totalPages, err := ctrl.departementService.FindAll(c.Request().Context(), keyword, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(departements, pagination))
}

func (ctrl DepartementControllerImpl) HandlerFindDepartementById(c echo.Context) error {
	departementId := c.Param("departementId")

	departement, err := ctrl.departementService.FindById(c.Request().Context(), departementId)

	if err != nil {
		switch err {
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(departement))
}

func (ctrl DepartementControllerImpl) HandlerDeleteDepartement(c echo.Context) error {
	panic("not implemented")
}

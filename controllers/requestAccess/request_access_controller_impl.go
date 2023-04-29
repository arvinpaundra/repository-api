package request_access

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/requestAccess/request"
	requestAccess "github.com/arvinpaundra/repository-api/services/requestAccess"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type RequestAccessControllerImpl struct {
	requestAccessService requestAccess.RequestAccessService
}

func NewRequestAccessController(requestAccessService requestAccess.RequestAccessService) RequestAccessController {
	return RequestAccessControllerImpl{
		requestAccessService: requestAccessService,
	}
}

func (ctrl RequestAccessControllerImpl) HandlerUpdateRequestAccess(c echo.Context) error {
	requestAccessId := c.Param("requestAccessId")
	var req request.UpdateRequestAccessRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.requestAccessService.Update(c.Request().Context(), req, requestAccessId)

	if err != nil {
		switch err {
		case utils.ErrRequestAccessNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl RequestAccessControllerImpl) HandlerFindAllRequestAccesses(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	status := c.QueryParam("status")

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	requestAccesses, totalRows, totalPages, err := ctrl.requestAccessService.FindAll(c.Request().Context(), keyword, status, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(requestAccesses, pagination))
}

func (ctrl RequestAccessControllerImpl) HandlerFindRequestAccessById(c echo.Context) error {
	requestAccessId := c.Param("requestAccessId")

	requestAccess, err := ctrl.requestAccessService.FindById(c.Request().Context(), requestAccessId)

	if err != nil {
		switch err {
		case utils.ErrRequestAccessNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(requestAccess))
}

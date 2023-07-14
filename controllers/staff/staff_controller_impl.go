package staff

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/staff/request"
	"github.com/arvinpaundra/repository-api/services/staff"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type StaffControllerImpl struct {
	staffService staff.StaffService
}

func NewStaffController(staffService staff.StaffService) StaffController {
	return StaffControllerImpl{
		staffService: staffService,
	}
}

func (ctrl StaffControllerImpl) HandlerRegister(c echo.Context) error {
	var req request.RegisterStaffRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.staffService.Register(c.Request().Context(), req)

	if err != nil {
		switch err {
		case utils.ErrEmailAlreadyUsed:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
				"email": err.Error(),
			}))
		case utils.ErrRoleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl StaffControllerImpl) HandlerLogin(c echo.Context) error {
	var req request.LoginStaffRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	token, err := ctrl.staffService.Login(c.Request().Context(), req)

	if err != nil {
		switch err {
		case utils.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(utils.ErrAuthenticationFailed.Error()))
		case utils.ErrStaffNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrAccountNotActivated:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(map[string]string{
		"token": token,
	}))
}

func (ctrl StaffControllerImpl) HandlerUpdateStaff(c echo.Context) error {
	var req request.UpdateStaffRequest

	staffId := c.Param("staffId")

	avatar, _ := c.FormFile("avatar")

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.staffService.Update(c.Request().Context(), req, avatar, staffId)

	if err != nil {
		switch err {
		case utils.ErrStaffNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrRoleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrEmailAlreadyUsed:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
				"email": err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl StaffControllerImpl) HandlerFindAllStaffs(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	roleId := c.QueryParam("role_id")

	query := request.StaffRequestQuery{
		Keyword: keyword,
		RoleId:  roleId,
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	staffs, totalRows, totalPages, err := ctrl.staffService.FindAll(c.Request().Context(), query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(staffs, pagination))
}

func (ctrl StaffControllerImpl) HandlerFindStaffById(c echo.Context) error {
	staffId := c.Param("staffId")

	staff, err := ctrl.staffService.FindById(c.Request().Context(), staffId)

	if err != nil {
		switch err {
		case utils.ErrStaffNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(staff))
}

func (ctrl StaffControllerImpl) HandlerUploadSignature(c echo.Context) error {
	staffId := c.Param("staffId")

	signature, _ := c.FormFile("signature")

	if signature == nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"signature": "Bagian ini wajib diisi",
		}))
	}

	if signature.Size > 1000000 {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"signature": "Max. size is 1MB",
		}))
	}

	err := ctrl.staffService.UploadSignature(c.Request().Context(), signature, staffId)

	if err != nil {
		switch err {
		case utils.ErrStaffNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

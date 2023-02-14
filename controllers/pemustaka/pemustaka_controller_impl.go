package pemustaka

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/request"
	"github.com/arvinpaundra/repository-api/services/pemustaka"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type PemustakaControllerImpl struct {
	pemustakaService pemustaka.PemustakaService
}

func NewPemustakaController(pemustakaService pemustaka.PemustakaService) PemustakaController {
	return PemustakaControllerImpl{
		pemustakaService: pemustakaService,
	}
}

func (ctrl PemustakaControllerImpl) HandlerRegister(c echo.Context) error {
	var req request.RegisterPemustakaRequest

	supportEvidence, _ := c.FormFile("support_evidence")

	if supportEvidence == nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"support_evidence": "This field is required",
		}))
	}

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.pemustakaService.Register(c.Request().Context(), req, supportEvidence)

	if err != nil {
		switch err {
		case utils.ErrEmailAlreadyUsed:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(err.Error()))
		case utils.ErrStudyProgramNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrRoleNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl PemustakaControllerImpl) HandlerLogin(c echo.Context) error {
	var req request.LoginPemustakaRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	token, err := ctrl.pemustakaService.Login(c.Request().Context(), req)

	if err != nil {
		switch err {
		case utils.ErrUserNotFound:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(utils.ErrAuthenticationFailed.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrWaitingForAcceptance:
			return c.JSON(http.StatusConflict, helper.ConflictResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(map[string]string{
		"token": token,
	}))
}

func (ctrl PemustakaControllerImpl) HandlerUpdatePemustaka(c echo.Context) error {
	var req request.UpdatePemustakaRequest

	pemustakaId := c.Param("pemustakaId")

	avatar, _ := c.FormFile("avatar")

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.pemustakaService.Update(c.Request().Context(), req, avatar, pemustakaId)

	if err != nil {
		switch err {
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrStudyProgramNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl PemustakaControllerImpl) HandlerFindAllPemustaka(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	roleId := c.QueryParam("role_id")
	departementId := c.QueryParam("departement_id")
	isCollectedFinalProject := c.QueryParam("is_collected_final_project")
	yearGen := c.QueryParam("year_gen")

	query := request.PemustakaRequestQuery{
		Keyword:                 keyword,
		RoleId:                  roleId,
		DepartementId:           departementId,
		IsCollectedFinalProject: isCollectedFinalProject,
		YearGen:                 yearGen,
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.limit": "Invalid number format",
		}))
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.page": "Invalid number format",
		}))
	}

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	pemustaka, totalRows, totalPages, err := ctrl.pemustakaService.FindAll(c.Request().Context(), query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(pemustaka, pagination))
}

func (ctrl PemustakaControllerImpl) HandlerFindPemustakaById(c echo.Context) error {
	pemustakaId := c.Param("pemustakaId")

	pemustaka, err := ctrl.pemustakaService.FindById(c.Request().Context(), pemustakaId)

	if err != nil {
		switch err {
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(pemustaka))
}

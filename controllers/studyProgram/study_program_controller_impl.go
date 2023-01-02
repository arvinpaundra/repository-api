package study_program

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/studyProgram/request"
	studyProgram "github.com/arvinpaundra/repository-api/services/studyProgram"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type StudyProgramControllerImpl struct {
	studyProgramService studyProgram.StudyProgramService
}

func NewStudyProgramController(studyProgramService studyProgram.StudyProgramService) StudyProgramController {
	return StudyProgramControllerImpl{
		studyProgramService: studyProgramService,
	}
}

func (ctrl StudyProgramControllerImpl) HandlerCreateStudyProgram(c echo.Context) error {
	var req request.CreateStudyProgramRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.studyProgramService.Create(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl StudyProgramControllerImpl) HandlerUpdateStudyProgram(c echo.Context) error {
	studyProgramId := c.Param("studyProgramId")
	var req request.UpdateStudyProgramRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.studyProgramService.Update(c.Request().Context(), req, studyProgramId)

	if err != nil {
		switch err {
		case utils.ErrStudyProgramNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl StudyProgramControllerImpl) HandlerFindAllStudyPrograms(c echo.Context) error {
	keyword := c.QueryParam("keyword")

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

	studyPrograms, totalRows, totalPages, err := ctrl.studyProgramService.FindAll(c.Request().Context(), keyword, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(studyPrograms, pagination))
}

func (ctrl StudyProgramControllerImpl) HandlerFindStudyProgramById(c echo.Context) error {
	studyProgramId := c.Param("studyProgramId")

	studyProgram, err := ctrl.studyProgramService.FindById(c.Request().Context(), studyProgramId)

	if err != nil {
		switch err {
		case utils.ErrStudyProgramNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(studyProgram))
}

func (ctrl StudyProgramControllerImpl) HandlerDeleteStudyProgram(c echo.Context) error {
	panic("not implemented")
}

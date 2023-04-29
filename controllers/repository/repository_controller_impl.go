package repository

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/repository/request"
	"github.com/arvinpaundra/repository-api/services/repository"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type RepositoryControllerImpl struct {
	repositoryService repository.RepositoryService
}

func NewRepositoryController(repositoryService repository.RepositoryService) RepositoryController {
	return RepositoryControllerImpl{
		repositoryService: repositoryService,
	}
}

func (ctrl RepositoryControllerImpl) HandlerCreateFinalProjectReport(c echo.Context) error {
	var req request.CreateFinalProjectReportRequest
	validationErrors := make(helper.ValidationError)

	validityPage, _ := c.FormFile("validity_page")

	if validityPage == nil {
		validationErrors["validity_page"] = "This field is required"
	}

	coverAndListContent, _ := c.FormFile("cover_and_list_content")

	if coverAndListContent == nil {
		validationErrors["cover_and_list_content"] = "This field is required"
	}

	chpOne, _ := c.FormFile("chp_one")

	if chpOne == nil {
		validationErrors["chp_one"] = "This field is required"
	}

	chpTwo, _ := c.FormFile("chp_two")

	if chpTwo == nil {
		validationErrors["chp_two"] = "This field is required"
	}

	chpThree, _ := c.FormFile("chp_three")

	if chpThree == nil {
		validationErrors["chp_three"] = "This field is required"
	}

	chpFour, _ := c.FormFile("chp_four")

	if chpFour == nil {
		validationErrors["chp_four"] = "This field is required"
	}

	chpFive, _ := c.FormFile("chp_five")

	if chpFive == nil {
		validationErrors["chp_five"] = "This field is required"
	}

	bibliography, _ := c.FormFile("bibliography")

	if bibliography == nil {
		validationErrors["bibliography"] = "This field is required"
	}

	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(validationErrors))
	}

	files := request.RepositoryInputFiles{
		ValidityPage:        validityPage,
		CoverAndListContent: coverAndListContent,
		ChpOne:              chpOne,
		ChpTwo:              chpTwo,
		ChpThree:            chpThree,
		ChpFour:             chpFour,
		ChpFive:             chpFive,
		Bibliography:        bibliography,
	}

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.repositoryService.CreateFinalProjectReport(c.Request().Context(), req, files)

	if err != nil {
		switch err {
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrCategoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrAlreadyCollectedFinalProject:
			return c.JSON(http.StatusNotFound, helper.BadRequestResponse([]string{err.Error()}))
		case utils.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrExaminerNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl RepositoryControllerImpl) HandlerCreateInternshipReport(c echo.Context) error {
	var req request.CreateInternshipReportRequest
	validationErrors := make(helper.ValidationError)

	validityPage, _ := c.FormFile("validity_page")

	if validityPage == nil {
		validationErrors["validity_page"] = "This field is required"
	}

	coverAndListContent, _ := c.FormFile("cover_and_list_content")

	if coverAndListContent == nil {
		validationErrors["cover_and_list_content"] = "This field is required"
	}

	chpOne, _ := c.FormFile("chp_one")

	if chpOne == nil {
		validationErrors["chp_one"] = "This field is required"
	}

	chpTwo, _ := c.FormFile("chp_two")

	if chpTwo == nil {
		validationErrors["chp_two"] = "This field is required"
	}

	chpThree, _ := c.FormFile("chp_three")

	if chpThree == nil {
		validationErrors["chp_three"] = "This field is required"
	}

	chpFour, _ := c.FormFile("chp_four")

	if chpFour == nil {
		validationErrors["chp_four"] = "This field is required"
	}

	chpFive, _ := c.FormFile("chp_five")

	if chpFive == nil {
		validationErrors["chp_five"] = "This field is required"
	}

	bibliography, _ := c.FormFile("bibliography")

	if bibliography == nil {
		validationErrors["bibliography"] = "This field is required"
	}

	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(validationErrors))
	}

	files := request.RepositoryInputFiles{
		ValidityPage:        validityPage,
		CoverAndListContent: coverAndListContent,
		ChpOne:              chpOne,
		ChpTwo:              chpTwo,
		ChpThree:            chpThree,
		ChpFour:             chpFour,
		ChpFive:             chpFive,
		Bibliography:        bibliography,
	}

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.repositoryService.CreateInternshipReport(c.Request().Context(), req, files)

	if err != nil {
		switch err {
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrAlreadyCollectedInternshipReport:
			return c.JSON(http.StatusNotFound, helper.BadRequestResponse([]string{err.Error()}))
		case utils.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusNotFound, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusNotFound, helper.SuccessCreatedResponse())
}

func (ctrl RepositoryControllerImpl) HandlerCreateResearchReport(c echo.Context) error {
	var req request.CreateResearchReportRequest
	validationErrors := make(helper.ValidationError)

	validityPage, _ := c.FormFile("validity_page")

	if validityPage == nil {
		validationErrors["validity_page"] = "This field is required"
	}

	coverAndListContent, _ := c.FormFile("cover_and_list_content")

	if coverAndListContent == nil {
		validationErrors["cover_and_list_content"] = "This field is required"
	}

	chpOne, _ := c.FormFile("chp_one")

	if chpOne == nil {
		validationErrors["chp_one"] = "This field is required"
	}

	chpTwo, _ := c.FormFile("chp_two")

	if chpTwo == nil {
		validationErrors["chp_two"] = "This field is required"
	}

	chpThree, _ := c.FormFile("chp_three")

	if chpThree == nil {
		validationErrors["chp_three"] = "This field is required"
	}

	chpFour, _ := c.FormFile("chp_four")

	if chpFour == nil {
		validationErrors["chp_four"] = "This field is required"
	}

	chpFive, _ := c.FormFile("chp_five")

	if chpFive == nil {
		validationErrors["chp_five"] = "This field is required"
	}

	bibliography, _ := c.FormFile("bibliography")

	if bibliography == nil {
		validationErrors["bibliography"] = "This field is required"
	}

	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(validationErrors))
	}

	files := request.RepositoryInputFiles{
		ValidityPage:        validityPage,
		CoverAndListContent: coverAndListContent,
		ChpOne:              chpOne,
		ChpTwo:              chpTwo,
		ChpThree:            chpThree,
		ChpFour:             chpFour,
		ChpFive:             chpFive,
		Bibliography:        bibliography,
	}

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.repositoryService.CreateResearchReport(c.Request().Context(), req, files)

	if err != nil {
		switch err {
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusNotFound, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusNotFound, helper.SuccessCreatedResponse())
}

func (ctrl RepositoryControllerImpl) HandlerDeleteRepository(c echo.Context) error {
	repositoryId := c.Param("repositoryId")

	err := ctrl.repositoryService.Delete(c.Request().Context(), repositoryId)

	if err != nil {
		switch err {
		case utils.ErrRepositoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl RepositoryControllerImpl) HandlerFindAllRepositories(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	collectionId := c.QueryParam("collection_id")
	departementId := c.QueryParam("departement_id")
	improvement := c.QueryParam("improvement")

	sort := c.QueryParam("sort")
	if sort != "created_at ASC" && sort != "created_at DESC" {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.sort": "Invalid value of sort. Only accept created_at ASC and created_at DESC",
		}))
	}

	query := request.RepositoryRequestQuery{
		Keyword:       keyword,
		CollectionId:  collectionId,
		DepartementId: departementId,
		Improvement:   improvement,
		Sort:          sort,
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	repositories, totalRows, totalPages, err := ctrl.repositoryService.FindAll(c.Request().Context(), query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(repositories, pagination))
}

func (ctrl RepositoryControllerImpl) HandlerFindRepositoryById(c echo.Context) error {
	repositoryId := c.Param("repositoryId")

	repository, err := ctrl.repositoryService.FindById(c.Request().Context(), repositoryId)

	if err != nil {
		switch err {
		case utils.ErrRepositoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(repository))
}

func (ctrl RepositoryControllerImpl) HandlerFindByAuthorId(c echo.Context) error {
	pemustakaId := c.Param("pemustakaId")

	keyword := c.QueryParam("keyword")
	collectionId := c.QueryParam("collection_id")
	departementId := c.QueryParam("departement_id")
	improvement := c.QueryParam("improvement")
	status := c.QueryParam("status")

	sort := c.QueryParam("sort")
	if sort != "created_at ASC" && sort != "created_at DESC" {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.sort": "Invalid value of sort. Only accept created_at ASC and created_at DESC",
		}))
	}

	query := request.RepositoryRequestQuery{
		Keyword:       keyword,
		CollectionId:  collectionId,
		DepartementId: departementId,
		Improvement:   improvement,
		Status:        status,
		Sort:          sort,
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	repositories, totalRows, totalPages, err := ctrl.repositoryService.FindByAuthorId(c.Request().Context(), pemustakaId, query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		switch err {
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(repositories, pagination))
}

func (ctrl RepositoryControllerImpl) HandlerFindByMentorId(c echo.Context) error {
	pemustakaId := c.Param("pemustakaId")

	keyword := c.QueryParam("keyword")
	collectionId := c.QueryParam("collection_id")
	departementId := c.QueryParam("departement_id")
	improvement := c.QueryParam("improvement")
	status := c.QueryParam("status")

	sort := c.QueryParam("sort")
	if sort != "created_at ASC" && sort != "created_at DESC" {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.sort": "Invalid value of sort. Only accept created_at ASC and created_at DESC",
		}))
	}

	query := request.RepositoryRequestQuery{
		Keyword:       keyword,
		CollectionId:  collectionId,
		DepartementId: departementId,
		Improvement:   improvement,
		Status:        status,
		Sort:          sort,
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	repositories, totalRows, totalPages, err := ctrl.repositoryService.FindByMentorId(c.Request().Context(), pemustakaId, query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		switch err {
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(repositories, pagination))
}

func (ctrl RepositoryControllerImpl) HandlerFindByExaminerId(c echo.Context) error {
	pemustakaId := c.Param("pemustakaId")

	keyword := c.QueryParam("keyword")
	collectionId := c.QueryParam("collection_id")
	departementId := c.QueryParam("departement_id")
	improvement := c.QueryParam("improvement")
	status := c.QueryParam("status")

	sort := c.QueryParam("sort")
	if sort != "created_at ASC" && sort != "created_at DESC" {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.sort": "Invalid value of sort. Only accept created_at ASC and created_at DESC",
		}))
	}

	query := request.RepositoryRequestQuery{
		Keyword:       keyword,
		CollectionId:  collectionId,
		DepartementId: departementId,
		Improvement:   improvement,
		Status:        status,
		Sort:          sort,
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	repositories, totalRows, totalPages, err := ctrl.repositoryService.FindByExaminerId(c.Request().Context(), pemustakaId, query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		switch err {
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(repositories, pagination))
}

func (ctrl RepositoryControllerImpl) HandlerFindByCollectionId(c echo.Context) error {
	collectionId := c.Param("collectionId")

	keyword := c.QueryParam("keyword")
	departementId := c.QueryParam("departement_id")
	improvement := c.QueryParam("improvement")
	status := c.QueryParam("status")

	sort := c.QueryParam("sort")
	if sort != "created_at ASC" && sort != "created_at DESC" {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.sort": "Invalid value of sort. Only accept created_at ASC and created_at DESC",
		}))
	}

	query := request.RepositoryRequestQuery{
		Keyword:       keyword,
		DepartementId: departementId,
		Improvement:   improvement,
		Status:        status,
		Sort:          sort,
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	repositories, totalRows, totalPages, err := ctrl.repositoryService.FindByCollectionId(c.Request().Context(), collectionId, query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		switch err {
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(repositories, pagination))
}

func (ctrl RepositoryControllerImpl) HandlerFindByDepartementId(c echo.Context) error {
	departementId := c.Param("departement_id")

	keyword := c.QueryParam("keyword")
	collectionId := c.QueryParam("collection_id")
	improvement := c.QueryParam("improvement")
	status := c.QueryParam("status")

	sort := c.QueryParam("sort")
	if sort != "created_at ASC" && sort != "created_at DESC" {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
			"request.query.sort": "Invalid value of sort. Only accept created_at ASC and created_at DESC",
		}))
	}

	query := request.RepositoryRequestQuery{
		Keyword:      keyword,
		CollectionId: collectionId,
		Improvement:  improvement,
		Status:       status,
		Sort:         sort,
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	pagination := &helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	repositories, totalRows, totalPages, err := ctrl.repositoryService.FindByDepartementId(c.Request().Context(), departementId, query, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		switch err {
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessOKResponseWithPagination(repositories, pagination))
}

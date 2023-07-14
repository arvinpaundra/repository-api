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

	files := request.RepositoryInputFiles{
		"validity_page":          nil,
		"cover_and_list_content": nil,
		"chp_one":                nil,
		"chp_two":                nil,
		"chp_three":              nil,
		"chp_four":               nil,
		"chp_five":               nil,
		"bibliography":           nil,
	}

	for key := range files {
		file, _ := c.FormFile(key)

		if file == nil {
			validationErrors[key] = "Bagian ini wajib diisi"
		} else {
			files[key] = file
		}
	}

	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(validationErrors))
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

	files := request.RepositoryInputFiles{
		"validity_page":          nil,
		"cover_and_list_content": nil,
		"chp_one":                nil,
		"chp_two":                nil,
		"chp_three":              nil,
		"chp_four":               nil,
		"chp_five":               nil,
		"bibliography":           nil,
	}

	for key := range files {
		file, _ := c.FormFile(key)

		if file == nil {
			validationErrors[key] = "Bagian ini wajib diisi"
		} else {
			files[key] = file
		}
	}

	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(validationErrors))
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

	files := request.RepositoryInputFiles{
		"validity_page":          nil,
		"cover_and_list_content": nil,
		"chp_one":                nil,
		"chp_two":                nil,
		"chp_three":              nil,
		"chp_four":               nil,
		"chp_five":               nil,
		"bibliography":           nil,
	}

	for key := range files {
		file, _ := c.FormFile(key)

		if file == nil {
			validationErrors[key] = "Bagian ini wajib diisi"
		} else {
			files[key] = file
		}
	}

	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(validationErrors))
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

func (ctrl RepositoryControllerImpl) HandlerUpdateFinalProjectReport(c echo.Context) error {
	var req request.UpdateFinalProjectReportRequest

	repositoryId := c.Param("repositoryId")

	files := request.RepositoryInputFiles{
		"validity_page":          nil,
		"cover_and_list_content": nil,
		"chp_one":                nil,
		"chp_two":                nil,
		"chp_three":              nil,
		"chp_four":               nil,
		"chp_five":               nil,
		"bibliography":           nil,
	}

	for key := range files {
		file, _ := c.FormFile(key)

		if file != nil {
			files[key] = file
		}
	}

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.repositoryService.UpdateFinalProjectReport(c.Request().Context(), req, files, repositoryId)

	if err != nil {
		switch err {
		case utils.ErrRepositoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrCategoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrExaminerNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDocumentsNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl RepositoryControllerImpl) HandlerUpdateInternshipReport(c echo.Context) error {
	var req request.UpdateInternshipReportRequest

	repositoryId := c.Param("repositoryId")

	files := request.RepositoryInputFiles{
		"validity_page":          nil,
		"cover_and_list_content": nil,
		"chp_one":                nil,
		"chp_two":                nil,
		"chp_three":              nil,
		"chp_four":               nil,
		"chp_five":               nil,
		"bibliography":           nil,
	}

	for key := range files {
		file, _ := c.FormFile(key)

		if file != nil {
			files[key] = file
		}
	}

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.repositoryService.UpdateInternshipReport(c.Request().Context(), req, files, repositoryId)

	if err != nil {
		switch err {
		case utils.ErrRepositoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrMentorNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDocumentsNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl RepositoryControllerImpl) HandlerUpdateResearchReport(c echo.Context) error {
	var req request.UpdateResearchReportRequest

	repositoryId := c.Param("repositoryId")

	files := request.RepositoryInputFiles{
		"validity_page":          nil,
		"cover_and_list_content": nil,
		"chp_one":                nil,
		"chp_two":                nil,
		"chp_three":              nil,
		"chp_four":               nil,
		"chp_five":               nil,
		"bibliography":           nil,
	}

	for key := range files {
		file, _ := c.FormFile(key)

		if file != nil {
			files[key] = file
		}
	}

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.repositoryService.UpdateResearchReport(c.Request().Context(), req, files, repositoryId)

	if err != nil {
		switch err {
		case utils.ErrRepositoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDepartementNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrDocumentsNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
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
	categoryId := c.QueryParam("category_id")
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
		CategoryId:    categoryId,
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
	categoryId := c.QueryParam("category_id")
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
		CategoryId:    categoryId,
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
	categoryId := c.QueryParam("category_id")
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
		CategoryId:    categoryId,
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
	categoryId := c.QueryParam("category_id")
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
		CategoryId:    categoryId,
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
	categoryId := c.QueryParam("category_id")
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
		CategoryId:    categoryId,
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
	categoryId := c.QueryParam("category_id")
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
		CategoryId:   categoryId,
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

func (ctrl RepositoryControllerImpl) HandlerGetTotalRepository(c echo.Context) error {
	status := c.QueryParam("status")

	totalRepository, err := ctrl.repositoryService.GetTotal(c.Request().Context(), status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(map[string]interface{}{
		"total_repositories": totalRepository,
	}))
}

func (ctrl RepositoryControllerImpl) HandlerConfirmRepository(c echo.Context) error {
	var req request.ConfirmRequest

	repositoryId := c.Param("repositoryId")

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.repositoryService.Confirm(c.Request().Context(), req, repositoryId)

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

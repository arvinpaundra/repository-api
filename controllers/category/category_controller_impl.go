package category

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/category/request"
	"github.com/arvinpaundra/repository-api/services/category"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type CategoryControllerImpl struct {
	categoryService category.CategoryService
}

func NewCategoryController(categoryService category.CategoryService) CategoryController {
	return CategoryControllerImpl{
		categoryService: categoryService,
	}
}

func (ctrl CategoryControllerImpl) HandlerCreateCategory(c echo.Context) error {
	var req request.CreateCategoryRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.categoryService.Create(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl CategoryControllerImpl) HandlerUpdateCategory(c echo.Context) error {
	categoryId := c.Param("categoryId")

	var req request.UpdateCategoryRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.categoryService.Update(c.Request().Context(), req, categoryId)

	if err != nil {
		switch err {
		case utils.ErrCategoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(nil, nil))
}

func (ctrl CategoryControllerImpl) HandlerDeleteCategory(c echo.Context) error {
	panic("not implemented")
}

func (ctrl CategoryControllerImpl) HandlerFindAllCategories(c echo.Context) error {
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

	categories, totalRows, totalPages, err := ctrl.categoryService.FindAll(c.Request().Context(), keyword, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessResponse(categories, pagination))
}

func (ctrl CategoryControllerImpl) HandlerFindCategoryById(c echo.Context) error {
	categoryId := c.Param("categoryId")

	category, err := ctrl.categoryService.FindById(c.Request().Context(), categoryId)

	if err != nil {
		switch err {
		case utils.ErrCategoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(category, nil))
}

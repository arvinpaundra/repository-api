package collection

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/collection/request"
	"github.com/arvinpaundra/repository-api/services/collection"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type CollectionControllerImpl struct {
	collectionService collection.CollectionService
}

func NewCollectionController(collectionService collection.CollectionService) CollectionController {
	return CollectionControllerImpl{
		collectionService: collectionService,
	}
}

func (ctrl CollectionControllerImpl) HandlerCreateCollection(c echo.Context) error {
	var req request.CreateCollectionRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.collectionService.Create(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessCreatedResponse())
}

func (ctrl CollectionControllerImpl) HandlerUpdateCollection(c echo.Context) error {
	collectionId := c.Param("collectionId")

	var req request.UpdateCollectionRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.collectionService.Update(c.Request().Context(), req, collectionId)

	if err != nil {
		switch err {
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(nil, nil))
}

func (ctrl CollectionControllerImpl) HandlerFindAllCollections(c echo.Context) error {
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

	collections, totalRows, totalPages, err := ctrl.collectionService.FindAll(c.Request().Context(), keyword, pagination.GetLimit(), pagination.GetOffset())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
	}

	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages

	return c.JSON(http.StatusOK, helper.SuccessResponse(collections, pagination))
}

func (ctrl CollectionControllerImpl) HandlerFindCollectionById(c echo.Context) error {
	collectionId := c.Param("collectionId")

	collection, err := ctrl.collectionService.FindById(c.Request().Context(), collectionId)

	if err != nil {
		switch err {
		case utils.ErrCollectionNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(collection, nil))
}

func (ctrl CollectionControllerImpl) HandlerDeleteCollectioin(c echo.Context) error {
	panic("not implemented")
}

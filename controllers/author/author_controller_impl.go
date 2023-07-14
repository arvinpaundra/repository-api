package author

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/services/author"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type AuthorControllerImpl struct {
	authorService author.AuthorService
}

func NewAuthorController(authorService author.AuthorService) AuthorController {
	return AuthorControllerImpl{
		authorService: authorService,
	}
}

func (ctrl AuthorControllerImpl) HandlerDeleteAuthor(c echo.Context) error {
	repositoryId := c.Param("repositoryId")
	pemustakaId := c.Param("pemustakaId")

	err := ctrl.authorService.Delete(c.Request().Context(), repositoryId, pemustakaId)

	if err != nil {
		switch err {
		case utils.ErrRepositoryNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

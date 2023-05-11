package identity_card

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/helper"
	identityCard "github.com/arvinpaundra/repository-api/services/identityCard"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type IdentityCardControllerImpl struct {
	identityCardService identityCard.IdentityCardService
}

func NewIdentityCardController(identityCardService identityCard.IdentityCardService) IdentityCardController {
	return IdentityCardControllerImpl{
		identityCardService: identityCardService,
	}
}

func (ctrl IdentityCardControllerImpl) HandlerGenerateIDCard(c echo.Context) error {
	pemustakaId := c.Param("pemustakaId")

	res, err := ctrl.identityCardService.Generate(c.Request().Context(), pemustakaId)

	if err != nil {
		switch err {
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.Blob(http.StatusOK, "image/png", res)
}

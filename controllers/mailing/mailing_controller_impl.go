package mailing

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/mailing/request"
	"github.com/arvinpaundra/repository-api/services/mailing"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type MailingControllerImpl struct {
	mailingService mailing.MailingService
}

func NewMailingController(mailingService mailing.MailingService) MailingController {
	return MailingControllerImpl{
		mailingService: mailingService,
	}
}

func (ctrl MailingControllerImpl) HandlerSendForgotPasswordWithExpirationToken(c echo.Context) error {
	var req request.SendEmailForgotPasswordRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.mailingService.SendForgotPasswordWithTokenExpiration(c.Request().Context(), req)

	if err != nil {
		switch err {
		case utils.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(utils.ErrEmailNotFound.Error()))
		case utils.ErrPemustakaNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		case utils.ErrTokenHasBeenSent:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
				"link": "link has been sent",
			}))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

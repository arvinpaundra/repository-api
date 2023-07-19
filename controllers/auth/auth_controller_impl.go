package auth

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/web/auth/request"
	"github.com/arvinpaundra/repository-api/services/auth"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	authService auth.AuthService
}

func NewAuthController(authService auth.AuthService) AuthController {
	return AuthControllerImpl{
		authService: authService,
	}
}

func (ctrl AuthControllerImpl) HandlerForgotPassword(c echo.Context) error {
	var req request.ForgotPasswordRequest

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.authService.ForgotPassword(c.Request().Context(), req)

	if err != nil {
		switch err {
		case utils.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(utils.ErrEmailNotFound.Error()))
		case utils.ErrTokenExpired:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
				"link": utils.ErrLinkExpired.Error(),
			}))
		case utils.ErrTokenNotMatch:
			return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(map[string]string{
				"token": utils.ErrTokenNotMatch.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

func (ctrl AuthControllerImpl) HandlerChangePassword(c echo.Context) error {
	var req request.ChangePasswordRequest

	userId := c.Param("userId")

	_ = c.Bind(&req)

	if err := helper.ValidateRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
	}

	err := ctrl.authService.ChangePassword(c.Request().Context(), userId, req)

	if err != nil {
		switch err {
		case utils.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, helper.NotFoundResponse(err.Error()))
		default:
			return c.JSON(http.StatusInternalServerError, helper.InternalServerErrorResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessOKResponse(nil))
}

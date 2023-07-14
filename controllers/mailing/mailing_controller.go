package mailing

import "github.com/labstack/echo/v4"

type MailingController interface {
	HandlerSendForgotPasswordWithExpirationToken(c echo.Context) error
}

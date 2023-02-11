package auth

import "github.com/labstack/echo/v4"

type AuthController interface {
	HandlerForgotPassword(c echo.Context) error
}

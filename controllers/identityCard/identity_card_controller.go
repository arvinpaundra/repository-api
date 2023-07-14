package identity_card

import "github.com/labstack/echo/v4"

type IdentityCardController interface {
	HandlerGenerateIDCard(c echo.Context) error
}

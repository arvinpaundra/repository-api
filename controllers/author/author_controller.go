package author

import "github.com/labstack/echo/v4"

type AuthorController interface {
	HandlerDeleteAuthor(c echo.Context) error
}

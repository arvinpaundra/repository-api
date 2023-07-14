package category

import "github.com/labstack/echo/v4"

type CategoryController interface {
	HandlerCreateCategory(c echo.Context) error
	HandlerUpdateCategory(c echo.Context) error
	HandlerDeleteCategory(c echo.Context) error
	HandlerFindAllCategories(c echo.Context) error
	HandlerFindCategoryById(c echo.Context) error
}

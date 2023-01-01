package controller

import "github.com/labstack/echo/v4"

type CollectionController interface {
	HandlerCreateCollection(c echo.Context) error
	HandlerUpdateCollection(c echo.Context) error
	HandlerDeleteCollectioin(c echo.Context) error
	HandlerFindAllCollections(c echo.Context) error
	HandlerFindCollectionById(c echo.Context) error
}

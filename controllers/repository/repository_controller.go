package repository

import "github.com/labstack/echo/v4"

type RepositoryController interface {
	HandlerCreateFinalProjectReport(c echo.Context) error
	HandlerCreateInternshipReport(c echo.Context) error
	HandlerCreateResearchReport(c echo.Context) error
	HandlerUpdateFinalProjectReport(c echo.Context) error
	HandlerUpdateInternshipReport(c echo.Context) error
	HandlerUpdateResearchReport(c echo.Context) error
	HandlerDeleteRepository(c echo.Context) error
	HandlerFindAllRepositories(c echo.Context) error
	HandlerFindRepositoryById(c echo.Context) error
	HandlerFindByAuthorId(c echo.Context) error
	HandlerFindByMentorId(c echo.Context) error
	HandlerFindByExaminerId(c echo.Context) error
	HandlerFindByCollectionId(c echo.Context) error
	HandlerFindByDepartementId(c echo.Context) error
	HandlerGetTotalRepository(c echo.Context) error
	HandlerConfirmRepository(c echo.Context) error
}

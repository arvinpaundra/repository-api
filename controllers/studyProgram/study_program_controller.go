package study_program

import "github.com/labstack/echo/v4"

type StudyProgramController interface {
	HandlerCreateStudyProgram(c echo.Context) error
	HandlerUpdateStudyProgram(c echo.Context) error
	HandlerDeleteStudyProgram(c echo.Context) error
	HandlerFindAllStudyPrograms(c echo.Context) error
	HandlerFindStudyProgramById(c echo.Context) error
}

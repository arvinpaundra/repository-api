package middlewares

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

func UploadRepositoryFiles() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			files := map[string]*multipart.FileHeader{
				"validity_page":          nil,
				"cover_and_list_content": nil,
				"chp_one":                nil,
				"chp_two":                nil,
				"chp_three":              nil,
				"chp_four":               nil,
				"chp_five":               nil,
				"bibliography":           nil,
			}

			for key := range files {
				file, _ := c.FormFile(key)

				if file != nil {
					files[key] = file
				}
			}

			validationErrors := make(helper.ValidationError)

			for key, file := range files {
				if file != nil {
					// check file size
					fileSize := file.Size

					if fileSize > 5000000 {
						validationErrors[key] = "Ukuran maks. 5MB"
					}

					// check file extension
					fileExtension := filepath.Ext(file.Filename)
					if fileExtension != ".pdf" {
						validationErrors[key] = utils.ErrFileTypeMustBePDF.Error()
					}
				}
			}

			if len(validationErrors) != 0 {
				return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(validationErrors))
			}

			return next(c)
		}
	}
}

func UploadAvatarValidator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := make(helper.ValidationError)

			file, _ := c.FormFile("avatar")

			if file != nil {
				// check file size
				fileSize := file.Size
				if fileSize > 1000000 {
					err["avatar"] = "Ukuran maks. 1MB"
				}

				// check file extension
				fileExtension := filepath.Ext(file.Filename)
				if fileExtension != ".jpg" && fileExtension != ".png" && fileExtension != ".jpeg" && fileExtension != ".webp" {
					err["avatar"] = "Hanya dapat mengunggah berkas berformat .jpg, .png, .jpeg, .webp"
				}
			}

			if len(err) != 0 {
				return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
			}

			return next(c)
		}
	}
}

func UploadSupportEvidence() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := make(helper.ValidationError)

			file, _ := c.FormFile("support_evidence")

			if file != nil {
				// check file size
				fileSize := file.Size
				if fileSize > 2000000 {
					err["support_evidence"] = "Ukuran maks. 2MB"
				}

				fileExtension := filepath.Ext(file.Filename)
				if fileExtension != ".jpg" && fileExtension != ".png" && fileExtension != ".jpeg" {
					err["support_evidence"] = "Hanya dapat mengunggah berkas berformat .jpg, .png, .jpeg"
				}
			}

			if len(err) != 0 {
				return c.JSON(http.StatusBadRequest, helper.BadRequestResponse(err))
			}

			return next(c)
		}
	}
}

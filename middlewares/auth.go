package middlewares

import (
	"net/http"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CheckRoles(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			payloads, err := utils.ExtractToken(c)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, helper.UnauthorizedResponse())
			}

			for _, role := range roles {
				switch {
				case payloads.Role == role:
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, helper.ForbiddenResponse())
		}
	}
}

func IsAuthenticated() echo.MiddlewareFunc {
	return echojwt.JWT([]byte(configs.GetConfig("JWT_SECRET")))
}

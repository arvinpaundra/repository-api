package utils

import (
	"time"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var jwtSecret = configs.GetConfig("JWT_SECRET")

type JWTCustomClaim struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id string, role string) (string, error) {
	claims := JWTCustomClaim{
		ID:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Local().Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func ExtractToken(c echo.Context) error {
	panic("implement me")
}

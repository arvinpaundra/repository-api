package utils

import (
	"strings"
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

func ExtractToken(c echo.Context) (*JWTCustomClaim, error) {
	tokenFromHeader := c.Request().Header.Get("Authorization")

	sanitizedTokenBearer := strings.Replace(tokenFromHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(sanitizedTokenBearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidTokenHeader
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)

		userId := claims["id"].(string)
		role := claims["role"].(string)

		customClaims := &JWTCustomClaim{
			ID:   userId,
			Role: role,
		}

		return customClaims, nil
	}

	return nil, ErrUnAuthorized
}

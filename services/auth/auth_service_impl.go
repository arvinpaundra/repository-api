package auth

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/arvinpaundra/repository-api/drivers/mysql/user"
	expirationToken "github.com/arvinpaundra/repository-api/drivers/redis/expirationToken"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/auth/request"
	"github.com/arvinpaundra/repository-api/utils"
)

type AuthServiceImpl struct {
	userRepository            user.UserRepository
	expirationTokenRepository expirationToken.ExpirationTokenRepository
}

func NewAuthService(
	userRepository user.UserRepository,
	expirationTokenRepository expirationToken.ExpirationTokenRepository,
) AuthService {
	return AuthServiceImpl{
		userRepository:            userRepository,
		expirationTokenRepository: expirationTokenRepository,
	}
}

func (service AuthServiceImpl) ForgotPassword(ctx context.Context, req request.ForgotPasswordRequest) error {
	decodedToken, err := base64.RawURLEncoding.DecodeString(req.Base64Token)

	if err != nil {
		return err
	}

	// get email from decodedToken
	email := strings.Split(string(decodedToken), "&")[1]

	if _, err := service.userRepository.FindByEmail(ctx, email); err != nil {
		return err
	}

	token, err := service.expirationTokenRepository.Get(ctx, email)

	if err != nil {
		return err
	}

	if token != string(decodedToken) {
		return utils.ErrTokenNotMatch
	}

	hashPassword := utils.HashPassword(req.Password)

	user := domain.User{
		Password: hashPassword,
	}

	if err := service.userRepository.Update(ctx, user, email); err != nil {
		return err
	}

	return nil
}

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
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	userRepository            user.UserRepository
	expirationTokenRepository expirationToken.ExpirationTokenRepository
	tx                        *gorm.DB
}

func NewAuthService(
	userRepository user.UserRepository,
	expirationTokenRepository expirationToken.ExpirationTokenRepository,
	tx *gorm.DB,
) AuthService {
	return AuthServiceImpl{
		userRepository:            userRepository,
		expirationTokenRepository: expirationTokenRepository,
		tx:                        tx,
	}
}

func (service AuthServiceImpl) ForgotPassword(ctx context.Context, req request.ForgotPasswordRequest) error {
	tx := service.tx.Begin()

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

	if err := service.userRepository.Update(ctx, tx, user, email); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}

func (service AuthServiceImpl) ChangePassword(ctx context.Context, userId string, req request.ChangePasswordRequest) error {
	tx := service.tx.Begin()

	user, err := service.userRepository.FindById(ctx, userId)

	if err != nil {
		return err
	}

	userDomain := domain.User{
		Password: utils.HashPassword(req.RepeatedPassword),
	}

	if err := service.userRepository.Update(ctx, tx, userDomain, user.Email); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}
		return err
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}

package auth_test

import (
	"context"
	"errors"
	"testing"

	userMock "github.com/arvinpaundra/repository-api/drivers/mysql/user/mocks"
	expTokenMock "github.com/arvinpaundra/repository-api/drivers/redis/expirationToken/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/auth/request"
	"github.com/arvinpaundra/repository-api/services/auth"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRepository            userMock.UserRepository
	expirationTokenRepository expTokenMock.ExpirationTokenRepository
	authService               auth.AuthService

	userDomain        domain.User
	forgotPasswordDTO request.ForgotPasswordRequest

	ctx   context.Context
	token string
	tx    *gorm.DB
)

func TestMain(m *testing.M) {
	authService = auth.NewAuthService(&userRepository, &expirationTokenRepository, tx)

	userDomain = domain.User{
		ID:       uuid.NewString(),
		Email:    "test@mail.com",
		Password: "12345678",
	}

	forgotPasswordDTO = request.ForgotPasswordRequest{
		Password:    "abcdefghijk",
		Base64Token: "NTViYTE2Y2M2ODU5NGM1MGEyOTViZDVjNDJlMGNmMjMmdGVzdEBtYWlsLmNvbQ",
	}

	ctx = context.Background()
	token = "55ba16cc68594c50a295bd5c42e0cf23&test@mail.com"

	m.Run()
}

func TestForgotPassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", ctx, userDomain.Email).Return(userDomain, nil).Once()

		expirationTokenRepository.Mock.On("Get", ctx, userDomain.Email).Return(token, nil).Once()

		userRepository.Mock.On("Update", ctx, mock.Anything, userDomain.Email).Return(nil).Once()

		err := authService.ForgotPassword(ctx, forgotPasswordDTO)

		assert.NoError(t, err)
	})

	// failed test
	t.Run("Failed | User email not found", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", ctx, userDomain.Email).Return(domain.User{}, utils.ErrEmailNotFound).Once()

		err := authService.ForgotPassword(ctx, forgotPasswordDTO)

		assert.Error(t, err)
	})

	t.Run("Failed | Token expired", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", ctx, userDomain.Email).Return(userDomain, nil).Once()

		expirationTokenRepository.Mock.On("Get", ctx, userDomain.Email).Return("", utils.ErrTokenExpired).Once()

		err := authService.ForgotPassword(ctx, forgotPasswordDTO)

		assert.Error(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", ctx, userDomain.Email).Return(userDomain, nil).Once()

		expirationTokenRepository.Mock.On("Get", ctx, userDomain.Email).Return(token, nil).Once()

		userRepository.Mock.On("Update", ctx, mock.Anything, userDomain.Email).Return(errors.New("error occurred")).Once()

		err := authService.ForgotPassword(ctx, forgotPasswordDTO)

		assert.Error(t, err)
	})
}

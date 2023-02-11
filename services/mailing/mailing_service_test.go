package mailing_test

import (
	"context"
	"errors"
	"testing"
	"time"

	pemustakaMock "github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka/mocks"
	userMock "github.com/arvinpaundra/repository-api/drivers/mysql/user/mocks"
	expTokenMock "github.com/arvinpaundra/repository-api/drivers/redis/expirationToken/mocks"
	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/mailing/request"
	"github.com/arvinpaundra/repository-api/services/mailing"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	expirationTokenRepository expTokenMock.ExpirationTokenRepository
	userRepository            userMock.UserRepository
	pemustakaRepository       pemustakaMock.PemustakaRepository
	mailingService            mailing.MailingService

	userDomain                 domain.User
	pemustakaDomain            domain.Pemustaka
	sendEmailForgotPasswordDTO request.SendEmailForgotPasswordRequest

	ctx   context.Context
	mail  helper.Mailing
	token string
)

func TestMain(m *testing.M) {
	mail = *helper.NewMailing()

	mailingService = mailing.NewMailingService(&expirationTokenRepository, &userRepository, &pemustakaRepository, mail)

	userDomain = domain.User{
		ID:       uuid.NewString(),
		Email:    "test@mail.com",
		Password: "12345678",
	}

	pemustakaDomain = domain.Pemustaka{
		ID:                      uuid.NewString(),
		UserId:                  userDomain.ID,
		StudyProgramId:          uuid.NewString(),
		DepartementId:           uuid.NewString(),
		RoleId:                  uuid.NewString(),
		MemberCode:              "ABCD",
		Fullname:                "test",
		IdentityNumber:          "000000000",
		YearGen:                 "0000",
		Gender:                  "test",
		Telp:                    "test",
		BirthDate:               "test",
		Address:                 "test",
		IsCollectedFinalProject: "0",
		IsActive:                "1",
		Avatar:                  "test",
	}

	sendEmailForgotPasswordDTO = request.SendEmailForgotPasswordRequest{
		Email: userDomain.Email,
	}

	ctx = context.Background()
	token = "55ba16cc68594c50a295bd5c42e0cf23&test@mail.com"

	m.Run()
}

func TestSendForgotPasswordWithExpirationToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expirationTokenRepository.Mock.On("Get", ctx, userDomain.Email).Return(token, utils.ErrTokenExpired).Once()

		userRepository.Mock.On("FindByEmail", ctx, userDomain.Email).Return(userDomain, nil).Once()

		expirationTokenRepository.Mock.On("Save", ctx, userDomain.Email, mock.Anything, 15*time.Minute).Return(nil).Once()

		mail.SendForgotPasswordMail(userDomain.Email, "Permintaan Reset Password Pengguna", mock.Anything)

		err := mailingService.SendForgotPasswordWithTokenExpiration(ctx, sendEmailForgotPasswordDTO)

		assert.NoError(t, err)
	})

	t.Run("Failed | Token has been sent", func(t *testing.T) {
		expirationTokenRepository.Mock.On("Get", ctx, userDomain.Email).Return(token, utils.ErrTokenHasBeenSent).Once()

		err := mailingService.SendForgotPasswordWithTokenExpiration(ctx, sendEmailForgotPasswordDTO)

		assert.Error(t, err)
	})

	t.Run("Failed | User email not found", func(t *testing.T) {
		expirationTokenRepository.Mock.On("Get", ctx, userDomain.Email).Return(token, utils.ErrTokenExpired).Once()

		userRepository.Mock.On("FindByEmail", ctx, userDomain.Email).Return(domain.User{}, utils.ErrEmailNotFound).Once()

		err := mailingService.SendForgotPasswordWithTokenExpiration(ctx, sendEmailForgotPasswordDTO)

		assert.Error(t, err)
	})

	t.Run("Failed | Error save token", func(t *testing.T) {
		expirationTokenRepository.Mock.On("Get", ctx, userDomain.Email).Return(token, utils.ErrTokenExpired).Once()

		userRepository.Mock.On("FindByEmail", ctx, userDomain.Email).Return(userDomain, nil).Once()

		expirationTokenRepository.Mock.On("Save", ctx, userDomain.Email, mock.Anything, 15*time.Minute).Return(errors.New("error occurred")).Once()

		err := mailingService.SendForgotPasswordWithTokenExpiration(ctx, sendEmailForgotPasswordDTO)

		assert.Error(t, err)
	})
}

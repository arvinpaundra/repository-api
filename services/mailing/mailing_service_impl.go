package mailing

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/drivers/mysql/user"
	expirationToken "github.com/arvinpaundra/repository-api/drivers/redis/expirationToken"
	"github.com/arvinpaundra/repository-api/helper/mailing"
	"github.com/arvinpaundra/repository-api/models/web/mailing/request"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
)

type MailingServiceImpl struct {
	expirationTokenRepository expirationToken.ExpirationTokenRepository
	userRepository            user.UserRepository
	pemustakaRepository       pemustaka.PemustakaRepository
	mailing                   mailing.Mailing
}

func NewMailingService(
	expirationTokenRepository expirationToken.ExpirationTokenRepository,
	userRepository user.UserRepository,
	pemustakaRepository pemustaka.PemustakaRepository,
	mailing mailing.Mailing,
) MailingService {
	return MailingServiceImpl{
		expirationTokenRepository: expirationTokenRepository,
		userRepository:            userRepository,
		pemustakaRepository:       pemustakaRepository,
		mailing:                   mailing,
	}
}

func (service MailingServiceImpl) SendForgotPasswordWithTokenExpiration(ctx context.Context, req request.SendEmailForgotPasswordRequest) error {
	// check token to prevent user request too much
	_, err := service.expirationTokenRepository.Get(ctx, req.Email)

	if err != utils.ErrTokenExpired {
		return utils.ErrTokenHasBeenSent
	}

	user, err := service.userRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		return err
	}

	// get unique string from uuid with deleted hyphens
	uniqueString := strings.ReplaceAll(uuid.NewString(), "-", "")

	token := fmt.Sprintf("%s&%s", uniqueString, user.Email)

	if err := service.expirationTokenRepository.Save(ctx, req.Email, token, 15*time.Minute); err != nil {
		return err
	}

	// encode token to base64 format
	encodeTokenToBase64 := base64.RawURLEncoding.EncodeToString([]byte(token))

	mailErr := service.mailing.SendForgotPasswordMail(user.Email, "Permintaan Reset Password Pengguna", encodeTokenToBase64)

	if mailErr != nil {
		return mailErr
	}

	return nil
}

func (service MailingServiceImpl) SendRequestAccessAccepted(ctx context.Context, pemustakaId string) error {
	panic("not implemented")
}

func (service MailingServiceImpl) SendRepositoryAccepted(ctx context.Context, pemustakaId string) error {
	panic("not implemented")
}

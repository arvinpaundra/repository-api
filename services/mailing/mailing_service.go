package mailing

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/mailing/request"
)

type MailingService interface {
	SendForgotPasswordWithTokenExpiration(ctx context.Context, req request.SendEmailForgotPasswordRequest) error
	SendRequestAccessAccepted(ctx context.Context, pemustakaId string) error
	SendRepositoryAccepted(ctx context.Context, pemustakaId string) error
}

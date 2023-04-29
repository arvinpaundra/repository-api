package auth

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/auth/request"
)

type AuthService interface {
	ForgotPassword(ctx context.Context, req request.ForgotPasswordRequest) error
	ChangePassword(ctx context.Context, userId string, req request.ChangePasswordRequest) error
}

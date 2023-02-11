package expiration_token

import (
	"context"
	"time"
)

type ExpirationTokenRepository interface {
	Save(ctx context.Context, email string, token string, ttl time.Duration) error
	Get(ctx context.Context, email string) (string, error)
}

package expiration_token

import (
	"context"
	"time"

	"github.com/arvinpaundra/repository-api/utils"
	"github.com/go-redis/redis/v8"
)

type ExpirationTokenRepositoryImpl struct {
	conn *redis.Client
}

func NewRedisRepository(conn *redis.Client) ExpirationTokenRepository {
	return ExpirationTokenRepositoryImpl{
		conn: conn,
	}
}

func (repository ExpirationTokenRepositoryImpl) Save(ctx context.Context, email string, token string, ttl time.Duration) error {
	err := repository.conn.Set(ctx, email, token, ttl).Err()

	if err != nil {
		return err
	}

	return nil
}

func (repository ExpirationTokenRepositoryImpl) Get(ctx context.Context, email string) (string, error) {
	token, err := repository.conn.Get(ctx, email).Result()

	if err == redis.Nil {
		return "", utils.ErrTokenExpired
	} else if err != nil {
		return "", err
	}

	return token, nil
}

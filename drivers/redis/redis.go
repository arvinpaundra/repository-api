package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	USERNAME string
	HOST     string
	PORT     string
	PASSWORD string
	DB       string
}

// NewRedis create new instance for Redis
func NewRedis(username, password, host, port, db string) *Redis {
	return &Redis{
		USERNAME: username,
		PASSWORD: password,
		HOST:     host,
		PORT:     port,
		DB:       db,
	}
}

func (config *Redis) Init(ctx context.Context) *redis.Client {
	address := fmt.Sprintf("%s:%s", config.HOST, config.PORT)
	db, _ := strconv.Atoi(config.DB)

	client := redis.NewClient(&redis.Options{
		Username:  config.USERNAME,
		Addr:      address,
		Password:  config.PASSWORD,
		DB:        db,
		TLSConfig: &tls.Config{},
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("error when connect to redis server: %v", err)
	}

	log.Println("connected to redis server")

	return client
}

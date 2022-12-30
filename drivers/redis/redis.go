package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/go-redis/redis/v8"
)

type ConfigRedis struct {
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       string
}

func New() *ConfigRedis {
	return &ConfigRedis{
		REDIS_HOST:     configs.GetConfig("REDIS_HOST"),
		REDIS_PORT:     configs.GetConfig("REDIS_PORT"),
		REDIS_PASSWORD: configs.GetConfig("REDIS_PASSWORD"),
		REDIS_DB:       configs.GetConfig("REDIS_DB"),
	}
}

func (config *ConfigRedis) Init(ctx context.Context) *redis.Client {
	address := fmt.Sprintf("%s:%s", config.REDIS_HOST, config.REDIS_PORT)
	db, _ := strconv.Atoi(config.REDIS_DB)

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.REDIS_PASSWORD,
		DB:       db,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("error when connect to redis server: %v", err)
	}

	log.Println("connected to redis server")

	return client
}

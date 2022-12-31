package main

import (
	"context"
	"net/http"
	"time"

	"github.com/arvinpaundra/repository-api/app/routes"
	"github.com/arvinpaundra/repository-api/configs"
	driverMySQL "github.com/arvinpaundra/repository-api/drivers/mysql"
	driverRedis "github.com/arvinpaundra/repository-api/drivers/redis"
	"github.com/arvinpaundra/repository-api/middlewares"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	// init mysql
	mysql := driverMySQL.New()
	mysqldb := mysql.Init()

	// init redis
	redis := driverRedis.New()
	redisdb := redis.Init(ctx)

	e := echo.New()

	e.Use(middlewares.CORS())

	route := routes.RouteConfig{
		Echo:  e,
		MySQl: mysqldb,
		Redis: redisdb,
	}

	route.New()

	go func() {
		if err := e.Start(configs.GetConfig("APP_PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down server")
		}
	}()

	wait := utils.GracefulShutdown(ctx, 5*time.Second, map[string]utils.Operation{
		"mysql": func(ctx context.Context) error {
			return driverMySQL.CloseMySQL(mysqldb)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	<-wait
}

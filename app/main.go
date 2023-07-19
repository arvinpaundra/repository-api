package main

import (
	"context"
	"net/http"
	"time"

	"github.com/arvinpaundra/repository-api/app/routes"
	"github.com/arvinpaundra/repository-api/configs"
	driverMySQL "github.com/arvinpaundra/repository-api/drivers/mysql"
	driverRedis "github.com/arvinpaundra/repository-api/drivers/redis"
	"github.com/arvinpaundra/repository-api/helper/cloudinary"
	"github.com/arvinpaundra/repository-api/helper/mailing"
	"github.com/arvinpaundra/repository-api/middlewares"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	// init mysql
	mysql := driverMySQL.NewMySQL(configs.GetConfig("MYSQL_USERNAME"), configs.GetConfig("MYSQL_PASSWORD"), configs.GetConfig("MYSQL_HOST"), configs.GetConfig("MYSQL_PORT"), configs.GetConfig("MYSQL_DBNAME"))
	mysqldb := mysql.Init()

	// init redis
	redis := driverRedis.NewRedis(configs.GetConfig("REDIS_USERNAME"), configs.GetConfig("REDIS_PASSWORD"), configs.GetConfig("REDIS_HOST"), configs.GetConfig("REDIS_PORT"), configs.GetConfig("REDIS_DB"))
	redisdb := redis.Init(ctx)

	e := echo.New()

	e.Use(middlewares.CORS())

	// init mail service
	mail := mailing.NewMailing(configs.GetConfig("SMTP_HOST"), configs.GetConfig("SMTP_PORT"), configs.GetConfig("EMAIL"), configs.GetConfig("PASSWORD_EMAIL"), configs.GetConfig("EMAIL_SENDER_NAME"), configs.GetConfig("FE_BASE_URL"))
	cloudinary := cloudinary.NewCloudinary()

	route := routes.RouteConfig{
		Echo:       e,
		MySQl:      mysqldb,
		Redis:      redisdb,
		Mailing:    mail,
		Cloudinary: cloudinary,
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

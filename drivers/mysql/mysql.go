package mysql

import (
	"fmt"
	"log"

	"github.com/arvinpaundra/repository-api/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	MYSQL_DBNAME   string
	MYSQL_PORT     string
}

// init mysql configs
func New() *MySQLConfig {
	return &MySQLConfig{
		MYSQL_USERNAME: configs.GetConfig("MYSQL_USERNAME"),
		MYSQL_PASSWORD: configs.GetConfig("MYSQL_PASSWORD"),
		MYSQL_HOST:     configs.GetConfig("MYSQL_HOST"),
		MYSQL_DBNAME:   configs.GetConfig("MYSQL_DBNAME"),
		MYSQL_PORT:     configs.GetConfig("MYSQL_PORT"),
	}
}

func (config *MySQLConfig) Init() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MYSQL_USERNAME,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_DBNAME,
	)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("error connect to mysql: %v", err)
	}

	log.Println("connected to mysql")

	return db
}

// perform to close mysql database connection
func CloseMySQL(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Fatalf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Fatalf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection closed")

	return nil
}

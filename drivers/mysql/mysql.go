package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	USERNAME string
	PASSWORD string
	HOST     string
	DBNAME   string
	PORT     string
}

// NewMySQL create new instace for MySQL
func NewMySQL(username, password, host, port, dbName string) *MySQL {
	return &MySQL{
		USERNAME: username,
		PASSWORD: password,
		HOST:     host,
		DBNAME:   dbName,
		PORT:     port,
	}
}

func (config *MySQL) Init() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.USERNAME,
		config.PASSWORD,
		config.HOST,
		config.PORT,
		config.DBNAME,
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

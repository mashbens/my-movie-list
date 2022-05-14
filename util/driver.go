package util

import (
	"github.com/mashbens/my-movie-list/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	MySQL DatabaseDriver = "mysql"

	PostgreSQL DatabaseDriver = "postgres"

	Static DatabaseDriver = "static"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	MySQL *gorm.DB

	PostgreSQL *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Database.Driver {
	case "MySQL":
		db.Driver = MySQL
		db.MySQL = newMySQL(config)
	case "PostgreSQL":
		db.Driver = PostgreSQL
		db.PostgreSQL = newPostgreSQL(config)
	default:
		panic("Database driver not supported")
	}

	return &db

}

func newPostgreSQL(config *config.AppConfig) *gorm.DB {
	db_url := config.Database.Db_url
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func newMySQL(config *config.AppConfig) *gorm.DB {
	db_url := config.Database.Db_url
	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.MySQL != nil {
		db, _ := db.MySQL.DB()
		db.Close()
	}
}

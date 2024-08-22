package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
}

func InitDB(config DatabaseConfig) *gorm.DB {
	dsn := "host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.Username +
		" dbname=" + config.DBName +
		" password=" + config.Password +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	return db
}

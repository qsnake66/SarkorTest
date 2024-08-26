package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
}

func InitDB(config DatabaseConfig) *sql.DB {
	dsn := "host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.Username +
		" dbname=" + config.DBName +
		" password=" + config.Password +
		" sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %s", err.Error())
	}

	return db
}

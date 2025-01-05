package database

import (
	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"shared/logger"
	"time"
)

var Connection *sql.DB

// Init - create connection to the database server
func Init() {
	var err error
	dbName := os.Getenv("DATABASE_NAME")
	Connection, err = sql.Open("sqlite3", dbName)
	if err != nil {
		logger.Logger.Fatalf("failed to connect to sqlite3 - %s", dbName)
	}
	Connection.SetMaxOpenConns(10)                  // Maximum 10 open connections
	Connection.SetMaxIdleConns(5)                   // Maximum 5 idle connections
	Connection.SetConnMaxLifetime(60 * time.Second) // Connections reused for 60 seconds
	err = Connection.Ping()
	if err != nil {
		logger.Logger.Fatalf("failed to ping sqlite3 - %s", dbName)
	}
	logger.Logger.Infof("connection established - %s", dbName)
}

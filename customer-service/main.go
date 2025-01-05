package main

import (
	"customer-service/database"
	"customer-service/server"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	database.Init(os.Getenv("DATABASE_NAME"))
	server.Init(os.Getenv("PORT"))
}

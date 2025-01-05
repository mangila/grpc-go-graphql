package main

import (
	_ "github.com/joho/godotenv/autoload"
	"order-service/database"
	"order-service/server"
	"os"
)

func main() {
	database.Init(os.Getenv("DATABASE_NAME"))
	server.Init(os.Getenv("PORT"))
}

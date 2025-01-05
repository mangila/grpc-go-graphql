package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"product-service/database"
	"product-service/server"
)

func main() {
	database.Init(os.Getenv("DATABASE_NAME"))
	server.Init(os.Getenv("PORT"))
}

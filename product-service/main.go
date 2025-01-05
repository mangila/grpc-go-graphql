package main

import (
	"product-service/database"
	"product-service/server"
)

func main() {
	database.Init()
	server.Init()
}

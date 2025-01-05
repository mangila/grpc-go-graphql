package main

import (
	"order-service/database"
	"order-service/server"
)

func main() {
	database.Init()
	server.Init()
}

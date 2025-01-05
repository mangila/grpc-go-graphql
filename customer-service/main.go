package main

import (
	"customer-service/database"
	"customer-service/server"
)

func main() {
	database.Init()
	server.Init()
}

package main

import (
	"fmt"
	"shared/logger"
	"ui-server/internal/server"
)

func main() {
	srv := server.NewServer()
	fmt.Printf("ui-server - %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Logger.Fatalf("cannot start server: %s", err)
	}
}

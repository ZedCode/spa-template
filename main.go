package main

import (
	"os"
	"spaserver"
)

func main() {
	port := os.Getenv("UI_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	server := spaserver.NewServer()
	server.Run(":" + port)
}

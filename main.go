package main

import (
	"log"
	"user-service/config"
	"user-service/db"
	"user-service/server"
)

func main() {
	// load configuration
	if err := config.Load("."); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// connect to database
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	server.Start()
}

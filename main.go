package main

import (
	"log"
	"user-service/config"
	"user-service/db"
	"user-service/server"
)

func main() {

	// Load configuration settings.
	if err := config.Load("."); err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	server.Start()
}

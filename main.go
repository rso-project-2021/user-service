package main

import (
	"log"
	"user-service/config"
	"user-service/db"
	"user-service/server"
)

func main() {

	// Load configuration settings.
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	if err := db.Connect(config.DBDriver, config.DBSource); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	server.Start(config.ServerAddress, config.GinMode)
}

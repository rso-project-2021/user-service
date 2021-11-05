package models

import (
	"log"
	"os"
	"testing"
	"user-service/config"
	"user-service/db"
)

func TestMain(m *testing.M) {

	// Load configuration settings.
	if err := config.Load("../."); err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Run tests.
	code := m.Run()

	os.Exit(code)
}

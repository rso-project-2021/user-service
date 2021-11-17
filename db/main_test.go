package db

import (
	"log"
	"os"
	"testing"
	"user-service/config"
)

var testStore *Store

func TestMain(m *testing.M) {

	// Load configuration settings.
	config, err := config.LoadConfig("./..")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	testStore, err = Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Run tests.
	code := m.Run()

	os.Exit(code)
}

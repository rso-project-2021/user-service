package db

import (
	"log"
	"user-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect() (err error) {

	// connect to database
	config := config.Read()
	db, err = sqlx.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		return
	}

	// test database connection
	if err = db.Ping(); err != nil {
		return
	}

	log.Println("Connected to database!")
	return
}

func GetDB() *sqlx.DB {
	return db
}

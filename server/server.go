package server

import (
	"user-service/config"
)

func Start() {
	config := config.Read()
	router := NewRouter()
	router.Run(config.ServerAddress)
}

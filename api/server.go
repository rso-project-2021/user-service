package api

import (
	"user-service/config"
	"user-service/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	router := gin.Default()
	gin.SetMode(config.GinMode)

	// Setup routing for server.
	v1 := router.Group("v1")
	{
		v1.GET("/users/:id", server.GetUserByID)
		v1.GET("/users", server.GetAllUsers)
		v1.POST("/users", server.CreateUser)
		v1.PUT("/users/:id", server.UpdateUser)
		v1.DELETE("/users/:id", server.DeleteUser)
	}

	/*
		health := router.Group("health")
		{
			health.GET("/live", _______)
			health.GET("/ready", _______)
		}
	*/

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

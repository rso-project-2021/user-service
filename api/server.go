package api

import (
	"user-service/config"
	"user-service/db"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	config config.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {

	gin.SetMode(config.GinMode)
	router := gin.Default()

	server := &Server{
		config: config,
		store:  store,
	}

	// Setup routing for server.
	v1 := router.Group("v1")
	{
		v1.GET("/users/:id", server.GetUserByID)
		v1.GET("/users", server.GetAllUsers)
		v1.POST("/users", server.CreateUser)
		v1.PUT("/users/:id", server.UpdateUser)
		v1.DELETE("/users/:id", server.DeleteUser)
	}

	// Setup health check routes.
	health := router.Group("health")
	{
		health.GET("/live", server.Live)
		health.GET("/ready", server.Ready)
	}

	// Setup metrics routes.
	metrics := router.Group("metrics")
	{
		metrics.GET("/", func(ctx *gin.Context) {
			handler := promhttp.Handler()
			handler.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

package api

import (
	cnfg "user-service/config"
	"user-service/db"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	config cnfg.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config cnfg.Config, store *db.Store) (*Server, error) {

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
		v1.POST("/users/login", server.DefinitelyNotLogin)
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

	// dynamic configuration with consul
	go cnfg.KeyWatcher("db_source", func(source string) {
		store, err := db.Connect(server.config.DBDriver, source)
		if err == nil {
			server.store = store
		}
	})

	return server.router.Run(address)
}

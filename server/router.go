package server

import (
	"user-service/config"
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	config := config.Read()
	gin.SetMode(config.GinMode)

	// Expose router paths.
	v1 := router.Group("v1")
	{
		user := new(controllers.UserController)
		v1.GET("/users/:id", user.GetByID)
		v1.GET("/users", user.GetAll)
		v1.POST("/users", user.Create)
		v1.PUT("/users/:id", user.Update)
		v1.DELETE("/users/:id", user.Delete)
	}

	return router
}

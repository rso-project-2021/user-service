package server

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(ginMode string) *gin.Engine {

	router := gin.Default()
	gin.SetMode(ginMode)

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

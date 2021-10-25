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

	// set routing paths
	v1 := router.Group("v1")
	{
		account := new(controllers.AccountController)
		v1.GET("/accounts/:id", account.GetByID)
	}

	return router
}

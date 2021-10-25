package controllers

import (
	"net/http"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

var accountModel = new(models.Account)

func (ac AccountController) GetByID(c *gin.Context) {

	if c.Param("id") != "" {
		account, err := accountModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve account", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Account founded!", "account": account})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

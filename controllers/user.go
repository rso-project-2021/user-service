package controllers

import (
	"net/http"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type getRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

var user = new(models.User)

func (uc UserController) GetByID(ctx *gin.Context) {
	var req getRequest

	// check request correctness
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// retrieve results
	result, err := user.GetByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (uc UserController) GetAll(ctx *gin.Context) {

	// retrieve results
	result, err := user.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

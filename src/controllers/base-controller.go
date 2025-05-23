package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController interface {
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
}

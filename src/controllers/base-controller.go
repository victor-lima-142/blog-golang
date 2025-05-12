package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController interface {
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Create(c *gin.Context)
}

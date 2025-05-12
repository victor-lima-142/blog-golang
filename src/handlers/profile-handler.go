package handlers

import (
	"blog/core/repositories"
	"blog/core/services"
	"blog/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileHandler(router *gin.RouterGroup, db *gorm.DB) {
	profileRepository := repositories.NewProfileRepository(db)

	profileService := services.NewProfileService(profileRepository)

	profileController := controllers.NewProfileController(profileService)

	profileRouter := router.Group("/profile")
	profileRouter.GET("/me/:id", profileController.GetMe)
	profileRouter.GET("/:id", profileController.FindByID)
	profileRouter.GET("/", profileController.FindAll)
	profileRouter.DELETE("/:id", profileController.Delete)
	profileRouter.PUT("/", profileController.Update)
}

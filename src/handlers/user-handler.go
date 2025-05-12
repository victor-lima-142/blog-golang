package handlers

import (
	"blog/core/repositories"
	"blog/core/services"
	"blog/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserHandler(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	profileRepository := repositories.NewProfileRepository(db)
	followRepository := repositories.NewFollowRepository(db)

	userService := services.NewUserService(userRepository)
	profileService := services.NewProfileService(profileRepository)
	followService := services.NewFollowService(followRepository)

	userController := controllers.NewUserController(userService, profileService, followService)

	userRouter := router.Group("/user")
	userRouter.GET("/me", userController.GetMe)
	userRouter.GET("/", userController.FindAll)
	userRouter.GET("/:id", userController.FindByID)
	userRouter.DELETE("/:id", userController.Delete)
	userRouter.POST("/", userController.Create)
	userRouter.PUT("/", userController.Update)
	userRouter.POST("/follow", userController.FollowUser)
	userRouter.POST("/unfollow", userController.UnfollowUser)
}

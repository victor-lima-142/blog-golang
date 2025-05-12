package handlers

import (
	"blog/core/repositories"
	"blog/core/services"
	"blog/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleHandler(router *gin.RouterGroup, db *gorm.DB) {
	articleRepository := repositories.NewArticleRepository(db)

	articleService := services.NewArticleService(articleRepository)

	articleController := controllers.NewArticleController(articleService)

	articleRouter := router.Group("/article")
	articleRouter.GET("/", articleController.FindAll)
	articleRouter.GET("/:id", articleController.FindByID)
	articleRouter.DELETE("/:id", articleController.Delete)
	articleRouter.POST("/", articleController.Create)
	articleRouter.PUT("/", articleController.Update)
}

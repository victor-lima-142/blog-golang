package controllers

import (
	"blog/core/models"
	"blog/core/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	BaseController
}

type articleController struct {
	articleService services.ArticleService
}

func NewArticleController(articleService services.ArticleService) ArticleController {
	return &articleController{articleService}
}

func (ctrl *articleController) FindAll(ctx *gin.Context) {
	var articles []models.Article

	articles, err := ctrl.articleService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, articles)
}

func (ctrl *articleController) FindByID(ctx *gin.Context) {
	var ID int64
	var article *models.Article
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	article, err = ctrl.articleService.FindByID(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, article)
}

func (ctrl *articleController) Create(ctx *gin.Context) {
	var article models.Article

	err := ctx.ShouldBindJSON(&article)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	err = ctrl.articleService.Create(&article)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *articleController) Update(ctx *gin.Context) {
	var ID int64
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	var article models.Article
	err = ctx.ShouldBindJSON(&article)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	article.ID = ID

	err = ctrl.articleService.Update(&article)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *articleController) Delete(ctx *gin.Context) {
	var ID int64
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	err = ctrl.articleService.Delete(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

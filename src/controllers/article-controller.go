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

func (ctrl *articleController) FindAll(c *gin.Context) {
	var articles []models.Article

	articles, err := ctrl.articleService.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, articles)
}

func (ctrl *articleController) FindByID(c *gin.Context) {
	var ID int64
	var article *models.Article
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	article, err = ctrl.articleService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, article)
}

func (ctrl *articleController) Create(c *gin.Context) {
	var article models.Article

	err := c.ShouldBindJSON(&article)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	err = ctrl.articleService.Create(&article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *articleController) Update(c *gin.Context) {
	var ID int64
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	var article models.Article
	err = c.ShouldBindJSON(&article)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	article.ID = ID

	err = ctrl.articleService.Update(&article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *articleController) Delete(c *gin.Context) {
	var ID int64
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	err = ctrl.articleService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

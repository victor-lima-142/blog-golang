package controllers

import (
	"blog/core/models"
	"blog/core/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	BaseController
	GetMe(c *gin.Context)
}

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) ProfileController {
	return &profileController{profileService: profileService}
}

func (ctrl *profileController) GetMe(c *gin.Context) {
	ctrl.FindByID(c)
}

func (ctrl *profileController) FindByID(c *gin.Context) {
	var ID int64
	var err error
	var profile *models.Profile

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	profile, err = ctrl.profileService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, profile)
}

func (ctrl *profileController) FindAll(c *gin.Context) {
	profiles, err := ctrl.profileService.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, profiles)
}

func (ctrl *profileController) Delete(c *gin.Context) {
	var ID int64
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	err = ctrl.profileService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *profileController) Update(c *gin.Context) {
	var ID int64
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	var profile models.Profile
	err = c.ShouldBindJSON(&profile)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	profile.ID = ID

	err = ctrl.profileService.Update(&profile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *profileController) Create(c *gin.Context) {
	var profile models.Profile

	err := c.ShouldBindJSON(&profile)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	err = ctrl.profileService.Create(&profile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

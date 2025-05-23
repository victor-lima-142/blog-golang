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
	GetMe(ctx *gin.Context)
}

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) ProfileController {
	return &profileController{profileService: profileService}
}

func (ctrl *profileController) GetMe(ctx *gin.Context) {
	ctrl.FindByID(ctx)
}

func (ctrl *profileController) FindByID(ctx *gin.Context) {
	var ID int64
	var err error
	var profile *models.Profile

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	profile, err = ctrl.profileService.FindByID(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, profile)
}

func (ctrl *profileController) FindAll(ctx *gin.Context) {
	profiles, err := ctrl.profileService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, profiles)
}

func (ctrl *profileController) Delete(ctx *gin.Context) {
	var ID int64
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	err = ctrl.profileService.Delete(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *profileController) Update(ctx *gin.Context) {
	var ID int64
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	var profile models.Profile
	err = ctx.ShouldBindJSON(&profile)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	profile.ID = ID

	err = ctrl.profileService.Update(&profile)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *profileController) Create(ctx *gin.Context) {
	var profile models.Profile

	err := ctx.ShouldBindJSON(&profile)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	err = ctrl.profileService.Create(&profile)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

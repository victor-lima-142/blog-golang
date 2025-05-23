package controllers

import (
	"blog/core/models"
	"blog/core/services"
	"blog/src/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	BaseController
	GetMe(ctx *gin.Context)
	FollowUser(ctx *gin.Context)
	UnfollowUser(ctx *gin.Context)
}

type userController struct {
	UserController
	userService    services.UserService
	profileService services.ProfileService
	followSerivce  services.FollowService
}

func NewUserController(userService services.UserService, profileService services.ProfileService, followService services.FollowService) UserController {
	return &userController{userService: userService, profileService: profileService, followSerivce: followService}
}

func (ctrl *userController) GetMe(ctx *gin.Context) {
	ctrl.FindByID(ctx)
}

func (ctrl *userController) FindByID(ctx *gin.Context) {
	var ID int64
	var user *models.User
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	user, err = ctrl.userService.FindByID(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)
}

func (ctrl *userController) FindAll(ctx *gin.Context) {
	var users []models.User
	var err error

	users, err = ctrl.userService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, users)
}

func (ctrl *userController) Delete(ctx *gin.Context) {
	var ID int64
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	err = ctrl.userService.Delete(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *userController) FollowUser(ctx *gin.Context) {
	var payload schemas.FollowAndUnfollowUserSchema

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid request": err.Error()})
	}

	err = ctrl.followSerivce.FollowUser(payload.FollowerID, payload.FollowingID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *userController) UnfollowUser(ctx *gin.Context) {
	var payload schemas.FollowAndUnfollowUserSchema

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid request": err.Error()})
	}

	err = ctrl.followSerivce.UnfollowUser(payload.FollowerID, payload.FollowingID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *userController) Update(ctx *gin.Context) {
	var ID int64
	var err error

	routeID := ctx.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	var user models.User
	err = ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	user.ID = ID

	err = ctrl.userService.Update(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *userController) Create(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	err = ctrl.userService.Create(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	ctx.JSON(http.StatusOK, nil)
}

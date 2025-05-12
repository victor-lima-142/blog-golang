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
	GetMe(c *gin.Context)
	FollowUser(c *gin.Context)
	UnfollowUser(c *gin.Context)
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

func (ctrl *userController) GetMe(c *gin.Context) {
	ctrl.FindByID(c)
}

func (ctrl *userController) FindByID(c *gin.Context) {
	var ID int64
	var user *models.User
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	user, err = ctrl.userService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *userController) FindAll(c *gin.Context) {
	var users []models.User
	var err error

	users, err = ctrl.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *userController) Delete(c *gin.Context) {
	var ID int64
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	err = ctrl.userService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *userController) FollowUser(c *gin.Context) {
	var payload schemas.FollowAndUnfollowUserSchema

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request": err.Error()})
	}

	err = ctrl.followSerivce.FollowUser(payload.FollowerID, payload.FollowingID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *userController) UnfollowUser(c *gin.Context) {
	var payload schemas.FollowAndUnfollowUserSchema

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request": err.Error()})
	}

	err = ctrl.followSerivce.UnfollowUser(payload.FollowerID, payload.FollowingID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *userController) Update(c *gin.Context) {
	var ID int64
	var err error

	routeID := c.Param("id")

	if routeID != "" {
		ID, err = strconv.ParseInt(routeID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
		}
	}

	var user models.User
	err = c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	user.ID = ID

	err = ctrl.userService.Update(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (ctrl *userController) Create(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	err = ctrl.userService.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

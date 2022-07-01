package controllers

import (
	"net/http"

	"github.com/cmani097/go-mongo/models"
	"github.com/cmani097/go-mongo/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (usercontroller *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = usercontroller.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (usercontroller *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := usercontroller.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (usercontroller *UserController) GetAll(ctx *gin.Context) {
	users, err := usercontroller.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (usercontroller *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = usercontroller.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (usercontroller *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := usercontroller.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (usercontroller *UserController) RegisterUserRotes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", usercontroller.CreateUser)
	userRoute.GET("/get/:name", usercontroller.GetUser)
	userRoute.GET("/getall", usercontroller.GetAll)
	userRoute.PATCH("/update", usercontroller.UpdateUser)
	userRoute.DELETE("/delete/:name", usercontroller.DeleteUser)
}

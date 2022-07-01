package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slice-mani/gin-postgres-project/models"
	"github.com/slice-mani/gin-postgres-project/services"
)

type SampleController struct {
	SampleService services.SampleService
}

func New(sampleservice services.SampleService) SampleController {
	return SampleController{
		SampleService: sampleservice,
	}
}

func (samplecontroller *SampleController) TestService(ctx *gin.Context) {
	fmt.Println("CreateUser of sample controller called")
	var user models.User
	err := samplecontroller.SampleService.TestService(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

package controller

import (
	"strconv"

	"github.com/Buddy-Git/golang-git-pro/model/entity"
	"github.com/Buddy-Git/golang-git-pro/service"
	"github.com/Buddy-Git/golang-git-pro/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func New(vservice service.VideoService) VideoController {
	validate := validator.New()

	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: vservice,
	}
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *videoController) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}

func (c *videoController) Delete(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	c.service.Delete(video)
	return nil
}

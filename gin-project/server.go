package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Buddy-Git/golang-git-pro/controller"
	"github.com/Buddy-Git/golang-git-pro/middlewares"
	"github.com/Buddy-Git/golang-git-pro/repository"
	"github.com/Buddy-Git/golang-git-pro/service"
	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	defer videoRepository.CloseDB()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "video input is valid"})
	})

	server.PUT("/videos", func(ctx *gin.Context) {
		err := videoController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "video input is valid"})
	})

	server.DELETE("/videos/:id", func(ctx *gin.Context) {
		err := videoController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "video input is valid"})
	})

	server.Run(":8000")
}

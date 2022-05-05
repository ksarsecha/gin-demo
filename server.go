package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ksarsecha/gin-demo/controller"
	"github.com/ksarsecha/gin-demo/service"
	"net/http"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.Save(context))
	})

	server.Run(":8080")
}

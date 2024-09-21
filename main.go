package main

import (
	"github.com/gin-gonic/gin"
	"github.come/salmanfaris22/Skill-map/controller"
	"github.come/salmanfaris22/Skill-map/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService) // Fixed naming
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		videoController.Save(ctx) // The Save method handles the response
	})

	server.Run()
}

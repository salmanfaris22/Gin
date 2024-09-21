package main

import (
	"github.com/gin-gonic/gin"
	"github.come/salmanfaris22/Skill-map/controller"
	"github.come/salmanfaris22/Skill-map/service"
)

var (
	videoService    service.VideoService       = service.New()
	vedioController controller.VedioController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, vedioController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, vedioController.Save(ctx))
	})
	server.Run()
}

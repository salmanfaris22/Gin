package controller

import (
	"github.com/gin-gonic/gin"
	"github.come/salmanfaris22/Skill-map/entity"
	"github.come/salmanfaris22/Skill-map/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	if err := ctx.BindJSON(&video); err != nil {
		ctx.JSON(400, gin.H{"error": "Bad Request"})
		return entity.Video{} // Return an empty video on error
	}
	savedVideo := c.service.Save(video)
	ctx.JSON(201, savedVideo) // Return the saved video with 201 status
	return savedVideo
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll() // Call the service's FindAll method
}

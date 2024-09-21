package controller

import (
	"github.com/gin-gonic/gin"
	"github.come/salmanfaris22/Skill-map/entity"
	"github.come/salmanfaris22/Skill-map/service"
)

type VedioController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VedioController {
	return &controller{ // Use pointer here
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var vdeo entity.Video
	ctx.BindJSON(&vdeo)
	c.service.Save(vdeo)
	return vdeo

}
func (controller *controller) FindAll() []entity.Video {
	return controller.FindAll()

}

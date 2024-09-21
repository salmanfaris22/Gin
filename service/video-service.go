package service

import (
	"github.come/salmanfaris22/Skill-map/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	video []entity.Video
}

func New() *videoService {
	return &videoService{}
}

func (serv *videoService) Save(video entity.Video) entity.Video {
	serv.video = append(serv.video, video)
	return video
}
func (serv *videoService) FindAll() []entity.Video {
	return serv.video
}

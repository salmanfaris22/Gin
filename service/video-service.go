package service

import "github.come/salmanfaris22/Skill-map/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() *videoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (serv *videoService) Save(video entity.Video) entity.Video {
	serv.videos = append(serv.videos, video)
	return video
}

func (serv *videoService) FindAll() []entity.Video {
	return serv.videos
}

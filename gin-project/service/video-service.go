package service

import (
	"github.com/Buddy-Git/golang-git-pro/model/entity"
	"github.com/Buddy-Git/golang-git-pro/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(entity.Video)
	Delete(entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (vservice *videoService) Save(video entity.Video) entity.Video {
	vservice.videoRepository.Save(video)
	return video
}

func (vservice *videoService) Update(video entity.Video) {
	vservice.videoRepository.Update(video)
}
func (vservice *videoService) Delete(video entity.Video) {
	vservice.videoRepository.Delete(video)
}

func (vservice *videoService) FindAll() []entity.Video {
	return vservice.videoRepository.FindAll()
}

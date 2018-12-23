package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	postVideoServiceImpl struct {
		postRepository contracts.PostRepository
	}
)

func NewPostVideoService(postRepository contracts.PostRepository) contracts.PostVideoService {
	return &postVideoServiceImpl{postRepository: postRepository}
}

func (s *postVideoServiceImpl) ChangePostVideo(postEntity *entities.PostEntity, postVideo *entities.FileEntity) common.Error {
	postEntity.SetVideo(postVideo)

	return s.postRepository.SavePost(postEntity)
}

func (s *postVideoServiceImpl) RemovePostVideo(postEntity *entities.PostEntity) common.Error {
	postEntity.RemoveVideo()

	return s.postRepository.SavePost(postEntity)
}

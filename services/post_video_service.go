package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	postVideoServiceImpl struct {
		postRepository contracts.PostRepository
	}
)

func NewPostVideoService(postRepository contracts.PostRepository) contracts.PostVideoService {
	return &postVideoServiceImpl{postRepository: postRepository}
}

func (s *postVideoServiceImpl) ChangePostVideo(
	postEntity *entities.PostEntity, postVideoEntity *entities.FileEntity,
) errors.Error {
	postEntity.SetVideo(postVideoEntity)

	return s.postRepository.SavePost(postEntity)
}

func (s *postVideoServiceImpl) RemovePostVideo(postEntity *entities.PostEntity) errors.Error {
	postEntity.RemoveVideo()

	return s.postRepository.SavePost(postEntity)
}

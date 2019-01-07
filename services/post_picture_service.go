package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	postPictureServiceImpl struct {
		postRepository contracts.PostRepository
	}
)

func NewPostPictureService(postRepository contracts.PostRepository) contracts.PostPictureService {
	return &postPictureServiceImpl{postRepository: postRepository}
}

func (s *postPictureServiceImpl) ChangePostPicture(postEntity *entities.PostEntity, postPicture *entities.FileEntity) errors.Error {
	postEntity.SetPicture(postPicture)

	return s.postRepository.SavePost(postEntity)
}

func (s *postPictureServiceImpl) RemovePostPicture(postEntity *entities.PostEntity) errors.Error {
	postEntity.RemovePicture()

	return s.postRepository.SavePost(postEntity)
}

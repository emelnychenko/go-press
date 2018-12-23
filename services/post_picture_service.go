package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	postPictureServiceImpl struct {
		postRepository contracts.PostRepository
	}
)

func NewPostPictureService(postRepository contracts.PostRepository) contracts.PostPictureService {
	return &postPictureServiceImpl{postRepository: postRepository}
}

func (s *postPictureServiceImpl) ChangePostPicture(postEntity *entities.PostEntity, postPicture *entities.FileEntity) common.Error {
	postEntity.SetPicture(postPicture)

	return s.postRepository.SavePost(postEntity)
}

func (s *postPictureServiceImpl) RemovePostPicture(postEntity *entities.PostEntity) common.Error {
	postEntity.RemovePicture()

	return s.postRepository.SavePost(postEntity)
}

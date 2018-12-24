package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	postAuthorServiceImpl struct {
		postRepository contracts.PostRepository
	}
)

func NewPostAuthorService(postRepository contracts.PostRepository) contracts.PostAuthorService {
	return &postAuthorServiceImpl{postRepository: postRepository}
}

func (s *postAuthorServiceImpl) ChangePostAuthor(
	postEntity *entities.PostEntity, postAuthorEntity *entities.UserEntity,
) common.Error {
	postEntity.SetAuthor(postAuthorEntity)

	return s.postRepository.SavePost(postEntity)
}

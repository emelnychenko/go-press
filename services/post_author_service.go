package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
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
) errors.Error {
	postEntity.SetAuthor(postAuthorEntity)

	return s.postRepository.SavePost(postEntity)
}

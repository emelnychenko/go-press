package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	CommentEntityFactory interface {
		CreateCommentEntity() *entities.CommentEntity
	}
)

package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	commentEntityFactoryImpl struct {
	}
)

func NewCommentEntityFactory() contracts.CommentEntityFactory {
	return new(commentEntityFactoryImpl)
}

func (*commentEntityFactoryImpl) CreateCommentEntity() *entities.CommentEntity {
	return entities.NewCommentEntity()
}

package apis

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	postAuthorApiImpl struct {
		eventDispatcher        contracts.EventDispatcher
		postAuthorEventFactory contracts.PostAuthorEventFactory
		postService            contracts.PostService
		userService            contracts.UserService
		postAuthorService      contracts.PostAuthorService
	}
)

func NewPostAuthorApi(
	eventDispatcher contracts.EventDispatcher,
	postAuthorEventFactory contracts.PostAuthorEventFactory,
	postService contracts.PostService,
	userService contracts.UserService,
	postAuthorService contracts.PostAuthorService,
) (postAuthorApi contracts.PostAuthorApi) {
	return &postAuthorApiImpl{
		eventDispatcher:        eventDispatcher,
		postAuthorEventFactory: postAuthorEventFactory,
		postService:            postService,
		userService:            userService,
		postAuthorService:      postAuthorService,
	}
}

func (a *postAuthorApiImpl) ChangePostAuthor(postId *models.PostId, postAuthorId *models.UserId) (err errors.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	postAuthorEntity, err := a.userService.GetUser(postAuthorId)

	if nil != err {
		return
	}

	err = a.postAuthorService.ChangePostAuthor(postEntity, postAuthorEntity)

	if nil != err {
		return
	}

	postAuthorEvent := a.postAuthorEventFactory.CreatePostAuthorChangedEvent(postEntity, postAuthorEntity)
	a.eventDispatcher.Dispatch(postAuthorEvent)

	return
}

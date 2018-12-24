package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	postPictureApiImpl struct {
		eventDispatcher         contracts.EventDispatcher
		postPictureEventFactory contracts.PostPictureEventFactory
		postService             contracts.PostService
		fileService             contracts.FileService
		postPictureService      contracts.PostPictureService
	}
)

func NewPostPictureApi(
	eventDispatcher contracts.EventDispatcher,
	postPictureEventFactory contracts.PostPictureEventFactory,
	postService contracts.PostService,
	fileService contracts.FileService,
	postPictureService contracts.PostPictureService,
) (postPictureApi contracts.PostPictureApi) {
	return &postPictureApiImpl{
		eventDispatcher:         eventDispatcher,
		postPictureEventFactory: postPictureEventFactory,
		postService:             postService,
		fileService:             fileService,
		postPictureService:      postPictureService,
	}
}

func (a *postPictureApiImpl) ChangePostPicture(postId *models.PostId, postPictureId *models.FileId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	postPictureEntity, err := a.fileService.GetFile(postPictureId)

	if nil != err {
		return
	}

	err = a.postPictureService.ChangePostPicture(postEntity, postPictureEntity)

	postPictureEvent := a.postPictureEventFactory.CreatePostPictureChangedEvent(postEntity, postPictureEntity)
	a.eventDispatcher.Dispatch(postPictureEvent)
	
	return
}

func (a *postPictureApiImpl) RemovePostPicture(postId *models.PostId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	err = a.postPictureService.RemovePostPicture(postEntity)

	postPictureEvent := a.postPictureEventFactory.CreatePostPictureRemovedEvent(postEntity)
	a.eventDispatcher.Dispatch(postPictureEvent)

	return
}

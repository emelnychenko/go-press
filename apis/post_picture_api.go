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
		contentTypeValidator    contracts.ContentTypeValidator
		postService             contracts.PostService
		fileService             contracts.FileService
		postPictureService      contracts.PostPictureService
	}
)

func NewPostPictureApi(
	eventDispatcher contracts.EventDispatcher,
	postPictureEventFactory contracts.PostPictureEventFactory,
	contentTypeValidator contracts.ContentTypeValidator,
	postService contracts.PostService,
	fileService contracts.FileService,
	postPictureService contracts.PostPictureService,
) (postPictureApi contracts.PostPictureApi) {
	return &postPictureApiImpl{
		eventDispatcher:         eventDispatcher,
		postPictureEventFactory: postPictureEventFactory,
		contentTypeValidator:    contentTypeValidator,
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

	err = a.contentTypeValidator.ValidateImage(postPictureEntity.Type)

	if nil != err {
		return
	}

	err = a.postPictureService.ChangePostPicture(postEntity, postPictureEntity)

	if nil != err {
		return
	}

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

	if nil != err {
		return
	}

	postPictureEvent := a.postPictureEventFactory.CreatePostPictureRemovedEvent(postEntity)
	a.eventDispatcher.Dispatch(postPictureEvent)

	return
}

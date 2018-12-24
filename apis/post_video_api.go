package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	postVideoApiImpl struct {
		eventDispatcher       contracts.EventDispatcher
		postVideoEventFactory contracts.PostVideoEventFactory
		postService           contracts.PostService
		fileService           contracts.FileService
		postVideoService      contracts.PostVideoService
	}
)

func NewPostVideoApi(
	eventDispatcher contracts.EventDispatcher,
	postVideoEventFactory contracts.PostVideoEventFactory,
	postService contracts.PostService,
	fileService contracts.FileService,
	postVideoService contracts.PostVideoService,
) (postVideoApi contracts.PostVideoApi) {
	return &postVideoApiImpl{
		eventDispatcher:       eventDispatcher,
		postVideoEventFactory: postVideoEventFactory,
		postService:           postService,
		fileService:           fileService,
		postVideoService:      postVideoService,
	}
}

func (a *postVideoApiImpl) ChangePostVideo(postId *models.PostId, postVideoId *models.FileId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	postVideoEntity, err := a.fileService.GetFile(postVideoId)

	if nil != err {
		return
	}

	err = a.postVideoService.ChangePostVideo(postEntity, postVideoEntity)

	postVideoEvent := a.postVideoEventFactory.CreatePostVideoChangedEvent(postEntity, postVideoEntity)
	a.eventDispatcher.Dispatch(postVideoEvent)

	return
}

func (a *postVideoApiImpl) RemovePostVideo(postId *models.PostId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	err = a.postVideoService.RemovePostVideo(postEntity)

	postVideoEvent := a.postVideoEventFactory.CreatePostVideoRemovedEvent(postEntity)
	a.eventDispatcher.Dispatch(postVideoEvent)

	return
}

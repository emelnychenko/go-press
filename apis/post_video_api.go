package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	postVideoApiImpl struct {
		postService        contracts.PostService
		fileService        contracts.FileService
		postVideoService contracts.PostVideoService
	}
)

func NewPostVideoApi(
	postService contracts.PostService,
	fileService contracts.FileService,
	postVideoService contracts.PostVideoService,
) (postVideoApi contracts.PostVideoApi) {
	return &postVideoApiImpl{
		postService:        postService,
		fileService:        fileService,
		postVideoService: postVideoService,
	}
}

func (a *postVideoApiImpl) ChangePostVideo(postId *models.PostId, postVideoId *models.FileId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	postVideo, err := a.fileService.GetFile(postVideoId)

	if nil != err {
		return
	}

	return a.postVideoService.ChangePostVideo(postEntity, postVideo)
}

func (a *postVideoApiImpl) RemovePostVideo(postId *models.PostId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	return a.postVideoService.RemovePostVideo(postEntity)
}

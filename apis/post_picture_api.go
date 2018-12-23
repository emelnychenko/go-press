package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	postPictureApiImpl struct {
		postService        contracts.PostService
		fileService        contracts.FileService
		postPictureService contracts.PostPictureService
	}
)

func NewPostPictureApi(
	postService contracts.PostService,
	fileService contracts.FileService,
	postPictureService contracts.PostPictureService,
) (postPictureApi contracts.PostPictureApi) {
	return &postPictureApiImpl{
		postService:        postService,
		fileService:        fileService,
		postPictureService: postPictureService,
	}
}

func (a *postPictureApiImpl) ChangePostPicture(postId *models.PostId, postPictureId *models.FileId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	postPicture, err := a.fileService.GetFile(postPictureId)

	if nil != err {
		return
	}

	return a.postPictureService.ChangePostPicture(postEntity, postPicture)
}

func (a *postPictureApiImpl) RemovePostPicture(postId *models.PostId) (err common.Error) {
	postEntity, err := a.postService.GetPost(postId)

	if nil != err {
		return
	}

	return a.postPictureService.RemovePostPicture(postEntity)
}

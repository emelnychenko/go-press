package providers

import (
	"bytes"
	"errors"
	"github.com/emelnychenko/go-press/aws_mocks"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAwsS3StorageProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewAwsS3StorageProvider", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().AccessKeyId().Return("")
		awsS3Parameters.EXPECT().SecretKeyId().Return("")
		awsS3Parameters.EXPECT().Region().Return("")

		awsS3WriterProxyFactory := mocks.NewMockAwsS3WriterProxyFactory(ctrl)

		awsS3UploaderFactory := mocks.NewMockAwsS3UploaderFactory(ctrl)
		uploaderAPI := aws_mocks.NewMockUploaderAPI(ctrl)
		awsS3UploaderFactory.EXPECT().Create(gomock.Any()).Return(uploaderAPI)

		awsS3DownloaderFactory := mocks.NewMockAwsS3DownloaderFactory(ctrl)
		downloaderAPI := aws_mocks.NewMockDownloaderAPI(ctrl)
		awsS3DownloaderFactory.EXPECT().Create(gomock.Any()).Return(downloaderAPI)

		var awsS3StorageProvider *awsS3StorageProviderImpl
		storageProvider := NewAwsS3StorageProvider(awsS3Parameters, awsS3WriterProxyFactory, awsS3UploaderFactory, awsS3DownloaderFactory)
		awsS3StorageProvider, isAwsS3StorageProvider := storageProvider.(*awsS3StorageProviderImpl)

		assert.True(t, isAwsS3StorageProvider)
		assert.Equal(t, awsS3WriterProxyFactory, awsS3StorageProvider.awsS3WriterProxyFactory)
		assert.Equal(t, awsS3Parameters, awsS3StorageProvider.awsS3Parameters)
		assert.Equal(t, uploaderAPI, awsS3StorageProvider.uploaderAPI)
		assert.Equal(t, downloaderAPI, awsS3StorageProvider.downloaderAPI)
	})

	t.Run("UploadFile", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		uploaderAPI := aws_mocks.NewMockUploaderAPI(ctrl)
		uploaderAPI.EXPECT().Upload(gomock.Any()).Return(nil, nil)

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters: awsS3Parameters,
			uploaderAPI:     uploaderAPI,
		}

		fileEntity := entities.NewFileEntity()
		assert.Nil(t, awsS3StorageProvider.UploadFile(fileEntity, nil))
	})

	t.Run("UploadFile:Error", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		uploaderAPI := aws_mocks.NewMockUploaderAPI(ctrl)
		uploaderAPI.EXPECT().Upload(gomock.Any()).Return(nil, errors.New(""))

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters: awsS3Parameters,
			uploaderAPI:     uploaderAPI,
		}

		fileEntity := entities.NewFileEntity()
		assert.Error(t, awsS3StorageProvider.UploadFile(fileEntity, nil))
	})

	t.Run("DownloadFile", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		fileDestination := bytes.NewBuffer(nil)
		awsS3WriterProxy := mocks.NewMockAwsS3WriterProxy(ctrl)

		awsS3WriterProxyFactory := mocks.NewMockAwsS3WriterProxyFactory(ctrl)
		awsS3WriterProxyFactory.EXPECT().Create(fileDestination).Return(awsS3WriterProxy)

		downloaderAPI := aws_mocks.NewMockDownloaderAPI(ctrl)
		downloaderAPI.EXPECT().Download(awsS3WriterProxy, gomock.Any()).Return(int64(0), nil)

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters:         awsS3Parameters,
			awsS3WriterProxyFactory: awsS3WriterProxyFactory,
			downloaderAPI:           downloaderAPI,
		}

		fileEntity := entities.NewFileEntity()
		err := awsS3StorageProvider.DownloadFile(fileEntity, fileDestination)

		assert.Nil(t, err)
	})

	t.Run("DownloadFile:Error", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		fileDestination := bytes.NewBuffer(nil)
		awsS3WriterProxy := mocks.NewMockAwsS3WriterProxy(ctrl)

		awsS3WriterProxyFactory := mocks.NewMockAwsS3WriterProxyFactory(ctrl)
		awsS3WriterProxyFactory.EXPECT().Create(fileDestination).Return(awsS3WriterProxy)

		downloaderAPI := aws_mocks.NewMockDownloaderAPI(ctrl)
		downloaderAPI.EXPECT().Download(awsS3WriterProxy, gomock.Any()).Return(int64(0), errors.New(""))

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters:         awsS3Parameters,
			awsS3WriterProxyFactory: awsS3WriterProxyFactory,
			downloaderAPI:           downloaderAPI,
		}

		fileEntity := entities.NewFileEntity()
		err := awsS3StorageProvider.DownloadFile(fileEntity, fileDestination)

		assert.Error(t, err)
	})
}

package providers

import (
	"bytes"
	"errors"
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

		awsS3Factory := mocks.NewMockAwsS3Factory(ctrl)
		awsSdkS3 :=  mocks.NewMockS3API(ctrl)
		awsS3Factory.EXPECT().Create(gomock.Any()).Return(awsSdkS3)

		awsS3UploaderFactory := mocks.NewMockAwsS3UploaderFactory(ctrl)
		awsSdkS3Uploader :=  mocks.NewMockUploaderAPI(ctrl)
		awsS3UploaderFactory.EXPECT().Create(gomock.Any()).Return(awsSdkS3Uploader)

		awsS3DownloaderFactory := mocks.NewMockAwsS3DownloaderFactory(ctrl)
		awsSdkS3Downloader :=  mocks.NewMockDownloaderAPI(ctrl)
		awsS3DownloaderFactory.EXPECT().Create(gomock.Any()).Return(awsSdkS3Downloader)

		var awsS3StorageProvider *awsS3StorageProviderImpl
		awsS3StorageProvider, isAwsS3StorageProvider := NewAwsS3StorageProvider(
			awsS3Parameters,
			awsS3WriterProxyFactory,
			awsS3Factory,
			awsS3UploaderFactory,
			awsS3DownloaderFactory,
		).(*awsS3StorageProviderImpl)

		assert.True(t, isAwsS3StorageProvider)
		assert.Equal(t, awsS3WriterProxyFactory, awsS3StorageProvider.awsS3WriterProxyFactory)
		assert.Equal(t, awsS3Parameters, awsS3StorageProvider.awsS3Parameters)
		assert.Equal(t, awsSdkS3, awsS3StorageProvider.awsSdkS3)
		assert.Equal(t, awsSdkS3Uploader, awsS3StorageProvider.awsSdkS3Uploader)
		assert.Equal(t, awsSdkS3Downloader, awsS3StorageProvider.awsSdkS3Downloader)
	})

	t.Run("UploadFile", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		awsSdkS3Uploader :=  mocks.NewMockUploaderAPI(ctrl)
		awsSdkS3Uploader.EXPECT().Upload(gomock.Any()).Return(nil, nil)

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters:  awsS3Parameters,
			awsSdkS3Uploader: awsSdkS3Uploader,
		}

		fileEntity := new(entities.FileEntity)
		assert.Nil(t, awsS3StorageProvider.UploadFile(fileEntity, nil))
	})

	t.Run("UploadFile:Error", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		awsSdkS3Uploader :=  mocks.NewMockUploaderAPI(ctrl)
		awsSdkS3Uploader.EXPECT().Upload(gomock.Any()).Return(nil, errors.New(""))

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters:  awsS3Parameters,
			awsSdkS3Uploader: awsSdkS3Uploader,
		}

		fileEntity := new(entities.FileEntity)
		assert.Error(t, awsS3StorageProvider.UploadFile(fileEntity, nil))
	})

	t.Run("DownloadFile", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		fileDestination := bytes.NewBuffer(nil)
		awsS3WriterProxy := mocks.NewMockAwsS3WriterProxy(ctrl)

		awsS3WriterProxyFactory := mocks.NewMockAwsS3WriterProxyFactory(ctrl)
		awsS3WriterProxyFactory.EXPECT().Create(fileDestination).Return(awsS3WriterProxy)

		awsSdkS3Downloader :=  mocks.NewMockDownloaderAPI(ctrl)
		awsSdkS3Downloader.EXPECT().Download(awsS3WriterProxy, gomock.Any()).Return(int64(0), nil)

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters:         awsS3Parameters,
			awsS3WriterProxyFactory: awsS3WriterProxyFactory,
			awsSdkS3Downloader:      awsSdkS3Downloader,
		}

		fileEntity := new(entities.FileEntity)
		assert.Nil(t, awsS3StorageProvider.DownloadFile(fileEntity, fileDestination))
	})

	t.Run("DownloadFile:Error", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		fileDestination := bytes.NewBuffer(nil)
		awsS3WriterProxy := mocks.NewMockAwsS3WriterProxy(ctrl)

		awsS3WriterProxyFactory := mocks.NewMockAwsS3WriterProxyFactory(ctrl)
		awsS3WriterProxyFactory.EXPECT().Create(fileDestination).Return(awsS3WriterProxy)

		awsSdkS3Downloader :=  mocks.NewMockDownloaderAPI(ctrl)
		awsSdkS3Downloader.EXPECT().Download(awsS3WriterProxy, gomock.Any()).Return(int64(0), errors.New(""))

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters:         awsS3Parameters,
			awsS3WriterProxyFactory: awsS3WriterProxyFactory,
			awsSdkS3Downloader:      awsSdkS3Downloader,
		}

		fileEntity := new(entities.FileEntity)
		assert.Error(t, awsS3StorageProvider.DownloadFile(fileEntity, fileDestination))
	})

	t.Run("DeleteFile", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		awsSdkS3 :=  mocks.NewMockS3API(ctrl)
		awsSdkS3.EXPECT().DeleteObject(gomock.Any()).Return(nil, nil)

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters: awsS3Parameters,
			awsSdkS3:        awsSdkS3,
		}

		fileEntity := new(entities.FileEntity)
		assert.Nil(t, awsS3StorageProvider.DeleteFile(fileEntity))
	})

	t.Run("DeleteFile:Error", func(t *testing.T) {
		awsS3Parameters := mocks.NewMockAwsS3Parameters(ctrl)
		awsS3Parameters.EXPECT().Bucket().Return("")

		awsSdkS3 :=  mocks.NewMockS3API(ctrl)
		awsSdkS3.EXPECT().DeleteObject(gomock.Any()).Return(nil, errors.New(""))

		awsS3StorageProvider := &awsS3StorageProviderImpl{
			awsS3Parameters: awsS3Parameters,
			awsSdkS3:        awsSdkS3,
		}

		fileEntity := new(entities.FileEntity)
		assert.Error(t, awsS3StorageProvider.DeleteFile(fileEntity))
	})
}

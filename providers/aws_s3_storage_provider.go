package providers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"io"
)

type (
	awsS3StorageProviderImpl struct {
		awsS3Parameters         contracts.AwsS3Parameters
		awsS3WriterProxyFactory contracts.AwsS3WriterProxyFactory
		awsSdkS3                s3iface.S3API
		awsSdkS3Uploader        s3manageriface.UploaderAPI
		awsSdkS3Downloader      s3manageriface.DownloaderAPI
	}
)

func NewAwsS3StorageProvider(
	parameters contracts.AwsS3Parameters,
	writerProxyFactory contracts.AwsS3WriterProxyFactory,
	awsS3Factory contracts.AwsS3Factory,
	uploaderFactory contracts.AwsS3UploaderFactory,
	downloaderFactory contracts.AwsS3DownloaderFactory,
) contracts.StorageProvider {
	// TODO: Add credentials check.
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(parameters.Region()),
		Credentials: credentials.NewStaticCredentials(
			parameters.AccessKeyId(),
			parameters.SecretKeyId(),
			"")},
	)

	awsSdkS3 := awsS3Factory.Create(sess)
	awsSdkS3Uploader := uploaderFactory.Create(sess)
	awsSdkS3Downloader := downloaderFactory.Create(sess)

	return &awsS3StorageProviderImpl{
		parameters,
		writerProxyFactory,
		awsSdkS3,
		awsSdkS3Uploader,
		awsSdkS3Downloader,
	}
}

func (a *awsS3StorageProviderImpl) UploadFile(fileEntity *entities.FileEntity, fileSource io.Reader) (err common.Error) {
	awsS3Bucket := a.awsS3Parameters.Bucket()

	_, err2 := a.awsSdkS3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awsS3Bucket),
		Key:    aws.String(fileEntity.Path),
		Body:   fileSource,
	})

	if err2 != nil {
		err = common.ServerError(fmt.Sprintf("Unable to upload %q to %q, %v", fileEntity.Path, awsS3Bucket, err2))
	}
	return
}

func (a *awsS3StorageProviderImpl) DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) (err common.Error) {
	awsS3WriterProxy := a.awsS3WriterProxyFactory.Create(fileDestination)

	_, err2 := a.awsSdkS3Downloader.Download(awsS3WriterProxy,
		&s3.GetObjectInput{
			Bucket: aws.String(a.awsS3Parameters.Bucket()),
			Key:    aws.String(fileEntity.Path),
		})

	if err2 != nil {
		err = common.ServerError(fmt.Sprintf("Unable to download item %q, %v", fileEntity.Path, err2))
	}

	return
}

func (a *awsS3StorageProviderImpl) DeleteFile(fileEntity *entities.FileEntity) (err common.Error) {
	awsS3Bucket := a.awsS3Parameters.Bucket()

	_, err2 := a.awsSdkS3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(awsS3Bucket),
		Key:    aws.String(fileEntity.Path),
	})

	if err2 != nil {
		err = common.ServerError(fmt.Sprintf("Unable to delete object %q from awsS3Bucket %q, %v", fileEntity.Path, awsS3Bucket, err2))
	}

	return
}

package providers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
		uploaderAPI             s3manageriface.UploaderAPI
		downloaderAPI           s3manageriface.DownloaderAPI
	}
)

func NewAwsS3StorageProvider(
	parameters contracts.AwsS3Parameters,
	writerProxyFactory contracts.AwsS3WriterProxyFactory,
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

	awsS3Uploader := uploaderFactory.Create(sess)
	awsS3Downloader := downloaderFactory.Create(sess)

	return &awsS3StorageProviderImpl{
		parameters,
		writerProxyFactory,
		awsS3Uploader,
		awsS3Downloader,
	}
}

func (a *awsS3StorageProviderImpl) UploadFile(fileEntity *entities.FileEntity, fileSource io.Reader) (err common.Error) {
	bucket := a.awsS3Parameters.Bucket()

	_, err2 := a.uploaderAPI.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileEntity.Path),
		Body:   fileSource,
	})
	if err2 != nil {
		// Print the error and exit.
		err = common.ServerError(fmt.Sprintf("Unable to upload %q to %q, %v", fileEntity.Path, bucket, err2))
	}
	return
}

func (a *awsS3StorageProviderImpl) DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) (err common.Error) {
	awsS3WriterProxy := a.awsS3WriterProxyFactory.Create(fileDestination)

	_, err2 := a.downloaderAPI.Download(awsS3WriterProxy,
		&s3.GetObjectInput{
			Bucket: aws.String(a.awsS3Parameters.Bucket()),
			Key:    aws.String(fileEntity.Path),
		})
	if err2 != nil {
		err = common.ServerError(fmt.Sprintf("Unable to download item %q, %v", fileEntity.Path, err2))
	}

	return
}

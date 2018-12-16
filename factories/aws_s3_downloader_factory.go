package factories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/emelnychenko/go-press/contracts"
)

type (
	awsS3DownloaderFactoryImpl struct{}
)

func NewAwsS3DownloaderFactory() contracts.AwsS3DownloaderFactory {
	return &awsS3DownloaderFactoryImpl{}
}

func (*awsS3DownloaderFactoryImpl) Create(sess *session.Session) s3manageriface.DownloaderAPI {
	return s3manager.NewDownloader(sess)
}

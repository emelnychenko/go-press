package factories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/emelnychenko/go-press/contracts"
)

type (
	awsS3UploaderFactoryImpl struct {
	}
)

func NewAwsS3UploaderFactory() contracts.AwsS3UploaderFactory {
	return &awsS3UploaderFactoryImpl{}
}

func (*awsS3UploaderFactoryImpl) Create(sess *session.Session) (awsSdkS3Uploader s3manageriface.UploaderAPI) {
	return s3manager.NewUploader(sess)
}

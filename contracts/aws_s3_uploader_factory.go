package contracts

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
)

type (
	AwsS3UploaderFactory interface {
		Create(sess *session.Session) (awsSdkS3Uploader s3manageriface.UploaderAPI)
	}
)

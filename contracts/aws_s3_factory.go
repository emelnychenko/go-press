package contracts

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type (
	AwsS3Factory interface {
		Create(sess *session.Session) (awsSdkS3 s3iface.S3API)
	}
)

package factories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/emelnychenko/go-press/contracts"
)

type (
	awsS3FactoryImpl struct {
	}
)

func NewAwsS3Factory() contracts.AwsS3Factory {
	return &awsS3FactoryImpl{}
}

func (*awsS3FactoryImpl) Create(sess *session.Session) (awsSdkS3 s3iface.S3API) {
	return s3.New(sess)
}

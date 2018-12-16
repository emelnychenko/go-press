package parameters

import (
	"github.com/emelnychenko/go-press/contracts"
	"os"
)

const (
	AwsS3AccessKeyIdParameter     string = "AWS_S3_ACCESS_KEY_ID"
	AwsS3SecretAccessKeyParameter string = "AWS_S3_SECRET_ACCESS_KEY"
	AwsS3BucketParameter          string = "AWS_S3_BUCKET"
	AwsS3RegionParameter          string = "AWS_S3_REGION"
)

type (
	awsS3ParametersImpl struct {
		accessKeyId     string
		secretAccessKey string
		bucket          string
		region          string
	}
)

func NewAwsS3Parameters() contracts.AwsS3Parameters {
	return &awsS3ParametersImpl{
		os.Getenv(AwsS3AccessKeyIdParameter),
		os.Getenv(AwsS3SecretAccessKeyParameter),
		os.Getenv(AwsS3BucketParameter),
		os.Getenv(AwsS3RegionParameter),
	}
}

func (p *awsS3ParametersImpl) AccessKeyId() string {
	return p.accessKeyId
}

func (p *awsS3ParametersImpl) SecretKeyId() string {
	return p.secretAccessKey
}

func (p *awsS3ParametersImpl) Bucket() string {
	return p.bucket
}

func (p *awsS3ParametersImpl) Region() string {
	return p.region
}

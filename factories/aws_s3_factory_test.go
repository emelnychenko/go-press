package factories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAwsS3Factory(t *testing.T) {
	t.Run("NewAwsS3Factory", func(t *testing.T) {
		_, isAwsS3Factory := NewAwsS3Factory().(*awsS3FactoryImpl)

		assert.True(t, isAwsS3Factory)
	})

	t.Run("Create", func(t *testing.T) {
		awsSdkSession, _ := session.NewSession()
		awsS3Factory := &awsS3FactoryImpl{}

		_, isAwsSdkS3 := awsS3Factory.Create(awsSdkSession).(*s3.S3)
		assert.True(t, isAwsSdkS3)
	})
}

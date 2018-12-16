package factories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAwsS3UploaderFactory(t *testing.T) {
	t.Run("NewAwsS3UploaderFactory", func(t *testing.T) {
		_, isAwsS3UploaderFactoryImpl := NewAwsS3UploaderFactory().(*awsS3UploaderFactoryImpl)

		assert.True(t, isAwsS3UploaderFactoryImpl)
	})

	t.Run("Create", func(t *testing.T) {
		sess, _ := session.NewSession()
		awsS3UploaderFactory := &awsS3UploaderFactoryImpl{}

		_, isUploader := awsS3UploaderFactory.Create(sess).(*s3manager.Uploader)
		assert.True(t, isUploader)
	})
}

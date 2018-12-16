package factories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAwsS3DownloaderFactory(t *testing.T) {
	t.Run("NewAwsS3DownloaderFactory", func(t *testing.T) {
		_, isAwsS3DownloaderFactory := NewAwsS3DownloaderFactory().(*awsS3DownloaderFactoryImpl)

		assert.True(t, isAwsS3DownloaderFactory)
	})

	t.Run("Create", func(t *testing.T) {
		sess, _ := session.NewSession()
		awsS3DownloaderFactory := &awsS3DownloaderFactoryImpl{}

		_, isDownloader := awsS3DownloaderFactory.Create(sess).(*s3manager.Downloader)
		assert.True(t, isDownloader)
	})
}

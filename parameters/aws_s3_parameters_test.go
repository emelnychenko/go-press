package parameters

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAwsS3Parameters(t *testing.T) {
	const awsS3AccessKeyIdParameter = "0"
	const awsS3SecretAccessKeyParameter = "1"
	const awsS3BucketParameter = "2"
	const awsS3RegionParameter = "3"

	t.Run("NewAwsS3Parameters", func(t *testing.T) {
		_ = os.Setenv(AwsS3AccessKeyIdParameter, awsS3AccessKeyIdParameter)
		_ = os.Setenv(AwsS3SecretAccessKeyParameter, awsS3SecretAccessKeyParameter)
		_ = os.Setenv(AwsS3BucketParameter, awsS3BucketParameter)
		_ = os.Setenv(AwsS3RegionParameter, awsS3RegionParameter)

		awsS3Parameters, isAwsS3Parameters := NewAwsS3Parameters().(*awsS3ParametersImpl)
		assert.True(t, isAwsS3Parameters)
		assert.Equal(t, awsS3AccessKeyIdParameter, awsS3Parameters.accessKeyId)
		assert.Equal(t, awsS3SecretAccessKeyParameter, awsS3Parameters.secretAccessKey)
		assert.Equal(t, awsS3BucketParameter, awsS3Parameters.bucket)
		assert.Equal(t, awsS3RegionParameter, awsS3Parameters.region)
	})

	t.Run("AccessKeyId", func(t *testing.T) {
		awsS3Parameters := &awsS3ParametersImpl{accessKeyId: awsS3AccessKeyIdParameter}
		assert.Equal(t, awsS3AccessKeyIdParameter, awsS3Parameters.AccessKeyId())
	})

	t.Run("SecretKeyId", func(t *testing.T) {
		awsS3Parameters := &awsS3ParametersImpl{secretAccessKey: awsS3SecretAccessKeyParameter}
		assert.Equal(t, awsS3SecretAccessKeyParameter, awsS3Parameters.SecretKeyId())
	})

	t.Run("Region", func(t *testing.T) {
		awsS3Parameters := &awsS3ParametersImpl{bucket: awsS3BucketParameter}
		assert.Equal(t, awsS3BucketParameter, awsS3Parameters.Bucket())
	})

	t.Run("Bucket", func(t *testing.T) {
		awsS3Parameters := &awsS3ParametersImpl{region: awsS3RegionParameter}
		assert.Equal(t, awsS3RegionParameter, awsS3Parameters.Region())
	})
}

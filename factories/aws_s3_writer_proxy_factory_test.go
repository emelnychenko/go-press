package factories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAwsS3WriterProxyFactory(t *testing.T) {
	t.Run("NewAwsS3WriterProxyFactory", func(t *testing.T) {
		_, isAwsS3WriterProxyFactory := NewAwsS3WriterProxyFactory().(*awsS3WriterProxyFactoryImpl)

		assert.True(t, isAwsS3WriterProxyFactory)
	})

	t.Run("Create", func(t *testing.T) {
		awsS3WriterProxyFactory := &awsS3WriterProxyFactoryImpl{}
		assert.NotNil(t, awsS3WriterProxyFactory.Create(nil))
	})
}

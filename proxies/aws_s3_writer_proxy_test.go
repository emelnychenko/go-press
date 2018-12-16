package proxies

import (
	"bufio"
	"bytes"
	"testing/iotest"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAwsS3WriterProxy(t *testing.T) {
	t.Run("NewAwsS3WriterProxy", func(t *testing.T) {
		writer := new(bufio.Writer)
		var awsS3WriterProxy = NewAwsS3WriterProxy(writer)
		assert.IsType(t, new(awsS3WriterProxyImpl), awsS3WriterProxy)
		assert.Equal(t, writer, awsS3WriterProxy.Writer())
	})

	t.Run("WriteAt", func(t *testing.T) {
		writer := iotest.NewWriteLogger("AwsS3WriterProxy", bytes.NewBuffer(nil))
		awsS3WriterProxy := NewAwsS3WriterProxy(writer)

		chunk := []byte("test")
		off := 5
		n, err := awsS3WriterProxy.WriteAt(chunk, int64(off))
		assert.Equal(t, len(chunk)+off, n)
		assert.Nil(t, err)
	})

	t.Run("WriteAt:Error", func(t *testing.T) {
		writer := iotest.NewWriteLogger("AwsS3WriterProxy", bytes.NewBuffer(nil))
		awsS3WriterProxy := NewAwsS3WriterProxy(writer)
		_, err := awsS3WriterProxy.WriteAt([]byte("test"), -1)
		assert.Error(t, err)
	})
}

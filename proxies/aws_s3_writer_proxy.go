package proxies

import (
	"github.com/emelnychenko/go-press/contracts"
	"io"
)

type (
	awsS3WriterProxyImpl struct {
		cur    int64
		writer io.Writer
	}
)

func NewAwsS3WriterProxy(writer io.Writer) contracts.AwsS3WriterProxy {
	return &awsS3WriterProxyImpl{0, writer}
}

func (a *awsS3WriterProxyImpl) Writer() io.Writer {
	return a.writer
}

func (a *awsS3WriterProxyImpl) WriteAt(p []byte, off int64) (n int, err error) {
	if off < a.cur {
		return 0, io.ErrUnexpectedEOF
	}

	s := make([]byte, off-a.cur)
	p = append(s, p...)
	n, err = a.writer.Write(p)
	a.cur += int64(n)
	return
}

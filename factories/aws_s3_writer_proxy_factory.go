package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/proxies"
	"io"
)

type (
	awsS3WriterProxyFactoryImpl struct {
	}
)

func NewAwsS3WriterProxyFactory() contracts.AwsS3WriterProxyFactory {
	return &awsS3WriterProxyFactoryImpl{}
}

func (*awsS3WriterProxyFactoryImpl) Create(destination io.Writer) contracts.AwsS3WriterProxy {
	return proxies.NewAwsS3WriterProxy(destination)
}

package contracts

import "io"

type (
	AwsS3WriterProxyFactory interface {
		Create(destination io.Writer) AwsS3WriterProxy
	}
)

package contracts

import "io"

type (
	AwsS3WriterProxy interface {
		io.WriterAt
		Writer() io.Writer
	}
)

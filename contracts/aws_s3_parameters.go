package contracts

type (
	AwsS3Parameters interface {
		AccessKeyId() string
		SecretKeyId() string
		Bucket() string
		Region() string
	}
)

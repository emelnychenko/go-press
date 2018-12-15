package enums

type (
	PostPrivacy string
)

const (
	PostPublicPrivacy PostPrivacy = "public"
	PostProtectedPrivacy PostPrivacy = "protected"
	PostPrivatePrivacy PostPrivacy = "private"
)

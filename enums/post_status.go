package enums

type (
	PostStatus string
)

const (
	PostPublishedStatus PostStatus = "published"
	PostScheduledStatus PostStatus = "scheduled"
	PostDraftStatus     PostStatus = "draft"
)

package resolvers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	subjectResolverImpl struct {
		userApi contracts.UserApi
	}
)

func NewSubjectResolver(userApi contracts.UserApi) (subjectResolver contracts.SubjectResolver) {
	return &subjectResolverImpl{userApi}
}

func (c *subjectResolverImpl) ResolveSubject(subjectId *models.ModelId, subjectType models.SubjectType) (subject interface{}, err errors.Error) {
	switch subjectType {
	case models.SystemSubjectType:
		subject = models.NewSystemUser()
	case models.UserSubjectType:
		subject, err = c.userApi.GetUser(subjectId)
	default:
		err = errors.NewSystemError("SubjectType out of range")
	}

	return
}

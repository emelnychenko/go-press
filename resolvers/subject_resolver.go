package resolvers

import (
	"errors"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/enums"
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

func (c *subjectResolverImpl) ResolveSubject(subjectId *common.ModelId, subjectType enums.SubjectType) (subject interface{}, err common.Error) {
	switch subjectType {
	case enums.SystemSubjectType:
		subject = models.NewSystemUser()
	case enums.UserSubjectType:
		subject, err = c.userApi.GetUser(subjectId)
	default:
		stringErr := errors.New("SubjectType out of range")
		err = common.NewSystemError(stringErr)
	}

	return
}

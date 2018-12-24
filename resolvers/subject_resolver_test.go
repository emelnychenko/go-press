package resolvers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubjectResolver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewSubjectResolver", func(t *testing.T) {
		userApi := mocks.NewMockUserApi(ctrl)
		subjectResolver, isSubjectResolver := NewSubjectResolver(userApi).(*subjectResolverImpl)

		assert.True(t, isSubjectResolver)
		assert.Equal(t, userApi, subjectResolver.userApi)
	})

	t.Run("ResolveSubject:SystemUser", func(t *testing.T) {
		subjectResolver := &subjectResolverImpl{}
		subject, err := subjectResolver.ResolveSubject(nil, enums.SystemSubjectType)

		assert.IsType(t, new(models.SystemUser), subject)
		assert.Nil(t, err)
	})

	t.Run("ResolveSubject:UserEntity", func(t *testing.T) {
		userId := common.NewModelId()
		user := new(models.User)
		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().GetUser(userId).Return(user, nil)

		subjectResolver := &subjectResolverImpl{userApi: userApi}
		subject, err := subjectResolver.ResolveSubject(userId, enums.UserSubjectType)

		assert.Equal(t, user, subject)
		assert.Nil(t, err)
	})

	t.Run("ResolveSubject:Error", func(t *testing.T) {
		subjectResolver := &subjectResolverImpl{}

		subject, err := subjectResolver.ResolveSubject(nil, "none")
		assert.Nil(t, subject)
		assert.Error(t, err)
	})
}

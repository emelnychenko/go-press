package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEchoHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	t.Run("NewUserEchoHelper", func(t *testing.T) {
		_, isUserParamParser := NewUserEchoHelper().(*userEchoHelperImpl)
		assert.True(t, isUserParamParser)
	})

	t.Run("ParseUserId", func(t *testing.T) {
		user := new(models.UserId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(UserIdParameterName).Return(user.String())

		userEchoHelper := &userEchoHelperImpl{}
		parsedUserId, err := userEchoHelper.ParseUserId(httpContext)
		assert.Equal(t, user.String(), parsedUserId.String())
		assert.Nil(t, err)
	})
}

package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewChannelHttpHelper", func(t *testing.T) {
		_, isChannelHttpHelper := NewChannelHttpHelper().(*channelHttpHelperImpl)
		assert.True(t, isChannelHttpHelper)
	})

	t.Run("ParseChannelId", func(t *testing.T) {
		channelId := new(models.ChannelId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(ChannelIdParameterName).Return(channelId.String())

		channelHttpHelper := &channelHttpHelperImpl{}
		parsedChannelId, err := channelHttpHelper.ParseChannelId(httpContext)
		assert.Equal(t, channelId.String(), parsedChannelId.String())
		assert.Nil(t, err)
	})
}

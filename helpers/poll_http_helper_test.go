package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPollHttpHelper", func(t *testing.T) {
		_, isPollHttpHelper := NewPollHttpHelper().(*pollHttpHelperImpl)
		assert.True(t, isPollHttpHelper)
	})

	t.Run("ParsePollId", func(t *testing.T) {
		poll := new(models.PollId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(PollIdParameterName).Return(poll.String())

		pollHttpHelper := &pollHttpHelperImpl{}
		parsedPollId, err := pollHttpHelper.ParsePollId(httpContext)
		assert.Equal(t, poll.String(), parsedPollId.String())
		assert.Nil(t, err)
	})
}

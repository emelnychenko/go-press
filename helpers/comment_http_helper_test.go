package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCommentHttpHelper", func(t *testing.T) {
		_, isCommentHttpHelper := NewCommentHttpHelper().(*commentHttpHelperImpl)
		assert.True(t, isCommentHttpHelper)
	})

	t.Run("ParseCommentId", func(t *testing.T) {
		comment := new(models.CommentId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(CommentIdParameterName).Return(comment.String())

		commentHttpHelper := &commentHttpHelperImpl{}
		parsedCommentId, err := commentHttpHelper.ParseCommentId(httpContext)
		assert.Equal(t, comment.String(), parsedCommentId.String())
		assert.Nil(t, err)
	})
}

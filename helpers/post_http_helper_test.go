package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostHttpHelper", func(t *testing.T) {
		_, isPostHttpHelper := NewPostHttpHelper().(*postHttpHelperImpl)
		assert.True(t, isPostHttpHelper)
	})

	t.Run("ParsePostId", func(t *testing.T) {
		post := new(models.PostId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(PostIdParameterName).Return(post.String())

		postHttpHelper := &postHttpHelperImpl{}
		parsedPostId, err := postHttpHelper.ParsePostId(httpContext)
		assert.Equal(t, post.String(), parsedPostId.String())
		assert.Nil(t, err)
	})
}

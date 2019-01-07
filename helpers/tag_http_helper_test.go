package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewTagHttpHelper", func(t *testing.T) {
		_, isTagHttpHelper := NewTagHttpHelper().(*tagHttpHelperImpl)
		assert.True(t, isTagHttpHelper)
	})

	t.Run("ParseTagId", func(t *testing.T) {
		tagId := new(models.TagId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(TagIdParameterName).Return(tagId.String())

		tagHttpHelper := &tagHttpHelperImpl{}
		parsedTagId, err := tagHttpHelper.ParseTagId(httpContext)
		assert.Equal(t, tagId.String(), parsedTagId.String())
		assert.Nil(t, err)
	})
}

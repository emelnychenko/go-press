package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	t.Run("NewCategoryHttpHelper", func(t *testing.T) {
		_, isCategoryHttpHelper := NewCategoryHttpHelper().(*categoryHttpHelperImpl)
		assert.True(t, isCategoryHttpHelper)
	})

	t.Run("ParseCategoryId", func(t *testing.T) {
		category := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(CategoryIdParameterName).Return(category.String())

		categoryHttpHelper := &categoryHttpHelperImpl{}
		parsedCategoryId, err := categoryHttpHelper.ParseCategoryId(httpContext)
		assert.Equal(t, category.String(), parsedCategoryId.String())
		assert.Nil(t, err)
	})
}

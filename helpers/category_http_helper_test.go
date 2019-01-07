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
		categoryId := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(CategoryIdParameterName).Return(categoryId.String())

		categoryHttpHelper := &categoryHttpHelperImpl{}
		parsedCategoryId, err := categoryHttpHelper.ParseCategoryId(httpContext)
		assert.Equal(t, categoryId.String(), parsedCategoryId.String())
		assert.Nil(t, err)
	})

	t.Run("ParseParentCategoryId", func(t *testing.T) {
		categoryParentId := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(ParentCategoryIdParameterName).Return(categoryParentId.String())

		categoryHttpHelper := &categoryHttpHelperImpl{}
		parsedCategoryId, err := categoryHttpHelper.ParseParentCategoryId(httpContext)
		assert.Equal(t, categoryParentId.String(), parsedCategoryId.String())
		assert.Nil(t, err)
	})
}

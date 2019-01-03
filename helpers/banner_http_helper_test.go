package helpers

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewBannerHttpHelper", func(t *testing.T) {
		_, isBannerHttpHelper := NewBannerHttpHelper().(*bannerHttpHelperImpl)
		assert.True(t, isBannerHttpHelper)
	})

	t.Run("ParseBannerId", func(t *testing.T) {
		banner := new(models.BannerId)
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(BannerIdParameterName).Return(banner.String())

		bannerHttpHelper := &bannerHttpHelperImpl{}
		parsedBannerId, err := bannerHttpHelper.ParseBannerId(httpContext)
		assert.Equal(t, banner.String(), parsedBannerId.String())
		assert.Nil(t, err)
	})
}

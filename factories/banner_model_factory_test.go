package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewBannerModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		bannerModelFactory, isBannerModelFactory := NewBannerModelFactory(paginationModelFactory).(*bannerModelFactoryImpl)

		assert.True(t, isBannerModelFactory)
		assert.Equal(t, paginationModelFactory, bannerModelFactory.paginationModelFactory)
	})

	t.Run("CreateBannerPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		bannerModelFactory := &bannerModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		bannerPaginationQuery := bannerModelFactory.CreateBannerPaginationQuery()

		assert.Equal(t, paginationQuery, bannerPaginationQuery.PaginationQuery)
	})

	t.Run("CreateBanner", func(t *testing.T) {
		bannerModelFactory := new(bannerModelFactoryImpl)
		assert.NotNil(t, bannerModelFactory.CreateBanner())
	})

	t.Run("CreateBannerCreate", func(t *testing.T) {
		bannerModelFactory := new(bannerModelFactoryImpl)
		assert.NotNil(t, bannerModelFactory.CreateBannerCreate())
	})

	t.Run("CreateBannerUpdate", func(t *testing.T) {
		bannerModelFactory := new(bannerModelFactoryImpl)
		assert.NotNil(t, bannerModelFactory.CreateBannerUpdate())
	})
}

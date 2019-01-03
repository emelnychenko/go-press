package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	bannerApiImpl struct {
		eventDispatcher  contracts.EventDispatcher
		bannerEventFactory contracts.BannerEventFactory
		bannerService      contracts.BannerService
		bannerAggregator   contracts.BannerAggregator
	}
)

func NewBannerApi(
	eventDispatcher contracts.EventDispatcher,
	bannerEventFactory contracts.BannerEventFactory,
	bannerService contracts.BannerService,
	bannerAggregator contracts.BannerAggregator,
) (bannerApi contracts.BannerApi) {
	return &bannerApiImpl{
		eventDispatcher,
		bannerEventFactory,
		bannerService,
		bannerAggregator,
	}
}

func (a *bannerApiImpl) ListBanners(
	bannerPaginationQuery *models.BannerPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	entityPaginationResult, err := a.bannerService.ListBanners(bannerPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.bannerAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

func (a *bannerApiImpl) GetBanner(bannerId *models.BannerId) (banner *models.Banner, err common.Error) {
	bannerEntity, err := a.bannerService.GetBanner(bannerId)

	if nil != err {
		return
	}

	banner = a.bannerAggregator.AggregateBanner(bannerEntity)
	return
}

func (a *bannerApiImpl) CreateBanner(data *models.BannerCreate) (banner *models.Banner, err common.Error) {
	bannerEntity, err := a.bannerService.CreateBanner(data)

	if nil != err {
		return
	}

	bannerCreatedEvent := a.bannerEventFactory.CreateBannerCreatedEvent(bannerEntity)
	a.eventDispatcher.Dispatch(bannerCreatedEvent)

	banner = a.bannerAggregator.AggregateBanner(bannerEntity)
	return
}

func (a *bannerApiImpl) UpdateBanner(bannerId *models.BannerId, data *models.BannerUpdate) (err common.Error) {
	bannerService := a.bannerService
	bannerEntity, err := bannerService.GetBanner(bannerId)

	if nil != err {
		return
	}

	err = bannerService.UpdateBanner(bannerEntity, data)

	if nil != err {
		return
	}

	bannerUpdatedEvent := a.bannerEventFactory.CreateBannerUpdatedEvent(bannerEntity)
	a.eventDispatcher.Dispatch(bannerUpdatedEvent)
	return
}

func (a *bannerApiImpl) DeleteBanner(bannerId *models.BannerId) (err common.Error) {
	bannerService := a.bannerService
	bannerEntity, err := bannerService.GetBanner(bannerId)

	if nil != err {
		return
	}

	err = bannerService.DeleteBanner(bannerEntity)

	if nil != err {
		return
	}

	bannerDeletedEvent := a.bannerEventFactory.CreateBannerDeletedEvent(bannerEntity)
	a.eventDispatcher.Dispatch(bannerDeletedEvent)

	return
}

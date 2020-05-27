package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IBannerService interface {
	GetBannerList(status, page, size int) (*[]model.WechatMallBannerDO, int)
	GetBannerById(id int) *model.WechatMallBannerDO
	AddBanner(banner *model.WechatMallBannerDO)
	UpdateBannerById(banner *model.WechatMallBannerDO)
}

type bannerService struct {
}

func NewBannerService() IBannerService {
	service := &bannerService{}
	return service
}

func (bs *bannerService) GetBannerList(status, page, size int) (*[]model.WechatMallBannerDO, int) {
	bannerList, err := dbops.BannerDao.List(status, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.BannerDao.CountByStatus(status)
	if err != nil {
		panic(err)
	}
	return bannerList, total
}

func (bs *bannerService) GetBannerById(id int) *model.WechatMallBannerDO {
	banner, err := dbops.BannerDao.QueryBanner(id)
	if err != nil {
		panic(err)
	}
	return banner
}

func (bs *bannerService) AddBanner(banner *model.WechatMallBannerDO) {
	_, err := dbops.BannerDao.Insert(banner)
	if err != nil {
		panic(err)
	}
}

func (bs *bannerService) UpdateBannerById(banner *model.WechatMallBannerDO) {
	err := dbops.BannerDao.Update(banner)
	if err != nil {
		panic(err)
	}
}

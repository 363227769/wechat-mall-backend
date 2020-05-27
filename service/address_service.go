package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

type IAddressService interface {
	GetAddressList(userId, page, size int) (*[]model.WechatMallUserAddressDO, int)
	GetAddress(id int) *model.WechatMallUserAddressDO
	GetDefaultAddress(userId int) *model.WechatMallUserAddressDO
	AddAddress(address *model.WechatMallUserAddressDO)
	UpdateAddress(address *model.WechatMallUserAddressDO)
}

type addressService struct {
}

func NewAddressService() IAddressService {
	service := addressService{}
	return &service
}

func (s *addressService) GetAddressList(userId, page, size int) (*[]model.WechatMallUserAddressDO, int) {
	addressList, err := dbops.UserAddressDao.ListByUserId(userId, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.UserAddressDao.CountByUserId(userId)
	if err != nil {
		panic(err)
	}
	return addressList, total
}

func (s *addressService) GetAddress(id int) *model.WechatMallUserAddressDO {
	addressDO, err := dbops.UserAddressDao.QueryById(id)
	if err != nil {
		panic(err)
	}
	return addressDO
}

func (s *addressService) GetDefaultAddress(userId int) *model.WechatMallUserAddressDO {
	addressDO, err := dbops.UserAddressDao.QueryDefaultAddress(userId)
	if err != nil {
		panic(err)
	}
	return addressDO
}

func (s *addressService) AddAddress(address *model.WechatMallUserAddressDO) {
	if address.IsDefault == 1 {
		clearDefaultAddress(address.UserId)
	}
	err := dbops.UserAddressDao.Insert(address)
	if err != nil {
		panic(err)
	}
}

func (s *addressService) UpdateAddress(address *model.WechatMallUserAddressDO) {
	if address.IsDefault == 1 {
		clearDefaultAddress(address.UserId)
	}
	err := dbops.UserAddressDao.Update(address)
	if err != nil {
		panic(err)
	}
}

func clearDefaultAddress(userId int) {
	addressDO, err := dbops.UserAddressDao.QueryDefaultAddress(userId)
	if err != nil {
		panic(err)
	}
	if addressDO.Id != defs.ZERO {
		addressDO.IsDefault = 0
		err = dbops.UserAddressDao.Update(addressDO)
		if err != nil {
			panic(err)
		}
	}
}

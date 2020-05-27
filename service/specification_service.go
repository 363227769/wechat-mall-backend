package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ISpecificationService interface {
	GetSpecificationList(page, size int) (*[]model.WechatMallSpecificationDO, int)
	GetSpecificationById(id int) *model.WechatMallSpecificationDO
	GetSpecificationByName(name string) *model.WechatMallSpecificationDO
	UpdateSpecificationById(spec *model.WechatMallSpecificationDO)
	AddSpecification(spec *model.WechatMallSpecificationDO)
	GetSpecificationAttrList(specId int) *[]model.WechatMallSpecificationAttrDO
	GetSpecificationAttrById(id int) *model.WechatMallSpecificationAttrDO
	GetSpecificationAttrByValue(value string) *model.WechatMallSpecificationAttrDO
	UpdateSpecificationAttrById(spec *model.WechatMallSpecificationAttrDO)
	AddSpecificationAttr(spec *model.WechatMallSpecificationAttrDO)
}

type specificationService struct {
}

func NewSpecificationService() ISpecificationService {
	service := specificationService{}
	return &service
}

func (ss *specificationService) GetSpecificationList(page, size int) (*[]model.WechatMallSpecificationDO, int) {
	specList, err := dbops.SpecDao.List(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.SpecDao.Count()
	if err != nil {
		panic(err)
	}
	return specList, total
}

func (ss *specificationService) GetSpecificationById(id int) *model.WechatMallSpecificationDO {
	spec, err := dbops.SpecDao.QueryById(id)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) GetSpecificationByName(name string) *model.WechatMallSpecificationDO {
	spec, err := dbops.SpecDao.QueryByName(name)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) UpdateSpecificationById(spec *model.WechatMallSpecificationDO) {
	err := dbops.SpecDao.UpdateById(spec)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) AddSpecification(spec *model.WechatMallSpecificationDO) {
	err := dbops.SpecDao.Insert(spec)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) GetSpecificationAttrList(specId int) *[]model.WechatMallSpecificationAttrDO {
	attrList, err := dbops.SpecAttrDao.ListBySpecId(specId)
	if err != nil {
		panic(err)
	}
	return attrList
}

func (ss *specificationService) GetSpecificationAttrById(id int) *model.WechatMallSpecificationAttrDO {
	attr, err := dbops.SpecAttrDao.QueryById(id)
	if err != nil {
		panic(err)
	}
	return attr
}

func (ss *specificationService) GetSpecificationAttrByValue(value string) *model.WechatMallSpecificationAttrDO {
	spec, err := dbops.SpecAttrDao.QueryByValue(value)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) UpdateSpecificationAttrById(attr *model.WechatMallSpecificationAttrDO) {
	err := dbops.SpecAttrDao.UpdateById(attr)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) AddSpecificationAttr(attr *model.WechatMallSpecificationAttrDO) {
	err := dbops.SpecAttrDao.Insert(attr)
	if err != nil {
		panic(err)
	}
}

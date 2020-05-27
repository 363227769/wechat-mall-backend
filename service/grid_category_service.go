package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IGridCategoryService interface {
	GetGridCategoryList(page, size int) (*[]model.WechatMallGridCategoryDO, int)
	GetGridCategoryById(id int) *model.WechatMallGridCategoryDO
	GetGridCategoryByName(name string) *model.WechatMallGridCategoryDO
	AddGridCategory(gridC *model.WechatMallGridCategoryDO)
	UpdateGridCategory(gridC *model.WechatMallGridCategoryDO)
	CountCategoryBindGrid(categoryId int) int
}

type gridCategoryService struct {
}

func NewGridCategoryService() IGridCategoryService {
	service := gridCategoryService{}
	return &service
}

func (g *gridCategoryService) GetGridCategoryList(page, size int) (*[]model.WechatMallGridCategoryDO, int) {
	gridCList, err := dbops.GridCategoryDao.List(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.GridCategoryDao.Count()
	if err != nil {
		panic(err)
	}
	return gridCList, total
}

func (g *gridCategoryService) GetGridCategoryById(id int) *model.WechatMallGridCategoryDO {
	gridC, err := dbops.GridCategoryDao.QueryById(id)
	if err != nil {
		panic(err)
	}
	return gridC
}

func (g *gridCategoryService) GetGridCategoryByName(name string) *model.WechatMallGridCategoryDO {
	gridC, err := dbops.GridCategoryDao.QueryByName(name)
	if err != nil {
		panic(err)
	}
	return gridC
}

func (g *gridCategoryService) AddGridCategory(gridC *model.WechatMallGridCategoryDO) {
	err := dbops.GridCategoryDao.Insert(gridC)
	if err != nil {
		panic(err)
	}
}

func (g *gridCategoryService) UpdateGridCategory(gridC *model.WechatMallGridCategoryDO) {
	err := dbops.GridCategoryDao.Update(gridC)
	if err != nil {
		panic(err)
	}
}

// 统计分类绑定的宫格
func (g *gridCategoryService) CountCategoryBindGrid(categoryId int) int {
	total, err := dbops.GridCategoryDao.CountByCategoryId(categoryId)
	if err != nil {
		panic(err)
	}
	return total
}

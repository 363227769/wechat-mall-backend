package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

type IBrowseRecordService interface {
	AddBrowseRecord(record *model.WechatMallGoodsBrowseRecord)
	ListBrowseRecord(userId, page, size int) (*[]defs.PortalBrowseRecordVO, int)
	ClearBrowseHistory(ids []int)
}

type browseRecordService struct{}

func NewBrowseRecordService() IBrowseRecordService {
	service := &browseRecordService{}
	return service
}

func (s *browseRecordService) AddBrowseRecord(record *model.WechatMallGoodsBrowseRecord) {
	recordDO, err := dbops.BrowseRecordDao.List(record.UserId, record.GoodsId)
	if err != nil {
		panic(err)
	}
	if recordDO.Id != 0 {
		err := dbops.BrowseRecordDao.DeleteById(recordDO.Id)
		if err != nil {
			panic(err)
		}
	}
	err = dbops.BrowseRecordDao.Insert(record)
	if err != nil {
		panic(err)
	}
}

func (s *browseRecordService) ListBrowseRecord(userId, page, size int) (*[]defs.PortalBrowseRecordVO, int) {
	records, err := dbops.BrowseRecordDao.ListByPage(userId, page, size)
	if err != nil {
		panic(err)
	}
	recordVOs := []defs.PortalBrowseRecordVO{}
	for _, recordDO := range *records {
		recordVO := defs.PortalBrowseRecordVO{}
		recordVO.Id = recordDO.Id
		recordVO.GoodsId = recordDO.GoodsId
		recordVO.Picture = recordDO.Picture
		recordVO.Title = recordDO.Title
		recordVO.Price = recordDO.Price
		recordVO.BrowseTime = recordDO.UpdateTime
		recordVOs = append(recordVOs, recordVO)
	}
	total, err := dbops.BrowseRecordDao.CountByUserId(userId)
	if err != nil {
		panic(err)
	}
	return &recordVOs, total
}

func (s *browseRecordService) ClearBrowseHistory(ids []int) {
	for _, v := range ids {
		err := dbops.BrowseRecordDao.DeleteById(v)
		if err != nil {
			panic(err)
		}
	}
}

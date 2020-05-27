package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

type goodsDao struct{}

var GoodsDao = new(goodsDao)

const goodsColumnList = `
id, brand_name, title, price, discount_price, category_id, online, picture, 
banner_picture, detail_picture, tags, sale_num, is_del, create_time, update_time
`

func (*goodsDao) List(keyword, order string, categoryId, online, page, size int) (*[]model.WechatMallGoodsDO, error) {
	sql := "SELECT " + goodsColumnList + " FROM wechat_mall_goods WHERE is_del = 0"
	if keyword != "" {
		sql += " AND title LIKE '%" + keyword + "%'"
	}
	if categoryId != defs.ALL {
		sql += " AND category_id = " + strconv.Itoa(categoryId)
	}
	if online != defs.ALL {
		sql += " AND online = " + strconv.Itoa(online)
	}
	if order != "" {
		sql += " ORDER BY " + order + " DESC "
	}
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	goodsList := []model.WechatMallGoodsDO{}
	for rows.Next() {
		goods := model.WechatMallGoodsDO{}
		err := rows.Scan(&goods.Id, &goods.BrandName, &goods.Title, &goods.Price, &goods.DiscountPrice,
			&goods.CategoryId, &goods.Online, &goods.Picture, &goods.BannerPicture, &goods.DetailPicture,
			&goods.Tags, &goods.SaleNum, &goods.Del, &goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
		goodsList = append(goodsList, goods)
	}
	return &goodsList, nil
}

func (*goodsDao) Count(keyword string, categoryId, online int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_goods WHERE is_del = 0"
	if keyword != "" {
		sql += " AND title LIKE '%" + keyword + "%'"
	}
	if categoryId != defs.ALL {
		sql += " AND category_id = " + strconv.Itoa(categoryId)
	}
	if online != defs.ALL {
		sql += " AND online = " + strconv.Itoa(online)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func (*goodsDao) Insert(goods *model.WechatMallGoodsDO) (int64, error) {
	sql := "INSERT INTO wechat_mall_goods ( " + goodsColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(goods.BrandName, goods.Title, goods.Price, goods.DiscountPrice,
		goods.CategoryId, goods.Online, goods.Picture, goods.BannerPicture, goods.DetailPicture,
		goods.Tags, 0, 0, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	i, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (*goodsDao) QueryById(id int) (*model.WechatMallGoodsDO, error) {
	sql := "SELECT " + goodsColumnList + " FROM wechat_mall_goods WHERE id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	goods := model.WechatMallGoodsDO{}
	for rows.Next() {
		err := rows.Scan(&goods.Id, &goods.BrandName, &goods.Title, &goods.Price,
			&goods.DiscountPrice, &goods.CategoryId, &goods.Online, &goods.Picture,
			&goods.BannerPicture, &goods.DetailPicture, &goods.Tags, &goods.SaleNum,
			&goods.Del, &goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &goods, nil
}

func (*goodsDao) UpdateById(goods *model.WechatMallGoodsDO) error {
	sql := `
UPDATE wechat_mall_goods 
SET brand_name = ?, title = ?, price = ?, discount_price = ?, category_id = ?,
online = ?, picture = ?, banner_picture = ?, detail_picture = ?, tags = ?, 
sale_num = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goods.BrandName, goods.Title, goods.Price, goods.DiscountPrice,
		goods.CategoryId, goods.Online, goods.Picture, goods.BannerPicture, goods.DetailPicture,
		goods.Tags, goods.SaleNum, goods.Del, time.Now(), goods.Id)
	if err != nil {
		return err
	}
	return nil
}

func (*goodsDao) CountByCategoryId(categoryId int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_goods WHERE is_del = 0 AND category_id = " + strconv.Itoa(categoryId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func (*goodsDao) UpdateOnlineStatus(categoryId, online int) error {
	sql := "UPDATE wechat_mall_goods SET update_time = now(), online = " + strconv.Itoa(online) + " WHERE is_del = 0 AND category_id = " + strconv.Itoa(categoryId)
	_, err := dbConn.Exec(sql)
	return err
}

func (*goodsDao) UpdateSaleNum(goodsId, num int) error {
	sql := "UPDATE wechat_mall_goods SET sale_num = sale_num + " + strconv.Itoa(num) + " WHERE id = " + strconv.Itoa(goodsId)
	_, err := dbConn.Exec(sql)
	return err
}

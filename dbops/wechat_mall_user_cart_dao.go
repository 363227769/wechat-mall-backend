package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

type cartDao struct{}

var CartDao = new(cartDao)

const cartColumnList = `
id, user_id, goods_id, sku_id, num, is_del, create_time, update_time
`

func (*cartDao) ListByUserId(userId, page, size int) (*[]model.WechatMallUserCartDO, error) {
	sql := "SELECT " + cartColumnList + " FROM wechat_mall_user_cart WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if page > 0 && size > 0 {
		sql += " ORDER BY update_time DESC LIMIT " + strconv.Itoa((page-1)*page) + " , " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	cartList := []model.WechatMallUserCartDO{}
	for rows.Next() {
		cartDO := model.WechatMallUserCartDO{}
		err := rows.Scan(&cartDO.Id, &cartDO.UserId, &cartDO.GoodsId, &cartDO.SkuId, &cartDO.Num, &cartDO.Del, &cartDO.CreateTime, &cartDO.UpdateTime)
		if err != nil {
			return nil, err
		}
		cartList = append(cartList, cartDO)
	}
	return &cartList, nil
}

// 购物车-数量
func (*cartDao) CountCartGoods(userId int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_user_cart WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, nil
		}
	}
	return total, nil
}

// 购物车-商品数量
func (*cartDao) CoundCartGoodsNum(userId int) (int, error) {
	sql := "SELECT IFNULL(SUM(num), 0) FROM wechat_mall_user_cart WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, nil
		}
	}
	return total, nil
}

func (*cartDao) Insert(cartDO *model.WechatMallUserCartDO) error {
	sql := "INSERT INTO wechat_mall_user_cart ( " + cartColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cartDO.UserId, cartDO.GoodsId, cartDO.SkuId, cartDO.Num, 0, time.Now(), time.Now())
	return err
}

func (*cartDao) QueryByParams(userId, goodsId, skuId int) (*model.WechatMallUserCartDO, error) {
	sql := "SELECT " + cartColumnList + " FROM wechat_mall_user_cart WHERE is_del = 0 AND user_id = ? AND goods_id = ? AND sku_id = ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(userId, goodsId, skuId)
	if err != nil {
		return nil, err
	}
	cartDO := model.WechatMallUserCartDO{}
	for rows.Next() {
		err := rows.Scan(&cartDO.Id, &cartDO.UserId, &cartDO.GoodsId, &cartDO.SkuId, &cartDO.Num, &cartDO.Del, &cartDO.CreateTime, &cartDO.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &cartDO, nil
}

func (*cartDao) UpdateById(cartDO *model.WechatMallUserCartDO) error {
	sql := `
UPDATE wechat_mall_user_cart 
SET user_id = ?, goods_id = ?, num = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cartDO.UserId, cartDO.GoodsId, cartDO.Num, cartDO.Del, time.Now(), cartDO.Id)
	return err
}

func (*cartDao) QueryById(id int) (*model.WechatMallUserCartDO, error) {
	sql := "SELECT " + cartColumnList + " FROM wechat_mall_user_cart WHERE id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	cartDO := model.WechatMallUserCartDO{}
	for rows.Next() {
		err := rows.Scan(&cartDO.Id, &cartDO.UserId, &cartDO.GoodsId, &cartDO.SkuId, &cartDO.Num, &cartDO.Del, &cartDO.CreateTime, &cartDO.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &cartDO, nil
}

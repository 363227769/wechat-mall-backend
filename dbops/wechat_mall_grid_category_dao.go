package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

type gridCategoryDao struct{}

var GridCategoryDao = new(gridCategoryDao)

const gridCategoryColumnList = `
id, title, name, category_id, picture, is_del, create_time, update_time
`

func (*gridCategoryDao) List(page, size int) (*[]model.WechatMallGridCategoryDO, error) {
	sql := "SELECT " + gridCategoryColumnList + " FROM wechat_mall_grid_category WHERE is_del = 0"
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + " ," + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var gridCList []model.WechatMallGridCategoryDO
	for rows.Next() {
		gridC := model.WechatMallGridCategoryDO{}
		err := rows.Scan(&gridC.Id, &gridC.Title, &gridC.Name, &gridC.CategoryId, &gridC.Picture, &gridC.Del, &gridC.CreateTime, &gridC.UpdateTime)
		if err != nil {
			return nil, err
		}
		gridCList = append(gridCList, gridC)
	}
	return &gridCList, nil
}

func (*gridCategoryDao) Count() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_grid_category WHERE is_del = 0"
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

func (*gridCategoryDao) Insert(gridC *model.WechatMallGridCategoryDO) error {
	sql := "INSERT INTO wechat_mall_grid_category( " + gridCategoryColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(gridC.Title, gridC.Name, gridC.CategoryId, gridC.Picture, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (*gridCategoryDao) QueryById(id int) (*model.WechatMallGridCategoryDO, error) {
	sql := "SELECT " + gridCategoryColumnList + " FROM wechat_mall_grid_category WHERE id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	gridC := model.WechatMallGridCategoryDO{}
	for rows.Next() {
		err := rows.Scan(&gridC.Id, &gridC.Title, &gridC.Name, &gridC.CategoryId, &gridC.Picture, &gridC.Del, &gridC.CreateTime, &gridC.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &gridC, nil
}

func (*gridCategoryDao) QueryByName(name string) (*model.WechatMallGridCategoryDO, error) {
	sql := "SELECT " + gridCategoryColumnList + " FROM wechat_mall_grid_category WHERE is_del = 0 AND name = '" + name + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	gridC := model.WechatMallGridCategoryDO{}
	for rows.Next() {
		err := rows.Scan(&gridC.Id, &gridC.Title, &gridC.Name, &gridC.CategoryId, &gridC.Picture, &gridC.Del, &gridC.CreateTime, &gridC.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &gridC, nil
}

func (*gridCategoryDao) Update(gridC *model.WechatMallGridCategoryDO) error {
	sql := `
UPDATE wechat_mall_grid_category 
SET title = ?, name = ?, category_id = ?, picture = ?, is_del = ?, update_time = ? 
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(gridC.Title, gridC.Name, gridC.CategoryId, gridC.Picture, gridC.Del, time.Now(), gridC.Id)
	if err != nil {
		return err
	}
	return nil
}

func (*gridCategoryDao) CountByCategoryId(categoryId int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_grid_category WHERE is_del = 0 AND category_id = " + strconv.Itoa(categoryId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		panic(err)
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

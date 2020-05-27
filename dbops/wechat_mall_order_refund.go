package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

type orderRefundDao struct{}

var OrderRefundDao = new(orderRefundDao)

const refundColumnList = `
id, refund_no, user_id, order_no, reason, refund_amount, status, is_del, refund_time, create_time, update_time
`

func (*orderRefundDao) Insert(record *model.WechatMallOrderRefund) error {
	sql := "INSERT INTO wechat_mall_order_refund (" + refundColumnList[4:] + ") VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.RefundNo, record.UserId, record.OrderNo, record.Reason, record.RefundAmount, 0, 0, record.RefundTime, time.Now(), time.Now())
	return err
}

// 查询退款单
func (*orderRefundDao) QueryByRefundNo(refundNo string) (*model.WechatMallOrderRefund, error) {
	sql := "SELECT " + refundColumnList + " FROM wechat_mall_order_refund WHERE refund_no = '" + refundNo + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	record := model.WechatMallOrderRefund{}
	for rows.Next() {
		err := rows.Scan(&record.Id, &record.RefundNo, &record.UserId, &record.OrderNo, &record.Reason, &record.RefundAmount, &record.Status, &record.Del, &record.RefundTime, &record.CreateTime, &record.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &record, nil
}

// 订单退款记录
func (*orderRefundDao) QueryByOrderNo(orderNo string) (*model.WechatMallOrderRefund, error) {
	sql := "SELECT " + refundColumnList + " FROM wechat_mall_order_refund WHERE order_no = '" + orderNo + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	record := model.WechatMallOrderRefund{}
	for rows.Next() {
		err := rows.Scan(&record.Id, &record.RefundNo, &record.UserId, &record.OrderNo, &record.Reason, &record.RefundAmount, &record.Status, &record.Del, &record.RefundTime, &record.CreateTime, &record.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &record, nil
}

func (*orderRefundDao) Update(id int, status int) error {
	sql := "UPDATE wechat_mall_order_refund SET status = " + strconv.Itoa(status) + ", refund_time = now(), update_time = now() WHERE status = 0 AND id = " + strconv.Itoa(id)
	_, err := dbConn.Exec(sql)
	return err
}

func (*orderRefundDao) CountPendingOrderRefund() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_order_refund WHERE status IN (0, 1)"
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

package model

import (
	"StoreMS/database"
	"time"
)

type Orders struct {
	ID          int       `json:"id"`
	OrderNo     int64     `json:"order_no"`
	UserId      int       `json:"user_id"`
	ShippingId  int       `json:"shipping_id"`
	Payment     float64   `json:"payment"`
	Postage     float64   `json:"postage"`
	Status      int       `json:"status"` //10未支付，11已支付，20待发货，21已发货,.....
	PaymentTime time.Time `json:"payment_time"`
	SendTime    time.Time `json:"send_time"`
	EndTime     time.Time `json:"end_time"`
	CloseTime   time.Time `json:"close_time"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	OrderDetail []OrderDetails
}

type OrderDetails struct {
	ID               int       `json:"id"`
	OrderNo          int64     `json:"order_no"`
	UserId           int       `json:"user_id"`
	ProductId        int       `json:"product_id"`
	ProductName      string    `json:"product_name"`
	ProductImage     string    `json:"product_image"`
	CurrentUnitPrice float64   `json:"current_unit_price"`
	Quantity         int       `json:"quantity"`
	TotalPrice       float64   `json:"total_price"`
	CreateTime       time.Time `json:"create_time"`
	UpdateTime       time.Time `json:"update_time"`
}

type PayInfo struct {
	ID             int       `json:"id"`
	UserId         int       `json:"user_id"`
	OrderNo        int64     `json:"order_no"`
	PayPlatform    int       `json:"pay_platform"`
	PlatformNumber string    `json:"platform_number"`
	PlatformStatus string    `json:"platform_status"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

//提交订单表
//func (o *Orders) AddOrder1() bool {
//
//	tx := database.DB.Begin()
//	if err := tx.Model(&Orders{}).Create(o).Error; err == nil {
//		//提交
//		if err := tx.Model(&Cart{}).Where(Cart{Checked: 1}).Delete(Cart{}).Error; err == nil {
//			tx.Commit()
//			return true
//		} else {
//			tx.Rollback()
//			return false
//		}
//	} else {
//		//回滚
//		tx.Rollback()
//		return false
//	}
//}

//添加订单
func AddOrder(or Orders, od []OrderDetails) bool {
	tx := database.DB.Begin()

	if err := tx.Exec("insert into orders(order_no,user_id,shipping_id,payment,postage,status,create_time,update_time) values(?,?,?,?,?,?,?,?)", or.OrderNo, or.UserId, or.ShippingId, or.Payment, or.Postage, or.Status, or.PaymentTime).Error; err != nil {
		//回滚
		tx.Rollback()
		return false
	}
	for _, v := range od {
		if err := tx.Model(&OrderDetails{}).Create(&v).Error; err != nil {
			//回滚
			tx.Rollback()
			return false
		} else {
			//改变商品容量
			if err := tx.Model(&Product{}).Where(Product{ID: v.ProductId}).Exec("update products set stock=stock-?", v.Quantity).Error; err != nil {
				tx.Rollback()
				return false
			}
		}
	}
	//删除购物车
	if err := tx.Model(&Cart{}).Where(Cart{UserId: or.UserId, Checked: 1}).Delete(Cart{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//取消订单
func DeleteOrder(o int64) bool {
	tx := database.DB.Begin()

	if err := tx.Model(&Orders{}).Where(&Orders{OrderNo: o}).Delete(Orders{}).Error; err != nil {
		//回滚
		tx.Rollback()
		return false
	}
	if err := tx.Model(&OrderDetails{}).Where(&OrderDetails{OrderNo: o}).Delete(OrderDetails{}).Error; err != nil {
		//回滚
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//显示全部订单
func (o *Orders) PrintOrder() (info Orders, boole bool) {
	tx := database.DB.Begin()

	if err := tx.Model(&Orders{}).Where(&Orders{UserId: o.UserId}).Find(&info).Error; err != nil {
		//回滚
		tx.Rollback()
		boole = false
		return
	}
	if err := tx.Model(&OrderDetails{}).Where(&OrderDetails{UserId: o.UserId}).Order("quantity").Find(&info.OrderDetail).Error; err != nil {
		//回滚
		tx.Rollback()
		boole = false
		return
	}
	tx.Commit()
	boole = true
	return
}

//获取指定订单
func (or *Orders) GetOrder() (err error, info []Orders) {
	tx := database.DB.Begin()

	if err = tx.Model(Orders{}).Where(Orders{UserId: or.UserId, Status: or.Status}).Find(info).Error; err == nil {
		tx.Commit()
		return
	}
	tx.Rollback()
	return
}

//支付该订单
func Pay(info *Orders) bool {

	tx := database.DB.Begin()

	if err := tx.Model(&Orders{}).Where(Orders{OrderNo: info.OrderNo}).Update(info).Error; err != nil {
		tx.Rollback()
		return false
	} else {
		//积分减少(未传递userid)
		if err := tx.Model(&UserPoint{}).Where(UserPoint{UserID: 1}).Exec("update user_points set balance=balance-?", info.Payment).Error; err != nil {
			tx.Rollback()
			return false
		}
		tx.Commit()
		return true
	}
}

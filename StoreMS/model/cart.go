package model

import (
	"StoreMS/database"
	"time"
)

type Cart struct {
	ID         int       `json:"id"`
	UserId     int       `json:"user_id"`
	ProductId  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Checked    int       `json:"checked"` //0未选中，1选中
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

//添加购物车
func (c *Cart) AddCart() bool {

	tx := database.DB.Begin()

	if err := tx.Model(&Cart{}).Where("user_id=? and product_id=?", c.UserId, c.Quantity).Create(c).Error; err == nil {
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		return false
	}
}

//添加前验证，有则返回对象
func (c *Cart) IsIn() (info Cart, boole bool) {

	tx := database.DB.Begin()

	if tx.Model(&Cart{}).Where("user_id=? and product_id=?", c.UserId, c.Quantity).First(&info).RecordNotFound() {
		tx.Commit()
		boole = true
		return
	} else {
		tx.Rollback()
		boole = false
		return
	}
}

//只添加购物车商品数量
func (c *Cart) EditCartBy2() bool {

	tx := database.DB.Begin()

	err := tx.Model(&Cart{}).Where("user_id=? and product_id=?", c.UserId, c.Quantity).Update(&Cart{Quantity: c.Quantity, UpdateTime: c.UpdateTime}).Error
	if err == nil {
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		return false
	}
}

//打印购物车信息（用户ID）
func (c *Cart) PrintCart() (info []Cart, boole bool) {

	tx := database.DB.Begin()

	if err := tx.Model(&Cart{}).Where(&Cart{UserId: c.UserId}).Find(&info).Error; err == nil {
		tx.Commit()
		boole = true
		return
	} else {
		tx.Rollback()
		boole = false
		return
	}
}

//删除购物车信息
func (c *Cart) DeleteCart() bool {

	tx := database.DB.Begin()

	if err := tx.Model(&Cart{}).Where(&Cart{ID: c.ID}).Delete(c).Error; err == nil {
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		return false
	}
}

//修改购物车信息
func (c *Cart) EditCartByID() (info Cart, boole bool) {

	tx := database.DB.Begin()

	if err := tx.Model(&Cart{}).Where(&Cart{ID: c.ID}).Update(c).First(&info).Error; err == nil {
		tx.Commit()
		boole = true
		return
	} else {
		tx.Rollback()
		boole = false
		return
	}
}

//全选
func (c *Cart) EditCart() (info []Cart, err error) {

	tx := database.DB.Begin()

	if err = tx.Model(&Cart{}).Where(&Cart{UserId: c.UserId}).Update(c).Count(&info).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

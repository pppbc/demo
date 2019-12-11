package model

import (
	"StoreMS/database"
	"time"
)

//用户表
type User struct {
	ID         int       `json:"id"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	Role       int       `json:"role"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Name       string    `json:"name"`
	Sex        string    `json:"sex"`
	Birthday   string    `json:"birthday"`
	Email      string    `json:"email"`
	Head       string    `json:"head"`
	IdCard     string    `json:"id_card"`
	Token      string    `json:"token,omitempty"`
}

//添加用户
func (u *User) AddUsers() bool {
	var count int
	//开启事务
	tx := database.DB.Begin()
	if err := tx.Model(&User{}).Where("phone=?", u.Phone).Count(&count).Create(u).Error; err == nil && count == 0 {
		//提交
		tx.Commit()
		return true
	} else {
		//回滚
		tx.Rollback()
		return false
	}
}

//验证用户名密码
func (u *User) CheckUser() (err error, info User) {
	//开启事务
	tx := database.DB.Begin()
	if err = tx.Model(&User{}).Where("phone=? and password=?", u.Phone, u.Password).First(&info).Error; err != nil {
		//回滚
		tx.Rollback()
		return
	} else {
		//提交
		tx.Commit()
		return
	}
}

//修改用户信息
func (u *User) EditUser() bool {
	var count int

	//开启事务
	tx := database.DB.Begin()

	err := tx.Model(&User{}).Where(&User{Phone: u.Phone}).Count(&count).Update(u).Error

	if err == nil && count == 1 {
		//提交
		tx.Commit()
		return true
	} else {
		//回滚
		tx.Rollback()
		return false
	}
}

//打印用户信息
func (u *User) PrintUser() (info User, boole bool) {

	tx := database.DB.Begin()

	if err := tx.Model(&User{}).Where(&User{Phone: u.Phone}).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		boole = true
		return
	} else {
		//回滚
		tx.Rollback()
		boole = false
		return
	}
}

//删除用户信息
func (u *User) DeleteUser() bool {
	var count int
	tx := database.DB.Begin()
	if err := tx.Model(&User{}).Where(&User{Phone: u.Phone}).Count(&count).Delete(u).Error; err == nil && count == 1 {
		//提交
		tx.Commit()
		return true
	} else {
		//回滚
		tx.Rollback()
		return false
	}
}

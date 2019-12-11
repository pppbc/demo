package model

import (
	"StoreMS/database"
	"time"
)

type Shipping struct {
	ID               int       `json:"id"`
	UserId           int       `json:"user_id"`
	ReceiverName     string    `json:"receiver_name"`
	ReceiverMobile   string    `json:"receiver_mobile"`
	ReceiverProvince string    `json:"receiver_province"`
	ReceiverCity     string    `json:"receiver_city"`
	ReceiverDistrict string    `json:"receiver_district"`
	ReceiverAddress  string    `json:"receiver_address"`
	ReceiverZip      string    `json:"receiver_zip"`
	CreateTime       time.Time `json:"create_time"`
	UpdateTime       time.Time `json:"update_time"`
}

//添加地址
func (s *Shipping) AddShip() bool {

	tx := database.DB.Begin()

	if err := tx.Model(&Shipping{}).Create(s).Error; err == nil {
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		return false
	}
}

//打印地址信息（用户ID）
func (s *Shipping) PrintShip() (info []Shipping, boole bool) {

	tx := database.DB.Begin()

	if err := tx.Model(&Shipping{}).Where(&Shipping{UserId: s.UserId}).Find(&info).Error; err == nil {
		tx.Commit()
		boole = true
		return
	} else {
		tx.Rollback()
		boole = false
		return
	}
}

//删除地址信息
func (s *Shipping) DeleteShip() bool {

	tx := database.DB.Begin()

	if err := tx.Model(&Shipping{}).Where(&Shipping{ID: s.ID}).Delete(s).Error; err == nil {
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		return false
	}
}

//修改地址信息
func (s *Shipping) EditShipByID() (info Shipping, boole bool) {

	tx := database.DB.Begin()

	if err := tx.Model(&Shipping{}).Where(&Shipping{ID: s.ID}).Update(s).First(&info).Error; err == nil {
		tx.Commit()
		boole = true
		return
	} else {
		tx.Rollback()
		boole = false
		return
	}
}

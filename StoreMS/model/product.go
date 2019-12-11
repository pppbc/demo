package model

import (
	"StoreMS/database"
	"time"
)

type Product struct {
	ID         int       `json:"id"`
	CategoryId int       `json:"category_id"`
	Name       string    `json:"name"`
	Subtitle   string    `json:"subtitle"`
	MainImages string    `json:"main_images"`
	SubImages  string    `json:"sub_images"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

//获取所有商品
func (p *Product) PrintAllProducts() (info []Product, boole bool) {

	//开启事务
	tx := database.DB.Begin()

	if err := tx.Model(&Product{}).Find(&info).Error; err == nil {
		//提交
		tx.Commit()
		boole = true
		return
	} else {
		//回滚
		boole = false
		tx.Rollback()
		return
	}

}

//获取指定商品信息
func (p *Product) GetProducts() (info Product, boole bool) {

	//开启事务
	tx := database.DB.Begin()

	if err := tx.Model(&Product{}).Where(&Product{ID: p.ID}).Find(&info).Error; err == nil {
		//提交
		tx.Commit()
		boole = true
		return
	} else {
		//回滚
		boole = false
		tx.Rollback()
		return
	}
}

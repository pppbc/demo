package model

import (
	"StoreMS/database"
	"time"
	"fmt"
)

//发行积分记录在数据库中
type Publish struct {
	ID           int       `json:"id"`             //主键ID
	PointMakerID int       `json:"point_maker_id"` //积分发行商ID，标识是哪一个积分发行商
	No           int       `json:"no"`             //发行序号，只是标识Execl的序号而已
	Account      string    `json:"account"`        //发行给具体账号
	Balance      int       `json:"balance"`        //发行金额
	Pick         int       `json:"pick"`           //10未领取  20领取
	PublishTime  time.Time `json:"publish_time"`   //发行时间
	ExpireTime   time.Time `json:"expire_time"`    //过期时间
}

//积分发行商
type PointMaker struct {
	ID         int       `json:"id"`     //ID
	Name       string    `json:"name"`   //名称
	Desc       string    `json:"desc"`   //介绍
	Logo       string    `json:"logo"`   //Logo商标图片
	Status     int       `json:"status"` //发行状态 10-待审核 20-发行成功 30-合作终止
	Total      int       `json:"total"`  //单次发行总量
	Fate       float64   `json:"fate"`   //通用积分比值
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

//用户其他积分表
type UserPoint struct {
	ID         int64     `json:"id"`          //ID
	PointID    int       `json:"point_id"`    //PointID
	UserID     int       `json:"user_id"`     //用户ID
	Balance    int       `json:"balance"`     //积分余额
	Total      int       `json:"total"`       //积分总量
	ExpireTime time.Time `json:"expire_time"` //过期时间
	CreateTime time.Time `json:"create_time"` //创建时间
	UpdateTime time.Time `json:"update_time"` //更新时间
}

//获取所有积分运营商
func (pm *PointMaker) Get() (err error, info []PointMaker) {

	tx := database.DB.Begin()

	if err = tx.Model(PointMaker{}).Where(PointMaker{}).Find(&info).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

//根据用户ID获取用户的所有积分
func (up *UserPoint) Get() (err error, info []UserPoint) {
	tx := database.DB.Begin()

	if err = tx.Model(UserPoint{}).Where(UserPoint{UserID: up.UserID}).Find(&info).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

func (pm *PointMaker) GetOne() (err error, info PointMaker) {
	tx := database.DB.Begin()

	if err = tx.Model(&PointMaker{}).Where(PointMaker{ID: pm.ID}).First(&info).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

//根据用户号码获取用户的所有积分
func GetPublish(phone string) (err error, info []Publish) {

	tx := database.DB.Begin()

	if err = tx.Model(Publish{}).Where(Publish{Account: phone,Pick:10}).Find(&info).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

//领取发放的积分
func (p *Publish) PickActive(userid int) (err error) {
	tx := database.DB.Begin()
	var info Publish
	//领取积分
	if err = tx.Model(&Publish{}).Where(Publish{ID: p.ID, Pick: 10}).First(&info).Update(&p).Error; err == nil {
		//写入用户积分表
		if err = tx.Model(&UserPoint{}).Create(&UserPoint{PointID: info.ID, UserID: userid, UpdateTime: time.Now(), CreateTime: time.Now(),ExpireTime:info.ExpireTime}).
			Error; err == nil {
			//同时满足，提交
			tx.Commit()
			return
		}
				fmt.Println("1",err)

		tx.Rollback()
		return
	} else {
				fmt.Println("2",err)

		tx.Rollback()
		return
	}
}

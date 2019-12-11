package model

import (
	"BookMSF/database"
	"log"
	"errors"
	)

type Lend struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Bookname   string `json:"bookname"`
	LendTime   string `json:"lend_time"`
	IsR        int    `json:"is_r"` //1 在借 2 归还
	ReturnTime string `json:"return_time"`
}

func (l *Lend) Append() (info Lend, err error) {
	var count int
	var bo Book
	tx := database.DB.Begin()

	//是否已经存在
	if err = tx.Model(&Lend{}).Where(&Lend{Username: l.Username, Bookname: l.Bookname, IsR: l.IsR}).Count(&count).Error; count == 0 && err == nil {
		//是否有余量
		if err=tx.Model(&Book{}).Where(&Book{Title:l.Bookname}).Exec("update books set num=num-1", ).First(&bo).Error;err==nil&&bo.Num>=0{
				//添加借书记录
			if err = tx.Model(&Lend{}).Create(l).Error; err == nil {
				tx.Commit()
				return
			} else {
				tx.Rollback()
				return
			}
		}else{
			err=errors.New("no")
			tx.Rollback()
			return
		}
	} else {
		tx.Rollback()
		return
	}
}

func (l *Lend) Return() (info Lend, err error) {
	var count int
	tx := database.DB.Begin()

	if err = tx.Model(&Lend{}).Where(&Lend{Username: l.Username, Bookname: l.Bookname, IsR: 1}).Count(&count).Update(l).Error; count == 1 && err == nil {
		log.Println("count=",count)
		tx.Commit()
		return
	} else {
		err=errors.New("no")
		tx.Rollback()
		return
	}
}

func (l *Lend) Getbyuser() (list []Lend, err error) {
	tx := database.DB.Begin()
	if err = tx.Model(&Lend{}).Where(&Lend{Username: l.Username}).Find(&list).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

func (l *Lend) Getbybook() (list []Lend, err error) {
	tx := database.DB.Begin()
	if err = tx.Model(&Lend{}).Where(&Lend{Bookname: l.Bookname}).Find(&list).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

func GetAllL() (list []Lend, err error) {
	tx := database.DB.Begin()
	if err = tx.Model(&Lend{}).Where(&Lend{}).Find(&list).Error; err == nil {
		tx.Commit()
		return
	} else {
		tx.Rollback()
		return
	}
}

//var lendList []*Lend
//
////存借书信息
//func PushLend(l *Lend) {
//	lendList = append(lendList, l)
//	//for _, v := range lendList {
//	//	log.Println(*v)
//	//}
//}
//
////取借书信息(用户名)
//func PullListByUserName(Name string) (info []*Lend, count int) {
//	count = 0
//	for _, v := range lendList {
//		if Name == v.UserName {
//			info = append(info, v)
//			count++
//		}
//	}
//	return
//}
//
////取借书信息(书名)
//func PullListByBookName(Name string) (info []*Lend, count int) {
//	count = 0
//	for _, v := range lendList {
//		if Name == v.BookName {
//			info = append(info, v)
//			count++
//		}
//	}
//	return
//}
//
////判断借书信息,返回id
//func GetListKey(u, b string) int {
//	for k, v := range lendList {
//		if v.BookName == b && v.UserName == u && v.IsR == 1 {
//			return k
//		}
//	}
//	return -1
//}
//
////还书
//func ReturnBook(key int, t string) {
//
//}
//
////
//func ShowList(key int) (l Lend) {
//	l = *lendList[key]
//	return
//}
//
////展示所有
//func ShowAll() (info []Lend) {
//	for _, v := range lendList {
//		info = append(info, *v)
//	}
//	return
//}

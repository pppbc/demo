package model

import "BookMSF/database"

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Num    int     `json:"num"`
	Src    string  `json:"src"`
	Desc   string  `json:"desc"`
}

//添加图书
func (b *Book) Append() (info Book, err error) {
	//开启事物
	tx := database.DB.Begin()

	//添加数据
	if err = tx.Model(&Book{}).Create(b).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//修改图书
func (b *Book) Change() (info Book, err error) {
	//开启事物
	tx := database.DB.Begin()

	//更新数据
	if err = tx.Model(&Book{}).Where(&Book{Title: b.Title}).Update(b).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//删除图书
func (b *Book) Delete() (err error) {
	//开启事物
	tx := database.DB.Begin()

	//删除
	if err = tx.Model(&Book{}).Where(&Book{Title: b.Title, ID: b.ID}).Delete(&Book{}).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//获取一本书籍
func (b *Book) Get() (info Book, err error) {
	//开启事物
	tx := database.DB.Begin()

	//获取数据
	if err = tx.Model(&Book{}).Where(&Book{Title: b.Title}).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//获取书籍列表
func GetAllB() (info []Book, err error) {
	//开启事物
	tx := database.DB.Begin()

	//添加数据
	if err = tx.Model(&Book{}).Where(&Book{}).Find(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//var books []*Book
//
////遍历所有图书
//func AllBook() (info []Book) {
//	for _, v := range books {
//		info = append(info, *v)
//	}
//	return
//}
//
////判断Title是否存在,返回bool+对象
//func JudgeTitleObj(title string) (*Book, bool) {
//	for _, v := range books {
//		if title == v.Title {
//			return v, true
//		}
//	}
//	return nil, false
//}
//
////判断Title是否存在,返回bool+对象
//func JudgeIDKey(id int) (int, bool) {
//	for k, v := range books {
//		if id == v.ID {
//			return k, true
//		}
//	}
//	return -1, false
//}
//
////通过id，修改
//func ChangeBook(b Book) bool {
//	for _, v := range books {
//		if v.ID == b.ID {
//			v.Title = b.Title
//			v.Author = b.Author
//			v.Num = b.Num
//			v.Price = b.Price
//			v.Src = b.Src
//			v.Desc = b.Desc
//
//			return true
//		}
//	}
//	return false
//}
//
////添加书籍
//func AppendBook(u *Book) {
//	books = append(books, u)
//	//for _, v := range books {
//	//	log.Println(v)
//	//}
//}
//
////删除
//func DeleteBookByID(k int) {
//
//	books = append(books[:k], books[k+1:]...)
//}
//
//func CountNum(name string) bool {
//	for _, v := range books {
//		if v.Title == name {
//			if v.Num > 0 {
//				v.Num--
//				return true
//			}
//		}
//	}
//	return false
//}
//func CountNumAdd(name string) {
//	for _, v := range books {
//		if v.Title == name {
//			v.Num++
//		}
//	}
//}

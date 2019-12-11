package model

import (
	"BookMSF/database"
	"BookMSF/redis"
	"strconv"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	HeadImg  string `json:"head_img"`
	Tel      string `json:"tel"`
	IsM      int    `json:"is_m"` //1,普通用户 2,图书管理员 ,3,权限管理员
	Desc     string `json:"desc"`
	Token    string `json:"token,omitempty"`
}

//注册
func (u *User) Register() (err error) {
	var count int
	//开启事务
	tx := database.DB.Begin()
	if err = tx.Model(&User{}).Where("username=?", u.Username).Count(&count).Create(u).Error; err == nil && count == 0 {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//登录验证
func (u *User) LoginCheck() (info User, err error) {
	//开启事务
	tx := database.DB.Begin()
	if err = tx.Model(&User{}).Where(User{Username: u.Username, Password: u.Password}).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//存token信息
func (u *User) LoginSave() (info User, err error) {
	//开启事务
	tx := database.DB.Begin()
	//存数据库
	if err = tx.Model(&User{}).Where(&User{ID: u.ID}).Update(u).First(&info).Error; err == nil {
		//存缓存
		if _, err = redis.DoRedisCommand("hset bookms_token_list token:user_id:"+strconv.Itoa(u.ID)+" "+u.Token, "int"); err == nil {
			log.Println("haahahahhahahhahhahaha,改好了")
			//提交
			tx.Commit()
			return
		} else {
			//回滚
			tx.Rollback()
			return
		}
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//修改信息
func (u *User) Change() (info User, err error) {
	var count int
	//开启事务
	tx := database.DB.Begin()
	if err = tx.Model(&User{}).Where("username=? and id<>?", u.Username, u.ID).Count(&count).Error; count == 0 && err == nil {
		if err = tx.Model(&User{}).Where(&User{Username: u.Username, ID: u.ID}).Update(u).First(&info).Error; err == nil {
			//提交
			tx.Commit()
			return
		} else {
			//回滚
			tx.Rollback()
			return
		}
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//修改权限
func (u *User) ChangeA() (info User, err error) {
	//开启事务
	tx := database.DB.Begin()
	//存数据库
	if err = tx.Model(&User{}).Where(&User{Username: u.Username}).Update(u).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//删除信息
func (u *User) Delete() (err error) {
	//开启事务
	tx := database.DB.Begin()
	//存数据库
	if err = tx.Model(&User{}).Where(&User{Username: u.Username}).Delete(u).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//获取个人信息
func (u *User) Get() (info User, err error) {
	//开启事务
	tx := database.DB.Begin()
	//存数据库
	if err = tx.Model(&User{}).Where(&User{ID: u.ID}).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}
//通过name获取
func (u *User) GetByName() (info User, err error) {
	//开启事务
	tx := database.DB.Begin()
	//存数据库
	if err = tx.Model(&User{}).Where(&User{Username: u.Username}).First(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//获取全部信息
func GetAllU() (info []User, err error) {
	//开启事务
	tx := database.DB.Begin()
	//存数据库
	if err = tx.Model(&User{}).Where(&User{}).Find(&info).Error; err == nil {
		//提交
		tx.Commit()
		return
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//修改密码
func ChangePassword(user User, new string) (err error) {
	var count int
	//开启事务
	tx := database.DB.Begin()
	//存数据库
	if err = tx.Model(&User{}).Where(&User{Username: user.Username, Password: user.Password}).Count(&count).Error; err == nil && count == 1 {
		if err = tx.Model(&User{Username: user.Username}).Update(&User{Password: new}).Error; err == nil {
			//提交
			tx.Commit()
			return
		} else {
			//回滚
			tx.Rollback()
			return
		}
	} else {
		//回滚
		tx.Rollback()
		return
	}
}

//var users []*User
////添加用户
//func Append(u *User) {
//	users = append(users, u)
//}
//
////批量导入
//func Appends(u []*User) bool {
//	users = append(users, u...)
//	return true
//}
//
////删除用户
//func Delete(key int) bool {
//	users = append(users[:key], users[key+1:]...)
//	return true
//}
//
////修改用户信息
//func Change(u User) bool {
//	for _, v := range users {
//		if v.ID == u.ID {
//			v.Name = u.Name
//			v.Desc = u.Desc
//			v.IsM = u.IsM
//			v.Age = u.Age
//			v.Tel = u.Tel
//			return true
//		}
//	}
//	return false
//}
//
////判断除自身外的所有name，有则返回true
//func JudgeNameOther(name string, key int) bool {
//	for k, v := range users {
//		if v.Name == name {
//			if k != key {
//				return true
//			}
//		}
//	}
//	return false
//}
//
////通过id获取下标
//func JudgeId(id int) (int, bool) {
//	for k, v := range users {
//		if v.ID == id {
//			return k, true
//		}
//	}
//	return -1, false
//}
//
////登录
//func JudgeInfo(user User) bool {
//	for _, v := range users {
//		if user.Name == v.Name && user.Password == v.Password {
//			return true
//		}
//	}
//	return false
//}
//
//
////通过ID,返回bool+对象
//func JudgeIDObj(id int) (*User, bool) {
//	for _, v := range users {
//		if id == v.ID {
//			return v, true
//		}
//	}
//	return nil, false
//}
//
////判断name是否存在,返回bool
//func JudgeName(name string) bool {
//	for _, v := range users {
//		if name == v.Name {
//			return true
//		}
//	}
//	return false
//}
//
////判断name是否存在,返回bool+key
//func JudgeNameKey(name string) (int, bool) {
//	for k, v := range users {
//		if name == v.Name {
//			return k, true
//		}
//	}
//	return -1, false
//}
//
////判断name是否存在,返回bool+对象
//func JudgeNameObj(name string) (*User, bool) {
//	for _, v := range users {
//		if name == v.Name {
//			return v, true
//		}
//	}
//	return nil, false
//}
//
////修改权限
//func ChangeA(key, ism int) bool {
//	for k, v := range users {
//		if k == key {
//			v.IsM = ism
//			return true
//		}
//	}
//	return false
//}
//
//func DeleteUserByName(key int) {
//	users = append(users[:key], users[key+1:]...)
//}
//
////遍历所有用户
//func AllUser() (info []User) {
//	for _, v := range users {
//		info = append(info, *v)
//	}
//	return
//}

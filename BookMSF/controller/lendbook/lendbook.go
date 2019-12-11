package lendbook

import (
	logs "BookMSF/log"
	"BookMSF/model"
	"BookMSF/util/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//借书-------------------------------------------------------------------------
func LendBookHandler(r *gin.Context) {
	//获取数据
	BookName := r.PostForm("bookname")
	UserName := r.PostForm("username")
	Time := time.Now().Format("2006-01-02 15:04:05")
	IsR := 1

	//封装
	l := model.Lend{Bookname: BookName, Username: UserName, LendTime: Time, IsR: IsR}

	if info, err := l.Append(); err == nil {
		logs.Loggers.Info("success", zap.String("status", l.Username+"借阅《"+l.Bookname+"》成功"))
		tools.JMOk(r, info, "借书成功！")
	} else {
		tools.JMFail(r, "借书失败")
	}
}

//还书------------------------------------------------------------------------
func ReturnBookHandler(r *gin.Context) {
	//获取数据
	BookName := r.PostForm("bookname")
	UserName := r.PostForm("username")
	Time := time.Now().Format("2006-01-02 15:04:05")
	IsR := 2

	b := &model.Lend{Bookname: BookName, Username: UserName, ReturnTime: Time, IsR: IsR}


	//还书状态改为2，加上还书时间，
	if info, err := b.Return(); err != nil {
		tools.JMFail(r, "没有这条在借记录")
	} else {
		logs.Loggers.Info("success", zap.String("status", b.Username+"归还《"+b.Bookname+"》成功"))
		tools.JMOk(r, info, "还书成功")
	}
}

//通过用户名查找-----------------------------------------------------------------
func GetListByUserName(r *gin.Context) {
	//获取数据
	Username := r.Query("username")

	l := &model.Lend{Username: Username}

	if list, err := l.Getbyuser(); err == nil {
		tools.JMOk(r, list, "find success!")
	} else {
		tools.JMFail(r, "find success!")
	}
}

//通过书名查找------------------------------------------------------------------
func GetListByBookName(r *gin.Context) {

	BookName := r.Query("bookname")

	l := &model.Lend{Bookname: BookName}

	if list, err := l.Getbybook(); err == nil {
		tools.JMOk(r, list, "find success!")
	} else {
		tools.JMFail(r, "find success!")
	}
}

//展示所有借书信息-----------------------------------------------------------------
func ShowBookHandler(r *gin.Context) {
	if list, err := model.GetAllL(); err == nil {
		tools.JMOk(r, list, "获取借书表成功！")
	} else {
		tools.JMFail(r, "获取借书表失败！")
	}
}

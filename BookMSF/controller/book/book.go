package book

import (
	logs "BookMSF/log"
	"BookMSF/model"
	"BookMSF/util/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//修改书籍信息------------------------------------------------------------------------
func ChangeInfoHandler(r *gin.Context) {
	//获取数据
	num, _ := strconv.Atoi(r.PostForm("Num"))
	title := r.PostForm("Title")
	src := r.PostForm("Src")
	price, _ := strconv.ParseFloat(r.PostForm("Price"), 64)
	author := r.PostForm("Author")
	desc := r.PostForm("Desc")

	b := &model.Book{Src: src, Title: title, Author: author, Price: price, Num: num, Desc: desc}

	//修改
	if info, err := b.Change(); err == nil {
		tools.JMOk(r, info, "修改成功!")
	} else {
		tools.JMFail(r, "修改失败!")
	}
}

//通过name寻找书籍------------------------------------------------------------------------
func GetInfoHandler(r *gin.Context) {
	//获取数据
	title := r.Query("title")
	//封装
	book := &model.Book{Title: title}

	//通过title获取该书籍信息

	if info, err := book.Get(); err != nil {
		tools.JMFail(r, "No such Book!")
	} else {
		tools.JMOk(r, info, "Find Success!")
	}
}

//展示所有书籍------------------------------------------------------------------------
func ShowBookAllHandler(r *gin.Context) {
	//获取
	if info, err := model.GetAllB(); err == nil {
		tools.JMOk(r, info, "find success")
	} else {
		tools.JMFail(r, "获取书籍列表失败！")
	}
}

//添加书籍------------------------------------------------------------------------
func AddBookHandler(r *gin.Context) {

	//获取数据
	num, err := strconv.Atoi(r.PostForm("num"))
	if err != nil {
		tools.JMFail(r, "书籍余量转换失败")
		return
	}
	title := r.PostForm("title")
	//src := r.PostForm("src")
	src := "http://img3m4.ddimg.cn/29/28/27926444-1_w_2.jpg"
	price, err := strconv.ParseFloat(r.PostForm("price"), 64)

	if err != nil {
		tools.JMFail(r, "价格转换失败")
		return
	}

	author := r.PostForm("author")
	desc := r.PostForm("desc")

	b := &model.Book{Src: src, Title: title, Author: author, Price: price, Num: num, Desc: desc}

	//添加
	if info, err := b.Append(); err == nil {
		logs.Loggers.Info("success", zap.String("status", "书籍《"+b.Title+"》添加成功!"))
		tools.JMOk(r, info, "书籍《"+b.Title+"》添加成功!")
	} else {
		tools.JMFail(r, "书籍《"+b.Title+"》添加失败!!")
	}
}

//删除书籍------------------------------------------------------------------------
func DeleteBookHandler(r *gin.Context) {
	//获取数据
	id, err := strconv.Atoi(r.Query("id"))
	title := r.Query("title")
	if err != nil {
		tools.JMFail(r, "id 转换失败")
	}
	book := &model.Book{ID: id, Title: title}
	if err := book.Delete(); err == nil {
		logs.Loggers.Info("success", zap.String("status", "书籍《"+book.Title+"》删除成功!"))
		tools.JMOk(r, nil, "书籍《"+book.Title+"》删除成功!")
	} else {
		tools.JMFail(r, "书籍《"+book.Title+"》删除失败!")
	}
}

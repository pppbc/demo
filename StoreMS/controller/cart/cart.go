package cart

import (
	logs "StoreMS/log"
	"StoreMS/model"
	"StoreMS/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

//添加购物车
func AddCartHandler(r *gin.Context) {

	//获取值
	userId, err := strconv.Atoi(r.Query("user_id"))
	if err != nil {
		tools.JMFail(r, "user_id转换失败")
		return
	}
	productId, err := strconv.Atoi(r.Query("product_id"))
	if err != nil {
		tools.JMFail(r, "product_id转换失败")
		return
	}
	quantity, err := strconv.Atoi(r.Query("quantity"))
	if err != nil {
		tools.JMFail(r, "quantity转换失败")
		return
	}
	//封装到结构体Cart
	c := &model.Cart{UserId: userId, Quantity: quantity, ProductId: productId, CreateTime: time.Now(), UpdateTime: time.Now()}

	//验证是否存在
	info, err1 := c.IsIn()

	if err1 {
		//不存在，加入购物车
		if c.AddCart() {
			logs.Loggers.Info("success", zap.String("status", "添加购物车成功"))
			tools.JMOk(r, nil, "add cart succeed")
		} else {
			tools.JMFail(r, "add cart failed")
		}
	} else {
		//存在，数量相加
		c.Quantity = info.Quantity + c.Quantity
		if c.EditCartBy2() {
			tools.JMOk(r, nil, "add quantity succeed")

		} else {
			tools.JMFail(r, "add quantity failed")
		}
	}
}

//显示购物车列表
func PrintCartHandler(r *gin.Context) {
	//获取值
	userId, err := strconv.Atoi(r.Query("user_id"))
	if err != nil {
		tools.JMFail(r, "user_id转换失败")
		return
	}
	//数据封装
	c := &model.Cart{UserId: userId}

	//
	info, err1 := c.PrintCart()
	if err1 == true {
		tools.JMOk(r, info, "print cart succeed")
	} else {
		tools.JMFail(r, "print cart failed")
	}
}

//删除本条购物车记录
func DeleteCartHandler(r *gin.Context) {
	//获取值
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		tools.JMFail(r, "id转换失败")
		return
	}
	//数据封装
	c := &model.Cart{ID: id}

	//删除信息
	if c.DeleteCart() {
		tools.JMOk(r, nil, "delete cart succeed")
	} else {
		tools.JMFail(r, "delete cart failed")
	}
}

//修改购物车记录
func EditCartHandler(r *gin.Context) {
	//获取信息
	id, err := strconv.Atoi(r.PostForm("id"))
	if err != nil {
		tools.JMFail(r, "id转换失败")
		return
	}
	quantity, err := strconv.Atoi(r.PostForm("quantity"))
	if err != nil {
		tools.JMFail(r, "quantity转换失败")
		return
	}
	checked, err := strconv.Atoi(r.PostForm("checked"))
	if err != nil {
		tools.JMFail(r, "checked转换失败")
		return
	}
	//数据封装
	c := &model.Cart{ID: id, Quantity: quantity, Checked: checked, UpdateTime: time.Now()}

	//修改信息
	info, err1 := c.EditCartByID()
	if err1 {
		tools.JMOk(r, info, "edit success")
	} else {
		tools.JMFail(r, "edit failed")
	}
}

//全选
func CheckAllHandler(r *gin.Context) {

	userId, err := strconv.Atoi(r.PostForm("userid"))
	if err != nil {
		tools.JMFail(r, "userid转换失败")
		return
	}
	c := &model.Cart{UserId: userId, Checked: 1, UpdateTime: time.Now()}

	if info, err1 := c.EditCart(); err1 == nil {
		tools.JMOk(r, info, "edit success")
	} else {
		tools.JMFail(r, "edit failed")
	}
}

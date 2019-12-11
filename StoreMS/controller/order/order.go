package order

import (
	"StoreMS/model"
	"StoreMS/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

//添加订单
func AddOrderHandler(r *gin.Context) {
	//获取数据
	info, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		tools.JMFail(r, "read request.body failed")
		return
	}

	//反序列化，接收数据
	var data model.Orders
	err = json.Unmarshal(info, &data)
	if err != nil {
		log.Println(err)

		tools.JMFail(r, "unmarshal failed")
		return
	}
	userId := data.UserId
	shippingId := data.ShippingId
	payment := data.Payment
	postage := data.Postage
	status := data.Status
	orderNo := tools.CreateOrderNo(userId)

	//订单表
	listAll := model.Orders{OrderNo: orderNo, UserId: userId, ShippingId: shippingId, Payment: payment, Postage: postage, Status: status, CreateTime: time.Now(), UpdateTime: time.Now()}

	//订单详情表
	listDetail := data.OrderDetail

	for k, _ := range listDetail {

		//获取信息
		listDetail[k].OrderNo = orderNo
		listDetail[k].UserId = userId
		listDetail[k].CreateTime = time.Now()
		listDetail[k].UpdateTime = time.Now()
	}

	if model.AddOrder(listAll, listDetail) {
		tools.JMOk(r, nil, "success")
		return
	} else {
		tools.JMFail(r, "faield")
		return
	}
}

//删除指定订单
func DeleteOrderHandler(r *gin.Context) {
	orderNo, err := strconv.ParseInt(r.Query("orderNo"), 10, 64)

	if err != nil {
		tools.JMFail(r, "orderNo转换失败")
		return
	}

	if model.DeleteOrder(orderNo) {
		tools.JMOk(r, nil, "delete order success")
		return
	} else {
		tools.JMFail(r, "delete order failed")
		return
	}
}

//获取所有订单信息
func PrintOrderAllHandler(r *gin.Context) {
	userId, err := strconv.Atoi(r.Query("userid"))

	if err != nil {
		tools.JMFail(r, "userid转换失败")
		return
	}

	o := model.Orders{UserId: userId}

	info, err1 := o.PrintOrder()
	if err1 == true {
		tools.JMOk(r, info, "print order success")
		return
	} else {
		tools.JMFail(r, "print order failed")
		return
	}
}

//支付订单
func PayHandler(r *gin.Context) {
	orderNo, err := strconv.ParseInt(r.Query("orderNo"), 10, 64)

	info := &model.Orders{OrderNo: orderNo, UpdateTime: time.Now(), Status: 11}
	if err != nil {
		tools.JMFail(r, "orderNo转换失败")
		return
	}

	if model.Pay(info) {
		tools.JMOk(r, nil, "订单支付成功")
	} else {
		tools.JMFail(r, "订单支付失败")
	}
}

//获取指定订单
func GetOrderByStatusHandler(r *gin.Context) {
	userId, err := strconv.Atoi(r.Query("userid"))
	if err != nil {
		tools.JMFail(r, "userid转换失败")
		return
	}
	status, err := strconv.Atoi(r.Query("userid"))
	if err != nil {
		tools.JMFail(r, "status转换失败")
		return
	}
	order := &model.Orders{UserId: userId, Status: status}

	if err, info := order.GetOrder(); err == nil {
		tools.JMOk(r, info, "success")
	} else {
		tools.JMFail(r, "failed")
	}
}

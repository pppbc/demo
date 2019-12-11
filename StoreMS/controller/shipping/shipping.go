package shipping

import (
	"StoreMS/model"
	"StoreMS/tools"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

//添加地址信息
func AddShipHandler(r *gin.Context) {

	//获取值
	userId, err := strconv.Atoi(r.Query("user_id"))
	if err != nil {
		tools.JMFail(r, "user_id转换失败")
		return
	}
	receiverName := r.Query("receiver_name")
	receiverMobile := r.Query("receiver_mobile")
	receiverProvince := r.Query("receiver_province")
	receiverCity := r.Query("receiver_city")
	receiverDistrict := r.Query("receiver_district")
	receiverAddress := r.Query("receiver_address")
	receiverZip := r.Query("receiver_zip")

	//验证数据格式
	if !tools.JudgePhone(receiverMobile) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}

	//封装到结构体
	s := &model.Shipping{UserId: userId, ReceiverMobile: receiverMobile, ReceiverProvince: receiverProvince, ReceiverCity: receiverCity, ReceiverDistrict: receiverDistrict, ReceiverAddress: receiverAddress, ReceiverZip: receiverZip, ReceiverName: receiverName, CreateTime: time.Now(), UpdateTime: time.Now()}

	//添加
	if s.AddShip() {
		tools.JMOk(r, nil, "add ship succeed")
	} else {
		tools.JMFail(r, "add ship failed")
	}
}

//显示地址
func PrintShipHandler(r *gin.Context) {
	//获取值
	userId, err := strconv.Atoi(r.Query("user_id"))
	if err != nil {
		tools.JMFail(r, "user_id转换失败")
		return
	}
	//数据封装
	s := &model.Shipping{UserId: userId}

	//
	info, err1 := s.PrintShip()
	if err1 == true {
		tools.JMOk(r, info, "print shipping succeed")
	} else {
		tools.JMFail(r, "print shipping failed")
	}
}

//删除选中地址信息
func DeleteShipHandler(r *gin.Context) {
	//获取值
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		tools.JMFail(r, "id转换失败")
		return
	}
	//数据封装
	s := &model.Shipping{ID: id}

	//删除信息
	if s.DeleteShip() {
		tools.JMOk(r, nil, "delete shipping succeed")
	} else {
		tools.JMFail(r, "delete shipping failed")
	}
}

//修改选中地址信息
func EditShipHandler(r *gin.Context) {
	//获取信息
	id, err := strconv.Atoi(r.PostForm("id"))
	if err != nil {
		tools.JMFail(r, "id转换失败")
		return
	}
	receiverName := r.PostForm("receiver_name")
	receiverMobile := r.PostForm("receiver_mobile")
	receiverProvince := r.PostForm("receiver_province")
	receiverCity := r.PostForm("receiver_city")
	receiverDistrict := r.PostForm("receiver_district")
	receiverAddress := r.PostForm("receiver_address")
	receiverZip := r.PostForm("receiver_zip")

	if !tools.JudgePhone(receiverMobile) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}

	//封装到结构体
	s := &model.Shipping{ID: id, ReceiverMobile: receiverMobile, ReceiverProvince: receiverProvince, ReceiverCity: receiverCity, ReceiverDistrict: receiverDistrict, ReceiverAddress: receiverAddress, ReceiverZip: receiverZip, ReceiverName: receiverName, UpdateTime: time.Now()}

	//修改信息
	info, err1 := s.EditShipByID()
	if err1 {
		tools.JMOk(r, info, "edit shipping success")
	} else {
		tools.JMFail(r, "edit shipping failed")
	}
}

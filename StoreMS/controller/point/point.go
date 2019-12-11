package point

import (
	logs "StoreMS/log"
	"StoreMS/model"
	"StoreMS/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//获取指定用户积分信息
func GetUserPointHandler(r *gin.Context) {
	userId, err := strconv.Atoi(r.Query("user_id"))
	if err != nil {
		tools.JMFail(r, "参数解析失败！")
		return
	}
	up := model.UserPoint{UserID: userId}

	if err, info := up.Get(); err == nil {
		tools.JMOk(r, info, "get userpoint success!")
	} else {
		tools.JMFail(r, "get userpoint failed!")
	}
}

//获取所有积分运营商
func GetMakerHandler(r *gin.Context) {

	pm := model.PointMaker{}

	err, info := pm.Get()

	if err == nil {
		logs.Loggers.Info("success", zap.String("status", "获取商户成功"))
		tools.JMOk(r, info, "get pointmaker success!")
	} else {
		logs.Loggers.Error("error")
		tools.JMFail(r, "get pointmaker failed!")
	}
}

//获取指定id运行商信息
func GetMakerInfoHandler(r *gin.Context) {
	pointId, err := strconv.Atoi(r.Query("user_id"))
	if err != nil {
		tools.JMFail(r, "参数解析失败！")
		return
	}
	pm := model.PointMaker{ID: pointId}

	if err, info := pm.GetOne(); err == nil {
		tools.JMOk(r, info, "get pointmaker success!")
	} else {
		tools.JMFail(r, "get pointmaker failed!")
	}
}

//获取用户所有积分信息（包括领取和未领取）
func GetPointPublishHandler(r *gin.Context) {

	phone := r.Query("phone")
	if !tools.JudgePhone(phone) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}

	if err, info := model.GetPublish(phone); err == nil {
		tools.JMOk(r, info, "get pointpublish success!")
	} else {
		tools.JMFail(r, "get pointpublish failed!")
	}
}

//积分领取
func PickActiveHandler(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		tools.JMFail(r, "参数解析失败！")
		return
	}

	userid, err := strconv.Atoi(r.Query("user_id"))
	if err != nil {
		tools.JMFail(r, "参数解析失败！")
		return
	}

	p := &model.Publish{ID: id, Pick: 20}

	if p.PickActive(userid) == nil {
		tools.JMOk(r, nil, "change  success!")
	} else {
		tools.JMFail(r, "change  failed!")
	}
}

package user

import (
	"StoreMS/config"
	logs "StoreMS/log"
	"StoreMS/model"
	"StoreMS/redis"
	"StoreMS/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
	"fmt"
)

//登录
func LoginHandler(r *gin.Context) {

	//获取请求中的数据
	phone := r.Query("phone")
	password := r.Query("password")
	udid:=r.Query("udid")
	

fmt.Println(phone)
fmt.Println(password)


	//验证数据格式
	if !tools.JudgePhone(phone) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}
	if !tools.JudgePassword(password) {
		tools.JMFail(r, "bad passwordFormat！")
		return
	}

	//封装到结构体里
	u := &model.User{Phone: phone, Password: tools.Md5Encode(password)}

	//调用方法，进行登录验证
	if err, info := u.CheckUser(); err == nil {

		//成功则生成token
		str, err1 := tools.CreateToken([]byte(tools.SecretKey), "xxx", info.ID, udid)

		//存token
		u = &model.User{Phone: phone, Password: tools.Md5Encode(password), Token: str,UpdateTime:time.Now()}

		//存数据库
		u.EditUser()

		//存redis缓存 
		_, err := redis.DoRedisCommand("hset token_list token:user_id:"+strconv.Itoa(info.ID)+" "+str, "int64")
		if err != nil {
			logs.Loggers.Error("error", zap.String("status", "登录失败，token存缓存出错"))
			tools.JMFail(r, "token failed1！")
			return
		}

		//做出响应
		if err1 == nil {
			//写日志
			logs.Loggers.Info("success", zap.String("status", "登录成功"))
			//返回数据（token）
			tools.JMOk(r, str, "login succeed！")
		} else {
			fmt.Println(err1)
			logs.Loggers.Error("error", zap.String("status", "登录失败，token出错"))
			tools.JMFail(r, "token failed2！")
		}
	} else {
		//提示失败
		tools.JMFail(r, "login failed！")
	}
}

//注册
func RegisterHandler(r *gin.Context) {

	//获取请求中的数据
	phone := r.Query("phone")
	password := r.Query("password")
	creatTime := time.Now()
	updateTime := time.Now()

	//验证数据格式
	if !tools.JudgePhone(phone) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}
	if !tools.JudgePassword(password) {
		tools.JMFail(r, "bad passwordFormat！")
		return
	}

	//封装到结构体里
	u := &model.User{Phone: phone, Password: tools.Md5Encode(password), CreateTime: creatTime, UpdateTime: updateTime}

	//调用方法，进行注册操作
	if u.AddUsers() {
		logs.Loggers.Info("success", zap.String("status", "注册成功"))
		tools.JMOk(r, nil, "register succeed！")
	} else {
		tools.JMFail(r, "register failed！")
	}
}

//打印信息
func PrintUserHandler(r *gin.Context) {
	phone := r.Query("phone")
	//验证数据格式
	if !tools.JudgePhone(phone) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}
	u := &model.User{Phone: phone}
	info, err := u.PrintUser()

	if err == true {
		logs.Loggers.Info("success", zap.String("status", "获取用户信息成功"))
		tools.JMOk(r, info, "print succeed！")
	} else {
		tools.JMFail(r, "print failed！")
	}
}

//修改user信息
func UpdateUserHandler(r *gin.Context) {

	//获取请求中的数据
	phone := r.PostForm("phone")
	password := r.PostForm("password")
	role, err := strconv.Atoi(r.PostForm("role"))
	if err != nil {
		tools.JMFail(r, "role转换失败")
		return
	}
	name := r.PostForm("name")
	sex := r.PostForm("sex")
	birthday := r.PostForm("birthday")
	email := r.PostForm("email")
	updateTime := time.Now()

	//验证数据格式
	if !tools.JudgePhone(phone) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}
	if !tools.JudgePassword(password) {
		tools.JMFail(r, "bad passwordFormat！")
		return
	}
	if !tools.JudgeEmail(email) {
		tools.JMFail(r, "bad emailFormat！")
		return
	}

	//封装到结构体里
	u := &model.User{Phone: phone, Password: tools.Md5Encode(password), Role: role, UpdateTime: updateTime, Name: name, Sex: sex, Birthday: birthday, Email: email}

	if u.EditUser() {
		logs.Loggers.Info("success", zap.String("status", "更新用户信息成功"))
		tools.JMOk(r, nil, "update succeed！")
	} else {
		tools.JMFail(r, "update failed！")
	}
}

//删除用户信息
func DeleteUserHandler(r *gin.Context) {
	phone := r.Query("phone")
	if !tools.JudgePhone(phone) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}
	u := &model.User{Phone: phone}
	if u.DeleteUser() {
		logs.Loggers.Info("success", zap.String("status", "注销用户成功"))
		tools.JMOk(r, nil, "delete succeed！")
	} else {
		tools.JMFail(r, "delete failed！")
	}
}

//上传头像
func UploadHeadHandler(r *gin.Context) {
	file, _ := r.FormFile("file")
	phone := r.PostForm("phone")

	filename := phone + ".png"

	if err := r.SaveUploadedFile(file, config.GetLocalPath()+filename); err != nil {
		tools.JMFail(r, "upload headImg failed")
		return
	}

	u := &model.User{Phone: phone, Head: config.GetFileServer() + filename}

	if u.EditUser() {
		logs.Loggers.Info("success", zap.String("status", "头像上传成功"))
		tools.JMOk(r, nil, "update headImg success")
	} else {
		tools.JMFail(r, "update headImg failed")
	}
}

//修改密码
func ChangePasswordHandler(r *gin.Context) {

	phone := r.Query("phone")
	password := r.Query("password")
	newPassword := r.Query("new_password")

	if !tools.JudgePhone(phone) {
		tools.JMFail(r, "bad phoneFormat！")
		return
	}
	if !tools.JudgePassword(password) {
		tools.JMFail(r, "bad passwordFormat！")
		return
	}
	if !tools.JudgePassword(newPassword) {
		tools.JMFail(r, "bad passwordFormat！")
		return
	}

	u := model.User{Phone: phone, Password: password}
	if err, _ := u.CheckUser(); err == nil {
		u = model.User{Phone: phone, Password: newPassword}
		if u.EditUser() {
			logs.Loggers.Info("success", zap.String("status", "修改密码成功"))
			tools.JMOk(r, nil, "change password success")
		} else {
			tools.JMFail(r, "change password failed")
		}
	} else {
		tools.JMFail(r, "change password failed")
	}
}

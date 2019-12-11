package user

import (
	logs "BookMSF/log"
	"BookMSF/model"
	"BookMSF/util/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"log"
)

//登录----------------------------------------------------------------------
func LoginHandler(r *gin.Context) {
	//获取Form中的数据
	Username := r.Query("username")
	Password := r.Query("password")
	UDid := r.Query("udid")


	//加密
	Password = tools.Md5Encode(Password)

	//数据封装
	u := &model.User{Username: Username, Password: Password}

	if info, err := u.LoginCheck(); err == nil {
		//生成token
		if str, err := tools.CreateToken([]byte(tools.SecretKey), "", info.ID, UDid); err == nil {
			//写数据库,写缓存
			u1 := &model.User{ID:info.ID,Username:info.Username,Password:info.Password,Token:str}
			if info, err := u1.LoginSave(); err == nil {
				//成功
				logs.Loggers.Info("success", zap.String("status", info.Username+"登录成功"))
				tools.JMOk(r, info, "登录成功！")
			} else {
				//提示失败
				tools.JMFail(r, "登录失败！存数据库失败！")
			}
		} else {
			//提示失败
			tools.JMFail(r, "登录失败！生成token失败！")
		}
	} else {
		//提示失败
		log.Println("111")
		tools.JMFail(r, "登录失败！请检查用户名或密码！")
	}
}

//注册----------------------------------------------------------------------
func RegisterHandler(r *gin.Context) {
	//获取Form中的数据
	Username := r.Query("username")
	Password := r.Query("password")
	IsM:=1
	//检查格式
	if !tools.JudgeNotNull(Username) {
		tools.JMFail(r, "用户名不能为空！")
		return
	}
	if !tools.JudgePassword(Password) {
		tools.JMFail(r, "密码格式有误，请确认密码长度介于6-18位之间！")
		return
	}
	
	//加密
	Password = tools.Md5Encode(Password)
	//封装
	u := &model.User{Username: Username, Password: Password,IsM:IsM}

	//判断用户信息是否已经存在数据库，不存在则注册
	if err := u.Register(); err == nil {
		logs.Loggers.Info("success", zap.String("status", u.Username+"注册成功"))
		tools.JMOk(r, nil, "注册成功！")

	} else {
		tools.JMFail(r, "注册失败！")
	}
}

//修改密码----------------------------------------------------------------------
func ChangePassword(r *gin.Context) {

	//获取Form中的数据
	name := r.PostForm("username")
	password := r.PostForm("password")
	newpassword1 := r.PostForm("new_password1")
	newpassword2 := r.PostForm("new_password2")

	if newpassword1 != newpassword2 {
		tools.JMFail(r, "两次密码不一致！")
	}

	u := model.User{Username: name, Password: password}

	//原密码判断,成功则修改
	if err := model.ChangePassword(u, newpassword1); err == nil {
		tools.JMOk(r, u, "密码修改成功！")
	} else {
		tools.JMFail(r, "密码修改失败！")
	}
}

//修改个人信息----------------------------------------------------------------------
func ChangeInfo(r *gin.Context) {
	//获取request Body的数据
	ID, err := strconv.Atoi(r.PostForm("id"))
	if err != nil {
		tools.JMFail(r, "id 转换失败！")
		return
	}
	Name := r.PostForm("username")
	Age, err := strconv.Atoi(r.PostForm("age"))
	if err != nil {
		tools.JMFail(r, "age 转换失败！")
		return
	}
	Tel:=r.PostForm("tel")
	
	if !tools.JudgePhone(Tel) {
		tools.JMFail(r, "手机号格式有误，请确认是否输入错误！")
		return
	}
	Desc := r.PostForm("desc")

	//封装
	u := model.User{ID: ID, Username: Name, Age: Age, Tel: Tel, Desc: Desc}

	//修改
	if info, err := u.Change(); err == nil {
		//判断用户名是否冲突,不冲突则修改
		tools.JMOk(r, info, "个人信息修改成功！")
	} else {
		//冲突或失败返回原数据
		tools.JMFail(r, "修改失败，注意用户名不能重复!")
	}
}


//通过姓名获取

func GetByNameHandler(r *gin.Context) {
	//获取form的数据
	Username:= r.Query("username")
	
	if !tools.JudgeNotNull(Username) {
		tools.JMFail(r, "用户名不能为空！")
		return
	}
	//封装
	u := &model.User{Username: Username}

	//获取该用户信息
	if info, err := u.GetByName(); err == nil {
		tools.JMOk(r, info, "获取成功！")
	} else {
		tools.JMFail(r, "获取失败!")
	}
}

//获取个人信息------------------------------------------------------------------
func GetMyselfHandler(r *gin.Context) {
	//获取form的数据
	ID, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		tools.JMFail(r, "id 转换失败！")
		return
	}
	//封装
	u := &model.User{ID: ID}

	//获取该用户信息
	if info, err := u.Get(); err == nil {
		tools.JMOk(r, info, "获取成功！")
	} else {
		tools.JMFail(r, "获取失败!")
	}
}

//修改权限---------------------------------------------------------
func ChangeAHandler(r *gin.Context) {
	//获取值
	IsM, err := strconv.Atoi(r.PostForm("is_m"))
	if err != nil {
		tools.JMFail(r, "权限 转换失败！")
		return
	}
	Username:=r.PostForm("username")
	
	if !tools.JudgeNotNull(Username) {
		tools.JMFail(r, "用户名不能为空！")
		return
	}
	
	//封装
	u := &model.User{Username: Username, IsM: IsM}

	//修改权限
	if info, err := u.ChangeA(); err == nil {
		tools.JMOk(r, info, "修改权限成功！")
	} else {
		tools.JMFail(r, "修改权限失败!")
	}
}

//删除用户---------------------------------------------------------
func DeleteUserHandler(r *gin.Context) {
	
	Username := r.Query("username")

	log.Println(".........",Username)
	
	if !tools.JudgeNotNull(Username) {
		tools.JMFail(r, "用户名不能为空！")
		return
	}
	u := &model.User{Username: Username}

	if err := u.Delete(); err == nil {
		logs.Loggers.Info("success", zap.String("status", u.Username+"注销成功"))
		tools.JMOk(r, nil, "删除成功!")
	} else {
		tools.JMFail(r, "没有该用户!")
	}
}

//展示所有用户信息-------------------------------------------------
func ShowUserHandler(c *gin.Context) {
	//获取信息
	if info, err := model.GetAllU(); err == nil {
		tools.JMOk(c, info, "获取所有用户信息成功")
	} else {
		tools.JMFail(c, "拉取用户信息失败")
	}
}

//上传头像
// func UploadHeadHandler(r *gin.Context) {
// 	file, _ := r.FormFile("file")
// 	username := r.PostForm("username")
// 	filename := username + ".png"

// 	if err := r.SaveUploadedFile(file, config.GetLocalPath()+filename); err != nil {
// 		tools.JMFail(r, "upload headImg failed")
// 		return
// 	}

// 	u := &model.User{Phone: phone, Head: config.GetFileServer() + filename}

// 	if u.EditUser() {
// 		logs.Loggers.Info("success", zap.String("status", "头像上传成功"))
// 		tools.JMOk(r, nil, "update headImg success")
// 	} else {
// 		tools.JMFail(r, "update headImg failed")
// 	}
// }

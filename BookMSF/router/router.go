package router

import (
	"BookMSF/config"
	"BookMSF/controller/book"
	"BookMSF/controller/lendbook"
	"BookMSF/controller/user"
	"BookMSF/util/tools"
	"github.com/gin-gonic/gin"
)

//路由包，分发各种请求到不同包执行

func InitRouter() {

	//init router
	router := gin.Default()

	//解决跨域问题
	router.Use(tools.MiddleWare())

	//创建接口
	//用户
	use := router.Group("/user")
	{
		use.GET("/login", user.LoginHandler)
		use.GET("/register", user.RegisterHandler)
		//验证token
		use.Use(tools.ValidateToken())
		use.POST("/changepassword", user.ChangePassword)
		use.POST("/change", user.ChangeInfo)
		use.GET("/get", user.GetMyselfHandler)
		use.GET("/getname", user.GetByNameHandler)
		use.POST("/changea", user.ChangeAHandler)
		use.GET("/delete", user.DeleteUserHandler)
		use.GET("/getall", user.ShowUserHandler)
	}

	//图书
	booker := router.Group("/book")
	{
		booker.GET("/getall", book.ShowBookAllHandler)
		booker.Use(tools.ValidateToken())
		booker.GET("/get", book.GetInfoHandler)
		booker.POST("/change", book.ChangeInfoHandler)
		booker.POST("/add", book.AddBookHandler)
		booker.GET("/delete", book.DeleteBookHandler)
	}

	//借书还书
	list := router.Group("/lend")
	{
		list.Use(tools.ValidateToken())
		list.POST("/lend", lendbook.LendBookHandler)
		list.GET("/getbybook", lendbook.GetListByBookName)
		list.GET("/getbyuser", lendbook.GetListByUserName)
		list.POST("/return", lendbook.ReturnBookHandler)
		list.GET("/getall", lendbook.ShowBookHandler)
	}
	//监听端口
	_ = router.Run(":" + config.GetPort())
}

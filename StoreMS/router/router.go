package router

import (
	"StoreMS/controller/cart"
	"StoreMS/controller/order"
	"StoreMS/controller/point"
	"StoreMS/controller/product"
	"StoreMS/controller/shipping"
	"StoreMS/controller/user"
	tools2 "StoreMS/tools"
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() {
	//默认路由
	router := gin.Default()

	//解决跨域问题
	router.Use(tools2.MiddleWare())

	//封装接口
	users := router.Group("/user")
	{
		users.GET("/login", user.LoginHandler)
		users.GET("/register", user.RegisterHandler)
		users.Use(tools2.ValidateToken())
		users.POST("/update", user.UpdateUserHandler)
		users.GET("/get", user.PrintUserHandler)
		users.GET("/delete", user.DeleteUserHandler)
		users.POST("/updateHeader", user.UploadHeadHandler)
		users.POST("/getcode", user.PostMessage)
		users.GET("/updatePassword", user.ChangePasswordHandler)
	}
	//token
	router.Use(tools2.ValidateToken())
	products := router.Group("/product")
	{
		products.GET("/get", product.GetProductsHandler)
		//router.POST("/upload", product.UploadHandler)
	}
	cars := router.Group("/cart")
	{
		cars.GET("/add", cart.AddCartHandler)
		cars.GET("/get", cart.PrintCartHandler)
		cars.GET("/delete", cart.DeleteCartHandler)
		cars.POST("/update", cart.EditCartHandler)
		cars.GET("/checkAll", cart.CheckAllHandler)
	}
	ships := router.Group("/ship")
	{
		ships.GET("/add", shipping.AddShipHandler)
		ships.GET("/get", shipping.PrintShipHandler)
		ships.GET("/delete", shipping.DeleteShipHandler)
		ships.POST("/update", shipping.EditShipHandler)
	}
	orders := router.Group("/order")
	{
		orders.POST("/add", order.AddOrderHandler)
		orders.GET("/delete", order.DeleteOrderHandler)
		orders.GET("/get", order.PrintOrderAllHandler)
		orders.GET("/pay", order.PayHandler)
		orders.GET("/getCheck", order.GetOrderByStatusHandler)
	}
	pointers := router.Group("/point")
	{
		pointers.GET("/getUserpoint", point.GetUserPointHandler)
		pointers.POST("/getMaker", point.GetMakerHandler)
		pointers.GET("/getMakerone", point.GetMakerInfoHandler)
		pointers.GET("/getPointpublish", point.GetPointPublishHandler)
		pointers.GET("/change", point.PickActiveHandler)
	}
	//监听
	_ = router.Run(":8088")
}

package main

import (
	"StoreMS/config"
	"StoreMS/database"
	"StoreMS/router"
)

func main() {
	//初始化log

	//初始化配置信息
	config.InitConfig()

	//数据库连接
	database.ConnectDB()

	//初始化路由
	router.InitRouter()

	//数据库断开连接
	defer database.DisconnectDB()

}

package main

import (
	"BookMSF/config"
	"BookMSF/database"
	"BookMSF/router"
)

func main() {
	//初始化log

	//初始化配置文件
	config.InitConfig()

	//初始化数据库
	database.InitDB()

	//初始化路由
	router.InitRouter()

	defer database.CloseDB()
}

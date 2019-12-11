package database

import (
	"BookMSF/config"
	logs "BookMSF/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

//数据库

var DB *gorm.DB

func InitDB() {
	var err error

	//获取数据库配置信息
	mysqlSetting := config.GetSqlSetting()
	conn := mysqlSetting.Username + ":" + mysqlSetting.Password + "@(" + mysqlSetting.IP + ")/" + mysqlSetting.DataName + "?charset=utf8&parseTime=true&loc=Local"

	//开启数据库
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	} else {
		logs.Loggers.Info("数据库连接成功")
	}

	//最大空闲数
	DB.DB().SetMaxIdleConns(50)
	//最大连接数
	DB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	//关闭数据库
	err := DB.Close()
	if err != nil {
		logs.Loggers.Error("error", zap.String("status", "关闭数据库失败"))
	}
}

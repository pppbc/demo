package database

import (
	"StoreMS/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	//获取数据库配置信息
	mysqlSetting := config.GetSqlSetting()
	conn := mysqlSetting.Username + ":" + mysqlSetting.Password + "@(" + mysqlSetting.IP + ")/" + mysqlSetting.DataName + "?charset=utf8&parseTime=true&loc=Local"

	//开启数据库
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	} else {
		log.Println("success init database")
	}

	//最大空闲数
	DB.DB().SetMaxIdleConns(50)
	//最大连接数
	DB.DB().SetMaxOpenConns(100)
}

func DisconnectDB() {

	//关闭数据库
	err := DB.Close()
	if err != nil {
		log.Println(err)
	}
}

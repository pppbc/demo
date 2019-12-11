package config

//配置信息
import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Config *goconfig.ConfigFile

type MySqlSetting struct {
	Username string
	Password string
	IP       string
	DataName string
}

func InitConfig() {
	cfg, err := goconfig.LoadConfigFile("config/config.ini")
	if err != nil {
		log.Println("Read failed[config.ini]")
		return
	}
	log.Println("success init config")
	Config = cfg
}
func GetPort() string {
	value, err := Config.GetValue("dev", "port")

	if err != nil {
		log.Printf("can't get value(%s):%s", value, err)
	}
	return value
}

func GetSqlSetting() *MySqlSetting {
	sec, _ := Config.GetSection("mysql")
	myconfig := &MySqlSetting{}
	myconfig.Username = sec["name"]
	myconfig.Password = sec["password"]
	myconfig.IP = sec["ip"]
	myconfig.DataName = sec["database"]
	return myconfig
}

func GetFileServer() string {
	value, err := Config.GetValue("file", "FileServer")
	if err != nil {
		log.Println("can't get value(%s):%s", value, err)
	}
	return value
}
func GetLocalPath() string {
	value, err := Config.GetValue("file", "LocalPath")
	if err != nil {
		log.Println("can't get value(%s):%s", value, err)
	}
	return value
}
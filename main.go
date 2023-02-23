package main

import (
	"loginsystem/Log"
	"loginsystem/conf"
	"loginsystem/models"
	"loginsystem/routers"
)

func main() {
	conf.InitConfig()               //初始化配置文件
	Log.ErrorLog = Log.InitErrLog() //初始化Err日志系统
	Log.Info = Log.InitLogLog()     //初始化info日志系统
	Log.Info.Printf("Config配置完成, 日志系统配置完成!")
	models.DB = models.InitDB() //初始化DB
	defer models.CloseDB(models.DB)
	routers.RunRouters() //初始静态资源 设置路由
}

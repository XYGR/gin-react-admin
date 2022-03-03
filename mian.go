package main

import (
	"gin-react-admin/core"
	"gin-react-admin/global"
	"gin-react-admin/initialize"
	"go.uber.org/zap"
)

//go:generate go version

func main() {

	global.GRA_VP = core.Viper() // 初始化Viper
	global.GRA_LOG = core.Zap()  //初始化zap日志库
	zap.ReplaceGlobals(global.GRA_LOG)
	global.GRA_DB = initialize.Gorm() // gorm链接数据库
	initialize.Timer()
	//	todo DBList
	//initialize.DBList()
	if global.GRA_DB != nil {
		initialize.RegisterTables(global.GRA_DB) // 初始化表
		//	程序结束前关闭数据库链接
		db, _ := global.GRA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}

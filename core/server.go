package core

import (
	"gin-react-admin/global"
	"gin-react-admin/service/system"
)

type server interface {
	ListenAdnServe() error
}

func RunWindowsServer() {
	if global.GRA_CONFIG.System.UseMultipoint || global.GRA_CONFIG.System.UseRedis {
		// 初始化redis服务
		// todo 初始化redis服务
		//initialize.Redis()
	}

	//	从db加载JWT数据
	if global.GRA_DB != nil {
		system.LoadAll()
	}
	// 初始化路由
	Router := initalize
}

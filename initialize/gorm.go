package initialize

import (
	"gin-react-admin/global"
	"gin-react-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	// todo 兼容pgsql
	return GormMysql()

}

//@function: RegisterTables
//@description: 注册数据库表专用

func RegisterTables(db *gorm.DB) {
	// todo 添加其他表
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
	)
	if err != nil {
		global.GRA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GRA_LOG.Info("register table success")
}

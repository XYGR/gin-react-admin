package system

import (
	"gin-react-admin/global"
	"gin-react-admin/model/system"
	"go.uber.org/zap"
)

type JwtService struct{}

func LoadAll() {
	var data []string
	err := global.GRA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GRA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}

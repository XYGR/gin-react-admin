package global

import (
	"gin-react-admin/config"
	"gin-react-admin/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GRA_VP     *viper.Viper
	GRA_CONFIG config.Server
	GRA_LOG    *zap.Logger
	GRA_DB     *gorm.DB
	GRA_DBList map[string]*gorm.DB
	GRA_Timer  timer.Timer = timer.NewTimerTask()
	BlackCache local_cache.Cache
)

package initialize

import (
	"gin-react-admin/global"
	"gin-react-admin/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

//@function: Routers
//@description: 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.StaticFS(global.GRA_CONFIG.Local.Path, http.Dir(global.GRA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址

	global.GRA_LOG.Info("use middleware logger")

	global.GRA_LOG.Info("use middleware cors")
	//注册swagger路由
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GRA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System
	//	todo 添加其他路由组
	// 创建公开路由组 不需要鉴权
	PublicGroup := Router.Group("")
	{
		// 健康/心跳检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	}

	//	创建私有路由组 需要鉴权
	PrivateGroup := Router.Group("")
	//	注册鉴权中间件
	PrivateGroup.Use()
	{

	}
}

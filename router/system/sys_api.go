package system

import "github.com/gin-gonic/gin"

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.Op) // 注册路由组和中间件
}

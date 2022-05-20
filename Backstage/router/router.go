package router

import (
	"dora/Backstage/controller"
	"dora/Backstage/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST("/api/register", controller.Register)

	r.POST("/api/login", controller.Login)

	//用户信息
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.AuthInfo)

	return r
}

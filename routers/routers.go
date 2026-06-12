// 管理程序的路由信息
package routers

import (
	"scaffold-demo/middlewares"
	"scaffold-demo/routers/auth"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")

	// 公开路由：登录、退出
	auth.RegisterSubRouters(apiGroup)

	// 需要 JWT 认证的路由
	protected := apiGroup.Group("")
	protected.Use(middlewares.JWTAuth)
	// 后续受保护路由在这里注册
}

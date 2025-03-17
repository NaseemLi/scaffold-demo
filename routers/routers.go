// 管理程序的路由信息
package routers

import (
	"scaffold-demo/routers/auth"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegisterRouters(r *gin.Engine) {
	//登陆的路由配置
	//1. 登陆login
	//2. 退出logout
	//3. api/auth/login
	apiGroup := r.Group("/api")
	auth.RegisterSubRouters(apiGroup)
}

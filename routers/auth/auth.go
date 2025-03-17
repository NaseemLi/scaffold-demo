package auth

import (
	"scaffold-demo/controllers/auth"

	"github.com/gin-gonic/gin"
)

func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

func logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}

func RegisterSubRouters(g *gin.RouterGroup) {
	//配置登陆功能的路由策略
	authGroup := g.Group("/auth")
	//登陆的功能
	login(authGroup)
	logout(authGroup)
}

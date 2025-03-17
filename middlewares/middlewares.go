// 中间件层
package middlewares

import (
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func JWTAuth(r *gin.Context) {
	//1. 除了login和logout之外所有的接口 都要验证时候携带token以及是否合法
	requestUrl := r.FullPath()
	logs.Debug(map[string]interface{}{"请求路径": requestUrl}, " ")
	if requestUrl == "api/auth/login" || requestUrl == "api/auth/logpit" {
		logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "登陆和退出不需要验证token")
		r.Next()
		return
	}
	//token
	//其他接口需要验证
	//获取是否需要携带token
	tokenString := r.Request.Header.Get("Authrization")
	if tokenString == "" {
		//说明没有携带token
		r.JSON(200, gin.H{
			"status":  401,
			"message": "请求未携带token,请登陆后尝试",
		})
		r.Abort()
		return
	}
	//如果不为空 要验证token合法性
	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		//说明没有携带token
		r.JSON(200, gin.H{
			"status":  401,
			"message": "token验证未通过",
		})
		r.Abort()
		return
	}
	r.Set("claims", claims)
	r.Next()
}

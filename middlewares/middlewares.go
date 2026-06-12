// 中间件层
package middlewares

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(r *gin.Context) {
	tokenString := r.Request.Header.Get("Authorization")
	if tokenString == "" {
		returnData := config.NewReturnData()
		returnData.Status = 401
		returnData.Message = "请求未携带token,请登陆后尝试"
		r.JSON(200, returnData)
		r.Abort()
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		returnData := config.NewReturnData()
		returnData.Status = 401
		returnData.Message = "token验证未通过"
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	r.Set("claims", claims)
	r.Next()
}

package auth

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登陆的逻辑
func Login(r *gin.Context) {
	// 1.获取前端传递的用户名和密码
	userInfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		r.JSON(200, returnData) //封装规范JSON
		return
	}
	logs.Debug(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "开始验证登陆信息")
	//验证用户名密码是否正确
	//数据库 环境变量
	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		//认证成功
		//生成JWT的token
		ss, err := jwtutil.GenToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"用户名": userInfo.Username, "错误信息": err.Error}, "用户名密码正确但是token失败")
			r.JSON(200, gin.H{
				"status":  401,
				"message": "生成token失败",
			})
			return
		}
		//token正常生成，返回给前端
		logs.Error(map[string]interface{}{"用户名": userInfo.Username}, "登陆成功")
		data := make(map[string]interface{})
		data["token"] = ss
		r.JSON(200, gin.H{
			"status":  200,
			"message": "登陆成功",
			"data":    data,
		})
	} else {
		//用户名 密码错误
		r.JSON(401, gin.H{
			"status":  401,
			"message": "用户名或者密码错误",
		})
	}
}

func Logout(r *gin.Context) {

	//退出
	//实现退出逻辑
	r.JSON(200, gin.H{
		"message": "退出成功",
		"status":  200,
	})
	logs.Debug(nil, "用户已退出")
	//验证用户名密码是否正确
}

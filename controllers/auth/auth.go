package auth

import (
	"scaffold-demo/config"
	"scaffold-demo/models"
	"scaffold-demo/utils"
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
	userInfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		returnData.Status = 400
		returnData.Message = err.Error()
		r.JSON(200, returnData)
		return
	}

	logs.Debug(map[string]interface{}{"用户名": userInfo.Username}, "开始验证登陆信息")

	var user models.User
	if err := config.DB.Where("username = ?", userInfo.Username).First(&user).Error; err != nil {
		returnData.Status = 401
		returnData.Message = "用户名或者密码错误"
		r.JSON(200, returnData)
		return
	}
	if !utils.CheckPassword(userInfo.Password, user.Password) {
		returnData.Status = 401
		returnData.Message = "用户名或者密码错误"
		r.JSON(200, returnData)
		return
	}

	ss, err := jwtutil.GenToken(userInfo.Username)
	if err != nil {
		logs.Error(map[string]interface{}{"用户名": userInfo.Username, "错误信息": err.Error()}, "用户名密码正确但是token失败")
		returnData.Status = 500
		returnData.Message = "生成token失败"
		r.JSON(200, returnData)
		return
	}

	logs.Info(map[string]interface{}{"用户名": userInfo.Username}, "登陆成功")
	returnData.Data["token"] = ss
	returnData.Message = "登陆成功"
	r.JSON(200, returnData)
}

func Logout(r *gin.Context) {
	returnData := config.NewReturnData()
	returnData.Message = "退出成功"
	r.JSON(200, returnData)
	logs.Debug(nil, "用户已退出")
}

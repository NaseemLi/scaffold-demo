// 项目总入口
package main

import (
	"fmt"
	"os"
	"scaffold-demo/config"
	"scaffold-demo/models"
	"scaffold-demo/routers"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("程序开始运行...")

	if err := config.InitDB(); err != nil {
		logs.Error(nil, "数据库连接失败: "+err.Error())
		os.Exit(1)
	}
	if err := models.AutoMigrate(); err != nil {
		logs.Error(nil, "数据库迁移失败: "+err.Error())
		os.Exit(1)
	}
	if err := models.InitAdminUser(); err != nil {
		logs.Error(nil, "初始化管理员账号失败: "+err.Error())
		os.Exit(1)
	}

	r := gin.Default()
	logs.Info(nil, "程序启动成功")
	//测试生成jwt token是否可用
	// ss, _ := jwtutil.GenToken("ddd") // 这里使用 ss
	// fmt.Println("测试是否能生成token:", ss)

	routers.RegisterRouters(r)
	r.Run(config.Port)
}

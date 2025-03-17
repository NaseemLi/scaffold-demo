// 项目总入口
package main

import (
	"fmt"
	"scaffold-demo/config"
	_ "scaffold-demo/config"
	"scaffold-demo/routers"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	//1. 加载程序配置
	//2. 配置gin
	fmt.Println("程序开始运行...")
	r := gin.Default()
	logs.Info(nil, "程序启动成功")
	//测试生成jwt token是否可用
	// ss, _ := jwtutil.GenToken("ddd") // 这里使用 ss
	// fmt.Println("测试是否能生成token:", ss)

	routers.RegisterRouters(r)
	r.Run(config.Port)
}

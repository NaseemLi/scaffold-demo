// 存放文件配置信息
package config

import (
	"scaffold-demo/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port       string
	JwtSignKey string
	JwtExpTime int64 //jwt token过期时间 单位：分钟
	Username   string
	Password   string
)

type RetrunData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数
func NewReturnData() RetrunData {
	returnData := RetrunData{}
	returnData.Status = 200
	data := make(map[string]interface{})
	returnData.Data = data
	return returnData
}

func initLogConfig(logLevel string) {
	//配置程序的日志输出级别
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	//文件名和行号
	logrus.SetReportCaller(true)
	//日志格式改为json
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "TimeFormat"})
}

func init() {
	logs.Debug(nil, "开始加载程序配置")
	//环境变量加载我们的程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	//获取程序启动端口号配置
	viper.SetDefault("PORT", ":8080")
	//获取jwt加密的secret
	viper.SetDefault("JWT_SIGN_KEY", "lizeyu")
	//获取jwt过期时间的配置
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	//获取用户名和密码
	//加密用户名和密码 md5
	//lizeyu password
	viper.SetDefault("USERNAME", "90A6F4835082FF380C3E94C7D0456BEA")
	viper.SetDefault("PASSWORD", "5F4DCC3B5AA765D61D8327DEB882CF99")

	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME")
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	//加载日志输出格式
	initLogConfig(logLevel)
}

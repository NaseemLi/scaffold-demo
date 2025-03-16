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
	Port string
)

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
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	port := viper.GetString("PORT")
	//加载日志输出格式
	initLogConfig(logLevel)
}

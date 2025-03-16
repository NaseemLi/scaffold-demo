package logs

import "github.com/sirupsen/logrus"

// 打印debug类型的日志
func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}
func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}
func Warning(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}
func Error(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}

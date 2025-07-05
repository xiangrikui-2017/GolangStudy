package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func InitLogrus() {
	// 日志文件
	file, err := os.OpenFile(Conf.Log.LogDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Logger.Out = file
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = Logger.Out
	// 设置日志级别
	log_level := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
	}
	Logger.Level = log_level[Conf.Log.LogLevel]

	// 日志格式
	Logger.Formatter = &logrus.JSONFormatter{}
}

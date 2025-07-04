package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func InitLogrus() {
	// 日志格式
	Logger.Formatter = &logrus.JSONFormatter{}
	// 日志文件
	f, _ := os.Create(Conf.Log.LogDir)
	Logger.Out = f
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
}

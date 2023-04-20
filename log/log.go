package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel) // 设置 日志等级
	logger.SetFormatter(&logrus.JSONFormatter{
		// PrettyPrint: true,
		// FieldMap: logrus.FieldMap{"callers":},
	}) // json格式
	file, err := os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.SetOutput(os.Stdout)
	}
	return logger
}

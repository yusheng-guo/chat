package global

import (
	"github.com/sirupsen/logrus"
	"github.com/yushengguo557/chat/log"
)

var Logger *logrus.Logger

func InitLog() {
	Logger = log.NewLog()
}

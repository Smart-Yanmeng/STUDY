package main

import "github.com/sirupsen/logrus"

func main() {
	// 修改日志输出的模式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 在日志中添加一个指定的字段，并且可以加多个字段
	log := logrus.WithField("app", "study").WithField("service", "logrus")

	// 在日志中添加多个指定的字段，并且会覆盖前面添加的字段
	//log = logrus.WithFields(logrus.Fields{
	//	"user": "York",
	//	"ip":   "192.168.200.254",
	//})

	// 如果想要添加不覆盖
	log = log.WithFields(logrus.Fields{
		"user": "York",
		"ip":   "192.168.200.254",
	})

	log.Errorf("你好")
}

package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		// 修改输出文字颜色
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetLevel(logrus.DebugLevel)

	logrus.Errorln("error")
	logrus.Infoln("info")
	logrus.Warnln("warn")
	logrus.Debugln("debug")
	logrus.Println("print")
}

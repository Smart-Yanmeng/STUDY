package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type MyHook struct {
}

func (hook MyHook) Levels() []logrus.Level {
	// return logrus.AllLevels

	// 修改日志输出等级
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook MyHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = "York"

	file, _ := os.OpenFile("goWeb/gin/unit7/demo07/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	line, _ := entry.String()

	file.Write([]byte(line))

	return nil
}

func main() {
	logrus.AddHook(&MyHook{})

	logrus.Warnln("WARN")
	logrus.Errorln("ERROR")
}

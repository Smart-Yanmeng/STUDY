package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	// 设置输出日志到文件
	file, _ := os.OpenFile("goWeb/logrus/unit1/demo05/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// 设置多个输出方法
	// writes := []io.Writer{
	// 	 file,
	//	 os.Stdout,
	// }

	// logrus.SetOutput(io.MultiWriter(writes...))
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))

	logrus.Infoln("INFO")
	logrus.Errorln("ERROR")
}

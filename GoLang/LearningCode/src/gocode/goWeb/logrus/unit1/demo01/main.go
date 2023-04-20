package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Error("出错了")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	// debug 信息默认不会输出
	logrus.Debugf("debug")
	logrus.Println("打印")

	// 获取打印日志等级，等级比较小的不会被打印出来
	fmt.Println(logrus.GetLevel())

	// 修改打印日志等级
	logrus.SetLevel(logrus.DebugLevel)

	// 这个 debug 信息就能被输出了
	logrus.Debugf("debug")
}

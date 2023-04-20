package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

const (
	cBlack  = 0
	cRed    = 1
	cGreen  = 2
	cYellow = 3
	cBlue   = 4
	cPurple = 5
	cCyan   = 6
	cGray   = 7
)

// PrintColor - 自封装函数
func PrintColor(colorCode int, text string, isBackgroundColor bool) {
	if isBackgroundColor {
		fmt.Printf("\033[4%dm%s\033[0m", colorCode, text)

		return
	}
	fmt.Printf("\033[3%dm%s\033[0m", colorCode, text)
}

type MyFormatter struct {
	Prefix     string
	TimeFormat string
}

func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var color int

	// 根据日志等级设置输出颜色
	switch entry.Level {
	case logrus.ErrorLevel:
		color = cRed
	case logrus.WarnLevel:
		color = cYellow
	case logrus.InfoLevel:
		color = cGreen
	case logrus.DebugLevel:
		color = cCyan
	default:
		color = cGray
	}

	// 设置 buffer 缓冲区
	var b *bytes.Buffer

	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	// 时间格式化
	formatTime := entry.Time.Format(f.TimeFormat)

	// 设置文件的行号
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

	// 设置格式
	fmt.Fprintf(b, "[%s] \033[3%dm[%s]\033[0m [%s] %s %s\n", f.Prefix, color, entry.Level, formatTime, fileVal, entry.Message)

	return b.Bytes(), nil
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&MyFormatter{Prefix: "GORM", TimeFormat: "2006-01-02 15:04:05"})

	logrus.Infoln("HELLO")
	logrus.Warnln("HELLO")
	logrus.Errorln("HELLO")
	logrus.Debugln("HELLO")
}

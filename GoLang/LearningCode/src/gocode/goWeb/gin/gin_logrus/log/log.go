package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type FileDataHook struct {
	file     *os.File
	logPath  string
	fileData string
	appName  string
}

func (hook FileDataHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook FileDataHook) Fire(entry *logrus.Entry) error {
	time := entry.Time.Format("2006-01-02_15-04")
	line, _ := entry.String()

	// 时间相等的情况下，在原来的文件继续写
	if hook.fileData == time {
		hook.file.Write([]byte(line))

		return nil
	}
	// 时间不相等的情况下，新增一个文件
	hook.file.Close()
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, time), os.ModePerm)
	fileName := fmt.Sprintf("%s/%s/%s.log", hook.logPath, time, hook.appName)

	hook.file, _ = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileData = time
	hook.file.Write([]byte(line))

	return nil
}

type MyFormatter struct {
	Prefix     string
	TimeFormat string
}

func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 设置 buffer 缓冲区
	var b *bytes.Buffer

	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	// 设置格式
	fmt.Fprintf(b, "%s\n", entry.Message)

	return b.Bytes(), nil
}

func InitFile(logPath, appName string) {
	logrus.SetFormatter(&MyFormatter{})

	fileData := time.Now().Format("2006-01-02_15-04")

	// 创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileData), os.ModePerm)
	if err != nil {
		logrus.Error(err)

		return
	}

	fileName := fmt.Sprintf("%s/%s/%s.log", logPath, fileData, appName)

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)

		return
	}

	fileHook := FileDataHook{file, logPath, fileData, appName}

	logrus.AddHook(&fileHook)
}

package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	allLog  = "all"
	errLog  = "err"
	warnLog = "warn"
	infoLog = "info"
)

type FileLevelHook struct {
	file     *os.File
	errFile  *os.File
	warnFile *os.File
	infoFile *os.File
	logPath  string
}

func (hook FileLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook FileLevelHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	switch entry.Level {
	case logrus.ErrorLevel:
		hook.errFile.Write([]byte(line))
	case logrus.WarnLevel:
		hook.warnFile.Write([]byte(line))
	case logrus.InfoLevel:
		hook.infoFile.Write([]byte(line))
	}
	hook.file.Write([]byte(line))

	return nil
}

func InitLevel(logPath string) {
	err := os.MkdirAll(fmt.Sprintf("%s", logPath), os.ModePerm)
	if err != nil {
		logrus.Error(err)

		return
	}

	allFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, allLog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	errFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, errLog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	warnFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, warnLog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	infoFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, infoLog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)

	fileHook := FileLevelHook{allFile, errFile, warnFile, infoFile, logPath}
	logrus.AddHook(&fileHook)
}

func main() {
	InitLevel("goWeb/gin/unit7/demo09/log")

	logrus.Errorln("ERROR")
	logrus.Warnln("WARN")
	logrus.Infoln("INFO")
}

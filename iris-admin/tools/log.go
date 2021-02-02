package tools

import (
	"log"
	"os"
	"time"

	"github.com/kataras/golog"
)

const (
	sysTime      = "2006/01/02 - 15:04:05"
	sysTimeShort = "2006-01-02"
	dir          = "./log/"
)

func init() {
	var (
		err error
	)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	setLogger()
	golog.Info("日志配置项启动")
	return
}

func setLogger() {
	var (
		logFile = newLogFile(dir + time.Now().Format(sysTimeShort) + ".log")
	)
	golog.AddOutput(logFile)
}

func newLogFile(filename string) *os.File {
	var (
		file *os.File
		err  error
	)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

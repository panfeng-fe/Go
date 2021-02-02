package log

import (
	"time"

	logger "github.com/phachon/go-logger"
)

// Logger ...
var Logger *logger.Logger

func init() {
	Logger = logger.NewLogger()
	Logger.Detach("console")
	consoleConfig := &logger.ConsoleConfig{
		Color:      true,
		JsonFormat: false,
		Format:     "",
	}
	fileConfig := &logger.FileConfig{
		LevelFileName: map[int]string{
			Logger.LoggerLevel("error"): "./logs/" + time.Now().Format("2006-01-02") + "-error.log", // Error 级别日志被写入 error .log 文件
			Logger.LoggerLevel("info"):  "./logs/" + time.Now().Format("2006-01-02") + "-info.log",  // Info 级别日志被写入到 info.log 文件中
			Logger.LoggerLevel("debug"): "./logs/" + time.Now().Format("2006-01-02") + "-debug.log", // Debug 级别日志被写入到 debug.log 文件中
		},
		MaxSize:    0,
		MaxLine:    0,
		DateSlice:  "d",
		JsonFormat: false,
	}
	Logger.Attach("console", logger.LOGGER_LEVEL_DEBUG, consoleConfig)
	Logger.Attach("file", logger.LOGGER_LEVEL_DEBUG, fileConfig)
}

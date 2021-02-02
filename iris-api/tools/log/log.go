package log

import (
	"time"

	logger "github.com/phachon/go-logger"
)

// Log ...
var Log *logger.Logger

func init() {
	Log = logger.NewLogger()
	Log.Detach("console")
	consoleConfig := &logger.ConsoleConfig{
		Color:      true,
		JsonFormat: false,
		Format:     "",
	}
	fileConfig := &logger.FileConfig{
		LevelFileName: map[int]string{
			Log.LoggerLevel("error"): "./logs/" + time.Now().Format("2006-01-02") + "-error.log", // Error 级别日志被写入 error .log 文件
			Log.LoggerLevel("info"):  "./logs/" + time.Now().Format("2006-01-02") + "-info.log",  // Info 级别日志被写入到 info.log 文件中
			Log.LoggerLevel("debug"): "./logs/" + time.Now().Format("2006-01-02") + "-debug.log", // Debug 级别日志被写入到 debug.log 文件中
		},
		MaxSize:    0,
		MaxLine:    0,
		DateSlice:  "d",
		JsonFormat: false,
	}
	Log.Attach("console", logger.LOGGER_LEVEL_DEBUG, consoleConfig)
	Log.Attach("file", logger.LOGGER_LEVEL_DEBUG, fileConfig)
}

package config

import (
	"native/model/app"
	"native/tools/log"

	"gopkg.in/gcfg.v1"
)

// Setting 配置信息
var Setting = &app.Config{}
var Log = log.Logger

func init() {
	if err := gcfg.ReadFileInto(Setting, "./config/app.ini"); err != nil {
		Log.Infof("Failed to parse config file: %s", err)
		panic("配置解析异常")
	}
	Log.Info("项目配置项初始化完毕")

}

package config

import (
"iris-api/model/app"
"iris-api/tools/log"
"gopkg.in/gcfg.v1"
)

// Setting 配置信息
var Setting = &app.Config{}
var Log = log.Log
func init() {
	if err := gcfg.ReadFileInto(Setting, "./config/app.ini"); err != nil {
		Log.Infof("Failed to parse config file: %s", err)
		return
	}
	Log.Info("项目配置项初始化完毕")

}


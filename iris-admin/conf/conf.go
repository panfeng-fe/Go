package conf

import (
	"gopkg.in/gcfg.v1"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	_"iris-admin/tools"
)

var Setting = &Config{}

func init(){
    if err := gcfg.ReadFileInto(Setting, "./conf/app.ini");err != nil {
		panic("异常信息")
		golog.Println("Failed to parse config file: %s", err)
	}
	golog.Info("项目配置项初始化完毕")
	return
}

type Config struct{
	App struct{
		Name string
		Port int
	}
	DB struct{
		User     string
		Password string
		Host     string
		Dbname   string
		Port     int
	}
	Project iris.Configuration
}

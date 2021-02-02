package model

import (
	// 数据库驱动
	"fmt"
	"iris-api/config"
	. "iris-api/tools/log"
	"strconv"

	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

var err error

// Engin 全局数据库引擎
var Engin *gorose.Engin

func init() {
	var config = config.Setting.DB
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.User, config.Password, config.Host, strconv.Itoa(config.Port), config.Dbname)

	var DbConfig = &gorose.Config{Driver: "mysql", Dsn: dsn}
	Engin, err = gorose.Open(DbConfig)
	if err != nil {
		Log.Infof("数据库连接失败:%s", err)
		return
	}
}

// Mysql 初始化数据库
func Mysql() gorose.IOrm {
	return Engin.NewOrm()
}

// sqlserver dsn
// var dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.Dbname)

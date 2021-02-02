package db

import (
	// 数据库驱动
	"fmt"
	"iris-admin/conf"
	"strconv"

	"github.com/gohouse/gorose/v2"
	// 数据库驱动
	_ "github.com/denisenkom/go-mssqldb"
)

var err error
var engin *gorose.Engin

func init() {
	var config = conf.Setting.DB
	// var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",config.User,config.Password,config.Host,strconv.Itoa(config.Port),config.Dbname,)
	var dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.Dbname)
	var DbConfig = &gorose.Config{Driver: "mssql", Dsn: dsn}
	engin, err = gorose.Open(DbConfig)
	if err != nil {
		fmt.Print("数据库连接失败", err)
	}
}

//Database 初始化数据库
func Database() gorose.IOrm {
	return engin.NewOrm()
}

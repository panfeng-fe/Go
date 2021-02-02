package main

import (
	"fmt"
	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

var err error
var engin *gorose.Engin

func init() {
	// mysql Dsn示例 "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true"
	var DbConfig = &gorose.Config{Driver: "mysql", Dsn: "root:root@tcp(localhost:3306)/world?charset=utf8&parseTime=true"}
	engin, err = gorose.Open(DbConfig)
}

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	var db = engin.NewOrm()

	//根据查询的结果创建结构体，将查询出来的结果直接放在切片里，十分便于操作
	var res []user
	_, err := db.Table(&res).Query("select * from table")
	// gorose提供这种占位符方式查询，避免sql拼接，十分便捷
	//_, err := db.Table(&res).Query("select top(?)* from table",num)

	//这是普通的数据库查询，查询出来的结果是map类型，根据个人需求采用不同的方法
	// res, err := db.Query("select * from table")
	fmt.Print(res, err)
}

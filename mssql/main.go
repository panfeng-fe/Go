package main

import (
	"fmt"
	// 与mysql的第一个区别就是数据库驱动不同
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gohouse/gorose"
)

var err error
var engin *gorose.Engin

func init() {
	// 第二个区别是由于驱动不同，所以连接数据库的dsn可能有所变化，其他都一样
	var DbConfig = &gorose.Config{Driver: "mssql", Dsn: "server=127.0.0.1;user id=root;password=root;port=1433;database=world;encrypt=disable"}
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

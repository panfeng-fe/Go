package app

import "github.com/kataras/iris/v12"

// Config 数据库结构体
type Config struct {
	App     App
	Mysql   Mysql
	Mssql   Mssql
	Project iris.Configuration
}

// App ...
type App struct {
	Name     string
	Port     int
	TokenKey string
}

// Mysql ...
type Mysql struct {
	User     string
	Password string
	Host     string
	Dbname   string
	Port     int
}

// Mssql ...
type Mssql struct {
	User                string
	Password            string
	Host                string
	Port                int
	RYAccountsDB        string
	RYGameScoreDB       string
	RYLotteryDB         string
	RYPlatformDB        string
	RYPlatformManagerDB string
	RYRecordDB          string
	RYTreasureDB        string
	RYNativeWebDB       string
}

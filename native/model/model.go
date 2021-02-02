package model

import (
	"fmt"
	"strconv"

	"native/config"
	"native/tools/log"

	// 数据库驱动
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gohouse/gorose/v2"
)

// 全局数据库引擎
var (
	err             error
	Accounts        *gorose.Engin
	Lottery         *gorose.Engin
	Treasure        *gorose.Engin
	Record          *gorose.Engin
	Platform        *gorose.Engin
	PlatformManager *gorose.Engin
	GameScore       *gorose.Engin
	NativeWeb       *gorose.Engin
	Log             = log.Logger
)

func init() {
	var config = config.Setting.Mssql
	var (
		accounts        = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYAccountsDB)
		lottery         = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYGameScoreDB)
		treasure        = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYLotteryDB)
		record          = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYPlatformDB)
		platform        = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYPlatformManagerDB)
		platformManager = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYRecordDB)
		gameScore       = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYTreasureDB)
		nativeWeb       = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", config.Host, config.User, config.Password, strconv.Itoa(config.Port), config.RYNativeWebDB)
	)
	Accounts, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: accounts})
	if err != nil {
		Log.Infof("数据库RYAccountsDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	Lottery, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: lottery})
	if err != nil {
		Log.Infof("数据库RYLotteryDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	Treasure, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: treasure})
	if err != nil {
		Log.Infof("数据库RYTreasureDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	Record, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: record})
	if err != nil {
		Log.Infof("数据库RYRecordDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	Platform, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: platform})
	if err != nil {
		Log.Infof("数据库RYPlatformDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	PlatformManager, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: platformManager})
	if err != nil {
		Log.Infof("数据库RYPlatformManagerDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	GameScore, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: gameScore})
	if err != nil {
		Log.Infof("数据库RYGameScoreDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	NativeWeb, err = gorose.Open(&gorose.Config{Driver: "mssql", Dsn: nativeWeb})
	if err != nil {
		Log.Infof("数据库RYNativeWebDB连接失败:%s", err)
		panic("初始数据库失败")
	}
	Log.Info("数据库初始化成功")
}

// Mysql 初始化数据库
func RYAccountsDB() gorose.IOrm {
	return Accounts.NewOrm()
}
func RYLotteryDB() gorose.IOrm {
	return Lottery.NewOrm()
}
func RYTreasureDB() gorose.IOrm {
	return Treasure.NewOrm()
}
func RYRecordDB() gorose.IOrm {
	return Record.NewOrm()
}
func RYPlatformDB() gorose.IOrm {
	return Platform.NewOrm()
}
func RYPlatformManagerDB() gorose.IOrm {
	return PlatformManager.NewOrm()
}
func RYGameScoreDB() gorose.IOrm {
	return GameScore.NewOrm()
}
func RYNativeWebDB() gorose.IOrm {
	return NativeWeb.NewOrm()
}

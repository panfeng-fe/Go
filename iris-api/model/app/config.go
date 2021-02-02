package app

import "github.com/kataras/iris"

// Config 数据库结构体
type Config struct {
	App struct {
		Name     string
		Port     int
		TokenKey string
	}
	DB struct {
		User     string
		Password string
		Host     string
		Dbname   string
		Port     int
	}
	DG struct {
		Key             string
		Account         string
		WebURL          string
		Limit           int
		DefaultXianHong string
		UseRoomLevel    int
		ServerID        int
		KindID          int
		Index           int
		LoginKey    	string
	}
	Project iris.Configuration
}

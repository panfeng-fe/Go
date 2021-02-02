package controller

import (
	"fmt"
	"native/model/app"
	"native/tools/log"

	"github.com/kataras/iris/v12"
)

//var res = app.Context{
//	State:   false,
//	Msg:     "",
//	Content: nil,
//}
var err error

// Log ...
var Log = log.Logger

// Testing ...
func Testing(ctx iris.Context) {
	fmt.Print(ctx.Path())
}

// GetPages ...
func GetPages(ctx iris.Context) (page app.Page, err error) {
	page.Size, err = ctx.URLParamInt("size")
	if err != nil {
		return page, err
	}
	page.Page, err = ctx.URLParamInt("page")
	return page, err
}

// GetPostGameID ...
func GetPostGameID(ctx iris.Context) (GameID int, err error) {
	GameID, err = ctx.PostValueInt("GameID")
	return GameID,err
}

// GetPages ...
func GetGameID(ctx iris.Context) (GameID int, err error) {
	GameID, err = ctx.URLParamInt("GameID")
	return GameID,err
}

// GetPages ...
func GetURLGameID(ctx iris.Context) (GameID int, err error) {
	GameID, err = ctx.Params().GetInt("id")
	return GameID,err
}

func Info(ctx iris.Context){
	fmt.Print(ctx.Path())
}

func Find(ctx iris.Context){
	fmt.Print(ctx.Path())
}

func GetSelectSql(ctx iris.Context)(){}
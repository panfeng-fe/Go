package controller

import (
	"native/db"
	"native/model/app"

	"github.com/kataras/iris/v12"
)

// PlayerInfo 获取玩家信息
func PlayerInfo(ctx iris.Context) {
	var res = app.Context{
		State:   false,
		Msg:     "",
		Content: nil,
		Len:     0,
	}
	var page app.Page

	// 阶段一 参数验证
	page.Size, err = ctx.URLParamInt("size")
	if err != nil {
		Log.Errorf("接受参数失败size:%s!", err)
		res.Msg = "接受参数失败"
		ctx.JSON(res)
		return
	}
	page.Page, err = ctx.URLParamInt("page")
	if err != nil {
		Log.Errorf("接受参数失败page:%s!", err)
		res.Msg = "接受参数失败"
		ctx.JSON(res)
		return
	}

	// 阶段二 传入获取页面参数 获取页面类容与总人数
	res.Content, res.Len, err = db.SelectPlayer(page)
	if err != nil {
		Log.Errorf("查询玩家列表失败!，err:%s", err)
		res.Msg = "查询玩家列表失败"
		ctx.JSON(res)
		return
	}

	// 阶段三 返回数据
	res.State = true
	res.Msg = "数据成功返回"
	ctx.JSON(res)
}

// PlayerFind 获取个人玩家信息
func PlayerFind(ctx iris.Context) {
	var res = app.Context{
		State:   false,
		Msg:     "",
		Content: nil,
		Len:     0,
	}
	// 获取GameID
	GameID, err := ctx.Params().GetInt("id")
	if err != nil {
		Log.Errorf("接受参数失败:%s!", err)
		res.Msg = "接受参数失败"
		ctx.JSON(res)
		return
	}
	// 查询玩家
	res.Content, err = db.FindPlayer(GameID)
	if err != nil {
		Log.Errorf("查询玩家%s失败! err:%s", GameID, err)
		res.Msg = "查询玩家失败"
		ctx.JSON(res)
		return
	}
	res.State = true
	res.Msg = "数据成功返回"
	ctx.JSON(res)
}

// PlayerModify 修改个人玩家信息
func PlayerModify(ctx iris.Context) {}

// PlayerForbid 禁用个人玩家信息
func PlayerForbid(ctx iris.Context) {
	var res = app.Context{
		State:   false,
		Msg:     "",
		Content: nil,
		Len:     0,
	}
	GameID, err := ctx.PostValueInt("GameID")
	if err != nil {
		Log.Errorf("禁用玩家API接受参数失败! %s：", err)
		res.Msg = "接受参数失败"
		ctx.JSON(res)
		return
	}
	State, err := ctx.PostValueInt("State")
	if err != nil {
		Log.Infof("禁用玩家API接受参数失败! %s：", err)
		res.Msg = "接受参数失败"
		ctx.JSON(res)
		return
	}
	err = db.UpdatePlayer(GameID, State)
	if err != nil {
		Log.Errorf("更改玩家%s状态，接受参数失败! %s：", GameID, err)
		res.Msg = "更改玩家状态失败"
		ctx.JSON(res)
		return
	}
	res.State = true
	res.Msg = "更改玩家状态成功"
	ctx.JSON(res)
}

// PlayerUserInfo 获取个人玩家详细信息
func PlayerUserInfo(ctx iris.Context) {
	var (
		res        app.Context
		UserInfo   app.UserInfo
	)
	GameID, err := ctx.PostValueInt("GameID")
	if err != nil {
		Log.Errorf("获取玩家信息，接受参数失败! %s：", err)
		res.Msg = "接受参数失败"
		ctx.JSON(res)
		return
	}

	UserInfo.Info, UserInfo.BankCard, err = db.SelectUserInfo(GameID)
	if err != nil {
		Log.Errorf("获取玩家信息，数据库查询失败!，err:%s", err)
		res.Msg = "查询玩家列表失败"
		ctx.JSON(res)
		return
	}
	res.State = true
	res.Msg = "玩家信息返回成功！"
	res.Content = UserInfo
	ctx.JSON(res)

}

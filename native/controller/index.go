package controller

import (
	"native/db"
	"native/model/app"

	"github.com/kataras/iris/v12"
)

// Login 登录.
func Login(ctx iris.Context) {
	var res = app.Context{
		State:   false,
		Msg:     "",
		Content: nil,
		Len:     0,
	}
	var param = app.LoginParam{
		Account:  ctx.PostValue("account"),
		Password: ctx.PostValue("password"),
	}
	if param.Account == "" || param.Password == "" {
		Log.Info("登录参数错误!")
		res.Msg = "登录参数错误"
		ctx.JSON(res)
		return
	}
	data,err := db.SelectLoginParams(param.Account)
	if err != nil {
		Log.Infof("查询登录参数错误! err：%s",err)
		res.Msg = "查询登录参数错误!"
		ctx.JSON(res)
		return
	}
	if data.Username == "" {
		Log.Info("查无此人!")
		res.Msg = "查无此人"
		ctx.JSON(res)
		return
	}
	if data.Nullity != 0 {
		Log.Info("您已经被禁用!")
		res.Msg = "您已经被禁用!"
		ctx.JSON(res)
		return
	}
	if param.Password != data.Password {
		Log.Infof("用户%s密码不对!", param.Account)
		res.Msg = "用户密码不对!"
		ctx.JSON(res)
		return
	}
	res.Msg = "登陆成功!"
	res.State = true
	ctx.JSON(res)
	return
}

// Index 首页。
func Index(ctx iris.Context) {
	var res = app.Context{
		State:   false,
		Msg:     "",
		Content: nil,
	}
	data,err := db.SelectIndex()
	if err != nil {
		Log.Infof("首页数据查询错误! err：%s",err)
		res.Msg = "首页数据查询错误!"
		ctx.JSON(res)
		return
	}
	res.Content = app.IndexParam{
		AllPlayer:       data[0].Month,
		TodayPlayer:     data[0].Today,
		AllAgent:        data[1].Month,
		TodayAgent:      data[1].Today,
		MonthPresent:    data[2].Month,
		TodayPresent:    data[2].Today,
		MonthCaijin:     data[3].Month,
		TodayCaijin:     data[3].Today,
		MonthRecharge:   data[4].Month,
		TodayRecharge:   data[4].Today,
		MonthWithdrawal: data[5].Month,
		TodayWithdrawal: data[5].Today,
	}
	res.State = true
	res.Msg = "数据成功返回"
	ctx.JSON(res)
}

package controller

import (
	"native/db"
	"native/model/app"

	"github.com/kataras/iris/v12"
)

// ListInfo 订单信息
func ListInfo(ctx iris.Context) {
	var res app.Context
	page, err := GetPages(ctx)
	if err != nil {
		Log.Errorf("获取分页参数失败:%s!", err)
		res.Msg = "获取分页参数失败"
		ctx.JSON(res)
		return
	}
	res.Content, res.Len, err = db.ListInfo(page)
	if err != nil {
		Log.Errorf("查询分页失败:%s!", err)
		res.Msg = "查询分页失败"
		ctx.JSON(res)
		return
	}
	res.State = true
	res.Msg = "查询分页成功"
	ctx.JSON(res)
}

// ListFind 订单查询
func ListFind(ctx iris.Context){
	var res app.Context
	GameID ,err := GetURLGameID(ctx)
	if err != nil {
		Log.Errorf("接收查询参数失败:%s!", err)
		res.Msg = "接收查询参数失败"
		ctx.JSON(res)
		return
	}
	res.Content,err = db.ListFind(GameID)
	if err != nil {
		Log.Errorf("查询过程失败:%s!", err)
		res.Msg = "查询过程失败"
		ctx.JSON(res)
		return
	}
	res.State = true
	res.Msg = "成功返回数据"
	ctx.JSON(res)
}

// ListMsg 未读信息
func ListMsg(ctx iris.Context){
	var res app.Context
	res.Len,err = db.ListMsg()
	if err != nil {
		Log.Errorf("查询未处理订单失败:%s!", err)
		res.Msg = "查询未处理订单失败"
		ctx.JSON(res)
		return
	}
	res.State = true
	res.Msg = "成功返回数据"
	ctx.JSON(res)
}

// PresentInfo 赠送信息
func PresentInfo(ctx iris.Context){
	var res app.Context
	page, err := GetPages(ctx)
	if err != nil {
		Log.Errorf("获取分页参数失败:%s!", err)
		res.Msg = "获取分页参数失败"
		ctx.JSON(res)
		return
	}
	res.Content, res.Len, err = db.ListInfo(page)
	if err != nil {
		Log.Errorf("查询分页失败:%s!", err)
		res.Msg = "查询分页失败"
		ctx.JSON(res)
		return
	}
	res.State = true
	res.Msg = "查询分页成功"
	ctx.JSON(res)


}

// PresentFind 赠送查询
func PresentFind(ctx iris.Context){}

// ChargeInfo 充值信息
func ChargeInfo(ctx iris.Context){}

// ChargeFind 充值查询
func ChargeFind(ctx iris.Context){}

// RejectInfo 拒绝信息
func RejectInfo(ctx iris.Context){}

// RejectFind 拒绝查询
func RejectFind(ctx iris.Context){}

// WithdrawalInfo 提现信息
func WithdrawalInfo(ctx iris.Context){}

// WithdrawalFind 提现查询
func WithdrawalFind(ctx iris.Context){}
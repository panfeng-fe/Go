package controller

import (
	"github.com/kataras/iris"
	"iris-api/config"
)

var err error
var DG = config.Setting.DG
// login ...
//func Login(ctx iris.Context){
//	// recive Login params
//	loginParams := app.GetLoginParams{}
//	resObj := app.Context{}
// 	loginParams.UserID,err = ctx.Params().GetInt("userid")
// 	if err != nil {
//		resObj.Msg = "用户UserID未成功接受！"
//		resObj.State = false
//		ctx.JSON(resObj)
//		return
//	}
//	loginParams.UserName = ctx.Params().Get("username")
//	loginParams.Time = ctx.Params().Get("time")
//	loginParams.Signature = ctx.Params().Get("signature")
//	loginParams.Type = strings.ToUpper(DG.DefaultXianHong)
//
//	/*
//		1 签名参数错误（失败）
//		2 用户不存在,动态密码查询（失败）
//		3 签名错误（失败）
//		4 程序运行出错（失败）
//		5 签名成功（成功）
//	 */
//    state := dg_client.CheckUserSignature(loginParams)
//	switch state {
//		case 1:
//			resObj.Msg = "签名参数错误！"
//			resObj.State = false
//			ctx.JSON(resObj)
//			return
//		case 2:
//			resObj.Msg = "用户不存在或动态密码不存在！"
//			resObj.State = false
//			ctx.JSON(resObj)
//			return
//		case 3:
//			resObj.Msg = "签名错误！"
//			resObj.State = false
//			ctx.JSON(resObj)
//			return
//		case 4:
//			resObj.Msg = "程序运行出错！"
//			resObj.State = false
//			ctx.JSON(resObj)
//			return
//		default:
//
//	}
//	random := dg_client.CreateCaptcha()
//	token := dg_client.MD5Signature(DG.Account,DG.Key,random)
//
//	var postLoginParams app.PostLoginParams
//	postLoginParams.UserID = loginParams.UserID
//	postLoginParams.UserName = loginParams.UserName
//	postLoginParams.Token = token
//	postLoginParams.LastLogin = time.Now().Format("2006-01-02 15:04:05")
//	postLoginParams.LastIp = ctx.Request().RemoteAddr
//	postLoginParams.Channel = "DG"
//	postLoginParams.AgentName = DG.Account
//
//	isUserToken := db.GetUserVideoToken(loginParams.UserID)
//
//	if isUserToken == false {
//		resObj.Msg = "用户未注册！"
//		resObj.State = false
//		ctx.JSON(resObj)
//		return
//	}
//	// UseRoomLevel = 1 限红
//	if DG.UseRoomLevel == 1 {
//		if state := dg_client.ChangeUserLimit(postLoginParams,loginParams.Type); state == false {
//			Log.Info("用户限红失败！")
//			return
//		}
//	}
//
//	loginRes,isLoginState := dg_client.DGLogin(postLoginParams)
//	if isLoginState == false {
//		Log.Info("用户登录失败！")
//		return
//	}
//
//	db.WriteUserVideoToken(postLoginParams,loginParams.UserName)
//
//
//
//
//
//}

// Detail ...
func Detail(ctx iris.Context){

}

// Free ...
func Free(ctx iris.Context){

}



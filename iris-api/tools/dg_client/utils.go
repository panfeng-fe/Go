package dg_client

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	// 日志
	. "iris-api/tools/log"
	"iris-api/config"
	"math/rand"
	"strconv"
	"time"

	"crypto/md5"
	"iris-api/db"
	"iris-api/model/app"
)
var DG = config.Setting.DG
/*
	state 返回参数
	1 签名参数错误（失败）
	2 用户不存在,动态密码查询（失败）
	3 签名错误（失败）
	4 程序运行出错（失败）
	5 签名成功（成功）
*/
func CheckUserSignature(loginStruct app.GetLoginParams)(state int){
	if loginStruct.UserID == 0 || loginStruct.Signature == "" || loginStruct.Time == "" {
		state = 1
		return state
	}
	Info,err := db.GetLoginInfo(loginStruct.UserID)
	if err != nil {
		state = 4
		return state
	}
	if Info.GameID < 0 || Info.DynamicPass == ""{
		state = 2
		return state
	}
	signature := MD5Signature(strconv.Itoa(loginStruct.UserID),Info.DynamicPass,loginStruct.Time,config.Setting.DG.LoginKey)

	if signature != loginStruct.Signature {
		state = 3
		return state
	}
	state = 5
	return state
}

// MD5Signature 签名
func MD5Signature(args ...string) (signature string){
	var param string
	for _,v := range args{
		param = param + v
	}
	data := []byte(param)
	has := md5.Sum(data)
	signature = fmt.Sprintf("%x", has)
	return signature
}

// 生成制定位数随机数
func CreateCaptcha() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}

// ChangeUserLimit 第三方限红
func ChangeUserLimit(postLoginParams app.PostLoginParams, Type string)(state bool){
	state = false
	userLimit := map[string]interface{}{
		"token" : postLoginParams.Token,
		"random" :  postLoginParams.Random,
		"data" : Type,
		"member" : map[string]interface{}{
			"username" : postLoginParams.UserName,
		},
	}
	postURL := fmt.Sprintf("%s/game/updateLimit/%s",DG.WebURL,DG.Account)
	res := Post(userLimit,postURL)
	res.Body()
	return state
}

//func DGLogin(postLoginParams app.PostLoginParams)(response interface{},state bool){
//	state = false
//	data := map[string]interface{}{
//		"token" : postLoginParams.Token,
//		"random" :  postLoginParams.Random,
//		"data" : Type,
//		"member" : map[string]interface{}{
//			"username" : postLoginParams.UserName,
//		},
//	}
//	postURL := fmt.Sprintf("%s/user/login/%s",DG.WebURL,DG.Account)
//	res := Post(data,postURL)
//	res.Body()
//	//response = res.Body()
//	return response,state
//}

// Post ...
func Post(data map[string]interface{},url string) (res *HttpRequest.Response){
	res, err := HttpRequest.Post(url, data)
	if err != nil {
		Log.Info("请求失败！")
		return
	}
	return res
}
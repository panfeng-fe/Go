package db

import (
	"iris-api/model"
	"iris-api/model/app"
	"iris-api/model/database"
	. "iris-api/tools/log"
)

var err error

//GetUserIDByGameID 通过GameID取UserID
func GetUserIDByGameID(GameID int) (UserID int, err error) {
	var mysql = model.Mysql()
	var res app.DBIntValue
	_, err = mysql.Table(&res).Query("select UserID from AccountsInfo where GameID = ?", GameID)
	UserID = res.IntValue
	return UserID, err
}

//GetUserIDByGameID 通过GameID取UserID
func GetGameIDByUserID(UserID int) (GameID int, err error) {
	var mysql = model.Mysql()
	var res app.DBIntValue
	_, err = mysql.Table(&res).Query("select GameID from AccountsInfo where UserID = ?", UserID)
	UserID = res.IntValue
	return UserID, err
}

//GetLoginInfo 根据UserID 获取动态密码与GameID
func GetLoginInfo(UserID int) (Info app.DBLoginInfo, err error) {
	var mysql = model.Mysql()
	_, err = mysql.Table(&Info).Query("select GameID,UserID,DynamicPass from AccountsInfo where UserID = ?", UserID)
	return Info, err
}

// GetUserVideoToken 查询玩家是否有记录
func GetUserVideoToken(UserID int)(state bool){
	state = false
	var mysql = model.Mysql()
	var res app.DBStringValue
	_, err = mysql.Table(&res).Query("select token from accountsvideologo where user_id = ?", UserID)
	if err != nil {
		Log.Infof("查询%d玩家token失败",UserID)
		return state
	}
	if res.StringValue != "" {
		state = true
		return state
	}else {
		return state
	}
}


// WriteUserVideoToken 写入用户token
func WriteUserVideoToken(postLoginParams app.PostLoginParams,UserName string){
	var mysql = model.Mysql()
	_,err = mysql.Table("Accountsvideologo").Update(&database.Accountsvideologo{
		UserId:postLoginParams.UserID,
		Token:postLoginParams.Token,
		Random:postLoginParams.Random,
		LastIp:postLoginParams.LastIp,
		LastLogin:postLoginParams.LastLogin,
		AgentName:postLoginParams.AgentName,
		Channel:postLoginParams.Channel})
	if err != nil {
		Log.Infof("更新玩家 %d的token失败",postLoginParams.UserID)
		return
	}
	_,err = mysql.Table("Recordvideologo").Insert(&database.Recordvideologo{
		UserId:postLoginParams.UserID,
		Token:postLoginParams.Token,
		UserName:UserName,
		LoginIp:postLoginParams.LastIp,
		LoginTime:postLoginParams.LastLogin,
		Channel:postLoginParams.Channel,
		DateId:0,
	})
}



















// GetJSON []database.Recordcaipiaotoken ...
//func GetJSON() []database.Recordcaipiaotoken {
//	var mysql = model.Mysql()
//	var res []database.Recordcaipiaotoken
//	_, err := mysql.Table(&res).Query("select * from Recordcaipiaotoken  limit 0,10")
//	if err != nil {
//		fmt.Print("select error")
//	}
//	return res
//}

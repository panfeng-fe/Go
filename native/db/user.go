package db

import (
	"native/model/app"
	"strconv"
)

// SelectPlayer(Param app.Page) ... data1 返回切片 data2返回总人数
func SelectPlayer(Param app.Page) (data []app.Player, len int, err error) {
	var amount app.IntValue
	sql := "select top(?) a.UserID,a.GameID,a.Accounts,Convert(int,isnull(b.Score/100,0)) as Score,Convert(int,isnull(b.InsureScore/100,0)) as InsureScore,"
	sql += "isnull(d.GameID,0) as AgentGameID,a.LastLogonIP,a.RegisterIP,ISNULL(e.Compellation,0) as AgentName,CONVERT(VARCHAR(16),a.RegisterDate,120) as RegisterDate,"
	sql += "CONVERT(VARCHAR(16),a.LastLogonDate,120) as LastLogonDate,a.Nullity from [RYAccountsDB].[dbo].[AccountsInfo] as a left join [RYTreasureDB].[dbo].[GameScoreInfo]"
	sql += "as b on a.UserID = b.UserID left join [RYTreasureDB].[dbo].[OnLineOrder] as c on a.UserID = c.UserID left join [RYAccountsDB].[dbo].[AccountsInfo] as d "
	sql += "on a.SpreaderID = d.UserID left join [RYAccountsDB].[dbo].[AccountsAgent] as e on a.SpreaderID = e.UserID  where a.UserID not in (select top((?-1)*15) "
	sql += "UserID FROM [RYAccountsDB].[dbo].[AccountsInfo] order by UserID DESC) group by a.UserID,a.GameID,a.Accounts,b.Score,b.InsureScore,"
	sql += "d.GameID,a.RegisterDate,a.LastLogonDate,a.Nullity,a.LastLogonIP,a.RegisterIP,e.Compellation order by UserID DESC"
	sql2 := "select count(*) as Tatol FROM [RYAccountsDB].[dbo].[AccountsInfo] where UserID > 3000"
	_, err = RYAccountsDB.Table(&data).Query(sql, Param.Size, Param.Page)
	if err != nil {
		return data,len,err
	}
	_, err = RYAccountsDB.Table(&amount).Query(sql2)
	if err != nil {
		return data,len,err
	}
	len = amount.Value
	return data, len, err
}

// FindPlayer ... data1 返回切片 data2返回总人数
func FindPlayer(GameID int) (data1 app.Player, err error) {
	sql := "select a.UserID,a.GameID,a.Accounts,Convert(int,isnull(b.Score/100,0)) as Score,Convert(int,isnull(b.InsureScore/100,0)) as InsureScore,"
	sql += "isnull(d.GameID,0) as AgentGameID,a.LastLogonIP,a.RegisterIP,ISNULL(e.Compellation,0) as AgentName,CONVERT(VARCHAR(16),a.RegisterDate,120) as RegisterDate,"
	sql += "CONVERT(VARCHAR(16),a.LastLogonDate,120) as LastLogonDate,a.Nullity from [RYAccountsDB].[dbo].[AccountsInfo] as a left join [RYTreasureDB].[dbo].[GameScoreInfo]"
	sql += "as b on a.UserID = b.UserID left join [RYTreasureDB].[dbo].[OnLineOrder] as c on a.UserID = c.UserID left join [RYAccountsDB].[dbo].[AccountsInfo] as d "
	sql += "on a.SpreaderID = d.UserID left join [RYAccountsDB].[dbo].[AccountsAgent] as e on a.SpreaderID = e.UserID where a.GameID = ? or a.UserID = ? or a.Accounts = ? "
	sql += "group by a.UserID,a.GameID,a.Accounts,b.Score,b.InsureScore,d.GameID,a.RegisterDate,a.LastLogonDate,a.Nullity,a.LastLogonIP,a.RegisterIP,e.Compellation"
	_, err = RYAccountsDB.Table(&data1).Query(sql, GameID, GameID, strconv.Itoa(GameID))
	return data1, err
}

func UpdatePlayer(GameID int, State int) (err error) {
	sql := "update [RYAccountsDB].[dbo].[AccountsInfo] set Nullity = ? ,NullityOverDate = getdate() where GameID = ?"
	_, err = RYAccountsDB.Query(sql, State, GameID)
	return err
}

func SelectUserInfo(GameID int) (data1 app.UserInfo_Info, data2 app.UserInfo_BankCard, err error) {
	sql := "select a.GameID,a.Accounts,a.LastLogonIP,a.RegisterIP,ISNULL(b.Score/100,0) as Score,ISNULL(b.InsureScore/100,0) as InsureScore,"
	sql += "Convert(int,SUM(CASE WHEN  c.OrderType=0 and c.OrderStatus = 2 Then c.PayAmount else 0 end)) as Recharge,"
	sql += "Convert(int,SUM(CASE WHEN c.orderType=1 and c.OrderStatus = 2 Then c.PayAmount/100 else 0 end)) as Withdrawal,"
	sql += "Convert(int,SUM(CASE WHEN d.Type = 1 Then d.AddGold/100 else 0 end)) as Present FROM [RYAccountsDB].[dbo].[AccountsInfo] as a "
	sql += "left join [RYTreasureDB].[dbo].[GameScoreInfo] as b on a.UserID = b.UserID "
	sql += "left join [RYTreasureDB].[dbo].[OnLineOrder] as c on a.UserID = c.UserID "
	sql += "left join [RYRecordDB].[dbo].[RecordGrantTreasure] as d on a.UserID = d.UserID "
	sql += "where a.GameID = ? Group by a.GameID,a.Accounts,b.Score,b.InsureScore,a.LastLogonIP,a.RegisterIP"
	_, err = RYAccountsDB.Table(&data1).Query(sql, GameID)
	if err != nil {
		return data1,data2,err
	}
	sql1 := "select a.UserRealName as Name,a.BankCard,a.BankName FROM [RYAccountsDB].[dbo].[AccountsBankCard] as a inner join "
	sql1 += "[RYAccountsDB].[dbo].[AccountsInfo] as b on a.UserId = b.UserID where a.CardType = 0 and b.GameID = ?"
	_, err = RYAccountsDB.Table(&data2).Query(sql1, GameID)
	if err != nil {
		return data1,data2,err
	}
	return data1,data2,err
}

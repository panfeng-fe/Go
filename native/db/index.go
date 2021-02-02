package db

import (
	"native/model/app"
)

// SelectLoginParams ...
func SelectLoginParams(accounts string) (data app.DB_LoginParam,err error) {
	_, err = RYAccountsDB.Table(&data).Query("select UserID,Username,Password,Nullity from [RYPlatformManagerDB].[dbo].[Base_Users] where Username = ?", accounts)
	return data,err
}

// SelectIndex ...
func SelectIndex() (data []app.DB_Index,err error) {
	sql := "select count(case when UserID > 3000 then UserID end) as Month,"                  // 总人数
	sql += "count(case when DateDiff(dd,RegisterDate,getdate())=0 then UserID end) as Today " // 今日人数
	sql += "FROM [RYAccountsDB].[dbo].[AccountsInfo] union all "
	sql += "select count(*),count(case when DateDiff(dd,CollectDate,getdate())=0 then UserID end) " // 总代理与今日代理
	sql += "FROM [RYAccountsDB].[dbo].[AccountsAgent] union all "
	sql += "select sum(case when DateDiff(mm,BuyDate,getdate())=0 then TotalAmount else 0 end)," // 当月提现
	sql += "sum(case when DateDiff(dd,BuyDate,getdate())=0 then TotalAmount else 0 end) "        // 今日提现
	sql += "FROM [RYNativeWebDB].[dbo].[AwardOrder] where OrderStatus in (1,2) union all "
	sql += "select convert(int,sum(case when DateDiff(mm,ApplyDate,getdate())=0 then OrderAmount else 0 end))," // 当月充值
	sql += "convert(int,sum(case when DateDiff(dd,ApplyDate,getdate())=0 then OrderAmount else 0 end)) "        // 今日充值
	sql += "FROM [RYTreasureDB].[dbo].[ShareDetailInfo] union all "
	sql += "select sum(case when DateDiff(mm,CollectDate,getdate())=0 and Type = 1 then AddGold else 0 end)," // 当月赠送
	sql += "sum(case when DateDiff(dd,CollectDate,getdate())=0 and Type = 1 then AddGold else 0 end) "        // 今日赠送
	sql += "FROM [RYRecordDB].[dbo].[RecordGrantTreasure] union all "
	sql += "select sum(case when DateDiff(mm,CollectDate,getdate())=0 and Type = 2 then AddGold else 0 end)," // 当月彩金
	sql += "sum(case when DateDiff(dd,CollectDate,getdate())=0 and Type = 2 then AddGold else 0 end) "        // 今日彩金
	sql += "FROM [RYRecordDB].[dbo].[RecordGrantTreasure]"

	_, err = RYAccountsDB.Table(&data).Query(sql)
	return data,err
}

// OnlinePlayer ...
func OnlinePlayer() {
	//sql := ""
}

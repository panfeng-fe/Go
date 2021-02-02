package db

import (
	"native/model/app"
)

// ListInfo 订单信息
func ListInfo(page app.Page) (data []app.ListInfo, len int, err error) {
	var amount app.IntValue
	sql1 := "select top(?) a.OrderID,a.AwardPrice,a.TotalAmount,a.Compellation,CONVERT(VARCHAR(16),a.BuyDate,120) as Date1,"
	sql1 += "a.MobilePhone,a.DwellingPlace,a.BuyIP,b.GameID FROM [RYNativeWebDB].[dbo].[AwardOrder] as a left join "
	sql1 += "[RYAccountsDB].[dbo].[AccountsInfo] as b on a.UserID = b.UserID where a.OrderStatus = 0 and a.BuyDate "
	sql1 += "not in (select top(?) BuyDate FROM [RYNativeWebDB].[dbo].[AwardOrder] order by BuyDate desc) order by BuyDate desc"
	_, err = RYNativeWebDB.Table(&data).Query(sql1, page.Size, (page.Page-1)*page.Size)
	if err != nil {
		return nil, 0, err
	}
	sql2 := "select count(1) as Tatol from [RYNativeWebDB].[dbo].[AwardOrder] where OrderStatus = 0"
	_, err = RYNativeWebDB.Table(&amount).Query(sql2)
	if err != nil {
		return nil, 0, err
	}
	len = amount.Value
	return data, len, err
}

// ListFind 订单查询
func ListFind(GameID int) (data []app.ListInfo,  err error) {
	sql := "select a.OrderID,a.AwardPrice,a.TotalAmount,a.Compellation,CONVERT(VARCHAR(16),a.BuyDate,120) as Date,"
	sql += "a.MobilePhone,a.DwellingPlace,a.BuyIP,b.GameID FROM [RYNativeWebDB].[dbo].[AwardOrder] as a left join "
	sql += "[RYAccountsDB].[dbo].[AccountsInfo] as b on a.UserID = b.UserID where b.GameID = ? and a.OrderStatus = 0 order by BuyDate desc"
	_, err = RYNativeWebDB.Table(&data).Query(sql, GameID)
	return data, err
}

// ListMsg 订单未读信息
func ListMsg()(len int,err error){
	var amount app.IntValue
	sql := "select count(1) as msg from [RYNativeWebDB].[dbo].[AwardOrder] where OrderStatus = 0"
	_, err = RYNativeWebDB.Table(&amount).Query(sql)
	len = amount.Value
	return len,err
}

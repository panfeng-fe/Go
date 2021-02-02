package db

import (
	"native/model"
	"native/tools/log"
)

var (
	RYAccountsDB = model.RYAccountsDB()
	RYLotteryDB = model.RYLotteryDB()
	RYTreasureDB = model.RYTreasureDB()
	RYRecordDB = model.RYRecordDB()
	RYPlatformDB = model.RYPlatformDB()
	RYPlatformManagerDB = model.RYPlatformManagerDB()
	RYGameScoreDB = model.RYGameScoreDB()
	RYNativeWebDB = model.RYNativeWebDB()
	Log = log.Logger
)
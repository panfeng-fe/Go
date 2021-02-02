package app

// GameID ...
type GameID struct {
	GameID int
}

// StructUserID ...
type StructUserID struct {
	UserID int
}

type  GetLoginParams struct {
	UserID int
	UserName string
	Time string
	Signature string
	Type string
}

type PostLoginParams struct {
	UserID int
	UserName string
	Token string
	Random string
	LastLogin string
	LastIp string
	Channel string
	AgentName string
}

type  DBLoginInfo struct {
	UserID int
	GameID int
	DynamicPass string
}

type Context struct {
	State bool
	Msg string
	Content interface{}
}

// 数据库查询Int
type DBIntValue struct {
	IntValue int
}

// 数据库查询String
type DBStringValue struct {
	StringValue string
}

type UserLimit struct {// UserLimit 限红结构体
	Token string
	Random string
	Data string
	Member struct{
		Username string
	}
}
package app

// LoginParam ...
type LoginParam struct {
	Account  string
	Password string
}

// DB_LoginParam ...
type DB_LoginParam struct {
	UserID   int
	Username string
	Password string
	Nullity  int
}

// DB_Index ...
type DB_Index struct {
	Month int
	Today int
}

type IndexParam struct {
	AllPlayer       int
	TodayPlayer     int
	AllAgent        int
	TodayAgent      int
	MonthPresent    int
	TodayPresent    int
	MonthCaijin     int
	TodayCaijin     int
	MonthRecharge   int
	TodayRecharge   int
	MonthWithdrawal int
	TodayWithdrawal int
}

// Context ...
type Context struct {
	State   bool
	Msg     string
	Content interface{}
	Len     int
}

// Page ...
type Page struct {
	Size int
	Page int
}

type Player struct {
	UserID        int
	GameID        int
	Accounts      string
	Score         int
	InsureScore   int
	AgentGameID   int
	AgentName     string
	LastLogonIP   string
	RegisterIP    string
	RegisterDate  string
	LastLogonDate string
	Nullity       int
}

type IntValue struct {
	Value int
}

type StringValue struct {
	Value string
}

type UserInfo struct {
	Info     UserInfo_Info
	BankCard UserInfo_BankCard
}

type UserInfo_Info struct {
	GameID      int
	Accounts    string
	LastLogonIP string
	RegisterIP  string
	Socre       int
	InsureSocre int
	Recharge    int
	Withdrawal  int
	Present     int
}

type UserInfo_BankCard struct {
	Name     string
	BankCard string
	BankName string
}

type ListInfo struct {
	OrderID       int
	AwardPrice    int
	TotalAmount   int
	Compellation  string
	Data         string
	MobilePhone   string
	DwellingPlace string
	BuyIP         string
	GameID        int
}

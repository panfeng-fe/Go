package database

type Recorddgbetdetail struct {
	Id            int     `:"id"`
	LobbyId       int     `:"lobby_id"`
	TableId       int     `:"table_id"`
	ShoeId        int     `:"shoe_id"`
	PlayId        int     `:"play_id"`
	GameType      int     `:"game_type"`
	GameId        int     `:"game_id"`
	MemberId      int     `:"member_id"`
	BetTime       string  `:"bet_time"`
	CalTime       string  `:"cal_time"`
	WinOrLoss     float64 `:"win_or_loss"`
	WinOrLossz    float64 `:"win_or_lossz"`
	BetPoints     float64 `:"bet_points"`
	BetPointsz    float64 `:"bet_pointsz"`
	AvailableBet  float64 `:"available_bet"`
	UserName      string  `:"user_name"`
	Result        string  `:"result"`
	BetDetail     string  `:"bet_detail"`
	BetDetailz    string  `:"bet_detailz"`
	Ip            string  `:"ip"`
	Ext           string  `:"ext"`
	IsRevocation  int     `:"is_revocation"`
	ParentBetid   int     `:"parent_betid"`
	CurrencyId    int     `:"currency_id"`
	DeviceType    int     `:"device_type"`
	Pluginid      int     `:"pluginid"`
	ParentId      int     `:"parent_id"`
	BalanceBefore float64 `:"balance_before"`
	Serial        int     `:"serial"`
	RoadId        int     `:"road_id"`
}

type Recordcaipiaotoken struct {
	Id       int    `:"id"`
	UserId   int    `:"user_id"`
	UserName string `:"user_name"`
	GameId   int    `:"game_id"`
	Token    string `:"token"`
	Time     string `:"time"`
	LoginIp  string `:"login_ip"`
}

type Recordcaipiaotransfer struct {
	TransferId     int    `:"transfer_id"`     // 转账ID
	UserId         int    `:"user_id"`         // 用户ID
	Token          string `:"token"`           // 转账Token
	TicketId       int    `:"ticket_id"`       // 注单ID
	Data           string `:"data"`            // 转账流水号
	Score          int    `:"score"`           // 转账金币
	Channel        string `:"channel"`         // 转账渠道
	TransferStatus int    `:"transfer_status"` // 转账状态
	Revenue        int    `:"revenue"`         // 游戏税收
	TransferType   int    `:"transfer_type"`   // 转账类型
	InsertTime     string `:"insert_time"`     // ()
	IsDelete       int    `:"is_delete"`
	Balance        int    `:"balance"`
}
type Recordvideotransfer struct {
	TransferId     int    `:"transfer_id"`     // 转账ID
	UserId         int    `:"user_id"`         // 用户ID
	Token          string `:"token"`           // 转账Token
	TicketId       string `:"ticket_id"`       // 注单ID
	Data           string `:"data"`            // 转账流水号
	Score          int    `:"score"`           // 转账金币
	Channel        string `:"channel"`         // 转账渠道
	TransferStatus int    `:"transfer_status"` // 转账状态
	Revenue        int    `:"revenue"`         // 游戏税收
	TransferType   int    `:"transfer_type"`   // 转账类型
	InsertTime     string `:"insert_time"`
	IsDelete       int    `:"is_delete"`
	Balance        int    `:"balance"`
}

type Recordzqbetdetail struct {
	Id            int     `:"id"`
	LobbyId       int     `:"lobby_id"`
	TableId       int     `:"table_id"`
	Type          int     `:"type"`
	PlayId        int     `:"play_id"`
	GameType      int     `:"game_type"`
	GameId        int     `:"game_id"`
	MemberId      int     `:"member_id"`
	BetTime       string  `:"bet_time"`
	CalTime       string  `:"cal_time"`
	WinOrLoss     float64 `:"win_or_loss"`
	BetPoints     float64 `:"bet_points"`
	AvailableBet  float64 `:"available_bet"`
	UserName      string  `:"user_name"`
	Result        string  `:"result"`
	BetDetail     string  `:"bet_detail"`
	Ip            string  `:"ip"`
	Ext           string  `:"ext"`
	IsRevocation  int     `:"is_revocation"`
	ParentBetid   int     `:"parent_betid"`
	CurrencyId    int     `:"currency_id"`
	DeviceType    int     `:"device_type"`
	Pluginid      int     `:"pluginid"`
	ParentId      int     `:"parent_id"`
	BalanceBefore float64 `:"balance_before"`
	Serial        int     `:"serial"`
	RoadId        int     `:"road_id"`
}

type User struct {
	Id   int    `:"id"`
	Name string `:"name"`
	Age  int    `:"age"`
}

type Accountsinfo struct {
	UserID           int    `:"UserID"`          // 用户标识
	GameID           int    `:"GameID"`          // 游戏标识
	ProtectID        int    `:"ProtectID"`       // 密保标识
	PasswordID       int    `:"PasswordID"`      // 口令索引
	SpreaderID       int    `:"SpreaderID"`      // 推广员标识
	Accounts         string `:"Accounts"`        // 用户帐号
	NickName         string `:"NickName"`        // 用户昵称
	RegAccounts      string `:"RegAccounts"`     // 注册帐号
	UnderWrite       string `:"UnderWrite"`      // 个性签名
	PassPortID       string `:"PassPortID"`      // 身份证号
	Compellation     string `:"Compellation"`    // 真实名字
	LogonPass        string `:"LogonPass"`       // 登录密码
	InsurePass       string `:"InsurePass"`      // 安全密码
	DynamicPass      string `:"DynamicPass"`     // 动态密码
	DynamicPassTime  string `:"DynamicPassTime"` // 动态密码更新时间
	FaceID           int    `:"FaceID"`          // 头像标识
	CustomID         int    `:"CustomID"`        // 自定标识
	Present          int    `:"Present"`         // 赠送礼物
	UserMedal        int    `:"UserMedal"`       // 用户奖牌
	Experience       int    `:"Experience"`      // 经验数值
	GrowLevelID      int    `:"GrowLevelID"`
	LoveLiness       int    `:"LoveLiness"`       // 用户魅力
	UserRight        int    `:"UserRight"`        // 用户权限
	MasterRight      int    `:"MasterRight"`      // 管理权限
	ServiceRight     int    `:"ServiceRight"`     // 服务权限
	MasterOrder      int    `:"MasterOrder"`      // 管理等级
	MemberOrder      int    `:"MemberOrder"`      // 会员等级
	MemberOverDate   string `:"MemberOverDate"`   // 过期日期
	MemberSwitchDate string `:"MemberSwitchDate"` // 切换时间
	CustomFaceVer    int    `:"CustomFaceVer"`    // 头像版本
	Gender           int    `:"Gender"`           // 用户性别
	Nullity          int    `:"Nullity"`          // 禁止服务
	NullityOverDate  string `:"NullityOverDate"`  // 禁止时间
	StunDown         int    `:"StunDown"`         // 关闭标志
	MoorMachine      int    `:"MoorMachine"`      // 固定机器
	IsAndroid        int    `:"IsAndroid"`        // 是否机器人
	WebLogonTimes    int    `:"WebLogonTimes"`    // 登录次数
	GameLogonTimes   int    `:"GameLogonTimes"`   // 登录次数
	PlayTimeCount    int    `:"PlayTimeCount"`    // 游戏时间
	OnLineTimeCount  int    `:"OnLineTimeCount"`  // 在线时间
	LastLogonIP      string `:"LastLogonIP"`      // 登录地址
	LastLogonDate    string `:"LastLogonDate"`    // 登录时间
	LastLogonMobile  string `:"LastLogonMobile"`  // 登录手机
	LastLogonMachine string `:"LastLogonMachine"` // 登录机器
	RegisterIP       string `:"RegisterIP"`       // 注册地址
	RegisterDate     string `:"RegisterDate"`     // 注册时间
	RegisterMobile   string `:"RegisterMobile"`   // 注册手机
	RegisterMachine  string `:"RegisterMachine"`  // 注册机器
	RegisterOrigin   int    `:"RegisterOrigin"`   // PC       0x00     ,
	PlatformID       int    `:"PlatformID"`
	UserUin          string `:"UserUin"`
	RankID           int    `:"RankID"`
	AgentID          int    `:"AgentID"`
	RegisterChannel  int    `:"RegisterChannel"` // 用户渠道
}

type Recordcaipiaologo struct {
	Id        int    `:"id"`
	UserId    int    `:"user_id"` // 用户ID
	UserName  string `:"user_name"`
	Token     string `:"token"`      // 登录Token
	LoginTime string `:"login_time"` // 登录时间
	LoginIp   string `:"login_ip"`   // 登录IP
	Channel   string `:"channel"`    // 登录彩票服务商
	DateId    int    `:"date_id"`    // 日期ID
}

type Accountsvideologo struct {
	Id int `:"id"`
	UserId int `:"user_id"` // 用户ID
	Token string `:"token"` // 视讯登录令牌

	Random string `:"random"` // 随机数

	LastLogin string `:"last_login"` // 最后登录时间
	LastIp string `:"last_ip"` // 最后登录IP
	Channel string `:"channel"` // 视讯服务渠道

	AgentName string `:"agent_name"`
}

type Recordvideologo struct {
	Id int `:"id"`
	UserId int `:"user_id"`
	Token string `:"token"`
	UserName string `:"user_name"`

	LoginTime string `:"login_time"`
	LoginIp string `:"login_ip"`
	Channel string `:"channel"`

	DateId int `:"date_id"`
}


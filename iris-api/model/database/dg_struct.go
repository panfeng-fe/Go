package database

// Member 成员
type Member struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Status   int     `json:"status"`
	Currency string  `json:"currency"`
	WinLimit float64 `json:"winLimit"`
	Amount   float64 `json:"amount"`
}

// Order 转账记录对象
type Order struct {
	Username string  `json:"username"`
	TicketID int     `json:"ticketId"`
	Serial   string  `json:"serial"`
	Amount   float64 `json:"amount"`
}

// BalanceRequest 余额请求
type BalanceRequest struct {
	Token  string `json:"token"`
	Member Member `json:"member"`
}

// BalanceResponse 余额请求返回
type BalanceResponse struct {
	CodeID int    `json:"codeId"`
	Token  string `json:"token"`
	Member Member `json:"member"`
}

// TransferRequest 转账请求
type TransferRequest struct {
	Token    string `json:"token"`
	TicketID int    `json:"ticketId"`
	Data     string `json:"data"`
	Member   Member `json:"member"`
}

// TransferResponse 转账请求返回
type TransferResponse struct {
	CodeID int    `json:"codeId"`
	Token  string `json:"token"`
	Data   string `json:"data"`
	Member Member `json:"member"`
}

// CheckTransferRequest 确认存取款结果接口
type CheckTransferRequest struct {
	Token string `json:"token"`
	Data  string `json:"data"`
}

// CheckTransferResponse 确认存取款结果接口返回
type CheckTransferResponse struct {
	Token  string `json:"token"`
	CodeID int    `json:"codeId"`
}

// InformRequest 请求回滚转账事务
type InformRequest struct {
	Token    string `json:"token"`
	TicketID int    `json:"ticketId"`
	Data     string `json:"data"`
	Member   Member `json:"member"`
}

// InformResponse 请求回滚转账事务返回
type InformResponse struct {
	CodeID int    `json:"codeId"`
	Token  string `json:"token"`
	Data   string `json:"data"`
	Member Member `json:"member"`
}

// OrderRequest 请求对账接口
type OrderRequest struct {
	Token    string `json:"token"`
	TicketID int    `json:"ticketId"`
}

// OrderResponse 请求对账接口返回
type OrderResponse struct {
	CodeID   int     `json:"codeId"`
	Token    string  `json:"token"`
	TicketID int     `json:"ticketId"`
	List     []Order `json:"list"`
}

// UnsettleRequest 查询未派彩和未回滚记录
type UnsettleRequest struct {
	Token string   `json:"token"`
	List  []string `json:"list"`
}

// UnsettleResponse 查询未派彩和未回滚记录返回
type UnsettleResponse struct {
	CodeID int     `json:"codeId"`
	Token  string  `json:"token"`
	List   []Order `json:"list"`
}

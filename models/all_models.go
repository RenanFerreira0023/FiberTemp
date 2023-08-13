package models

import "time"

var LayoutDate = "2006-01-02 15:04:05"

type MiddlewareStruct struct {
	NumRequests int
	LastRequest time.Time
}
type MessageError struct {
	MsgError string `json:"message_error"`
}

type QueryGetUsersAgent struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	SecondName     string    `json:"second_name"`
	Email          string    `json:"email"`
	PasswordAgent  string    `json:"password_agent"`
	CreateAccount  time.Time `json:"dt_create_account"`
	ExpiredAccount time.Time `json:"dt_expired_account"`
	AccountValid   bool      `json:"account_valid"`
	QuantityAlert  int       `json:"quantity_alerts"`
	AccountCopy    int       `json:"quantity_account_copy"`
}

type QueryBodyUsersAgent struct {
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	Email          string `json:"email"`
	Password_Agent string `json:"password_agent"`
	CreateAccount  string `json:"dt_create_account"`
	ExpiredAccount string `json:"dt_expired_account"`
	AccountValid   bool   `json:"account_valid"`
	QuantityAlert  int    `json:"quantity_alerts"`
	AccountCopy    int    `json:"quantity_account_copy"`
}

type QueryBodyCreateChannel struct {
	NameChannel   string `json:"channel_name"`
	AgentID       int    `json:"users_agent_id"`
	CreateChannel string `json:"dt_create_channel"`
}

type QueryGetLogin struct {
	AgentID       int `json:"agent_id"`
	ChannelID     int `json:"channel_id"`
	QuantityAlert int `json:"quantity_alerts"`
	AccountCopy   int `json:"quantity_account_copy"`
}

type QueryBodySendCopy struct {
	Symbol        string  `json:"symbol"`
	ActionType    string  `json:"action_type"`
	Ticket        int64   `json:"ticket"`
	Lot           float64 `json:"lot"`
	TargetPedding float64 `json:"target_pedding"`
	TakeProfit    float64 `json:"takeprofit"`
	StopLoss      float64 `json:"stoploss"`
	DateEntry     string  `json:"dt_send_order"`
	UserAgentID   int     `json:"user_agent_id"`
	ChannelID     int     `json:"channel_id"`
}

type QueryBodyUserReceptor struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	SecondName     string    `json:"second_name"`
	Email          string    `json:"email"`
	CreateAccount  time.Time `json:"dt_create_account"`
	ExpiredAccount time.Time `json:"dt_expired_account"`
}

type QueryGetUserReceptor struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	Email          string `json:"email"`
	CreateAccount  string `json:"dt_create_account"`
	ExpiredAccount string `json:"dt_expired_account"`
}

type QueryGetValidReceptor struct {
	ID             int       `json:"id"`
	ExpiredAccount time.Time `json:"dt_create_account"`
}

type QueryGetChannel struct {
	ID      int `json:"id"`
	AgentID int `json:"users_agent_id"`
}

type QueryRequestLoginReceptor struct {
	AgentID    int `json:"user_agent_id"`
	ReceptorID int `json:"user_receptor_id"`
	ChannelID  int `json:"channel_id"`
}

type BodyCopyTrader struct {
	ID            int       `json:"id"`
	Symbol        string    `json:"symbol"`
	ActionType    string    `json:"action_type"`
	Ticket        int       `json:"ticket"`
	Lot           float64   `json:"lot"`
	TargetPedding float64   `json:"target_pedding"`
	TakeProfit    float64   `json:"takeprofit"`
	StopLoss      float64   `json:"stoploss"`
	DateEntry     time.Time `json:"dt_send_order"`
	AgentID       int       `json:"user_agent_id"`
	ChannelID     int       `json:"channel_id"`
}

type StrutcURLCopyTrader struct {
	AgentID   int    `json:"user_agent_id"`
	ChannelID int    `json:"channel_id"`
	DateStart string `json:"dt_start"`
	DateEnd   string `json:"dt_end"`
	Offset    int    `json:"offset"`
	PageLimit int    `json:"limit"`
}

type QueryRequestReqCopy struct {
	ID            int    `json:"id"`
	DateSendOrder string `json:"dt_send_copy"`
	AllCopyID     int    `json:"all_copy_id"`
	ReceptorID    int    `json:"users_receptor_id"`
	ChannelID     int    `json:"channel_id"`
}

type QueryBodyInsertPermission struct {
	ID             int `json:"id"`
	UserReceptorID int `json:"user_receptor_id"`
	ChannelID      int `json:"channel_id"`
}

type QueryRequestToken struct {
	UserID int    `json:"id"`
	Token  string `json:"token"`
}

type JsonRequest200 struct {
	DataBaseID int    `json:"id"`
	MsgInsert  string `json:"message_insert"`
}

type StrutcURLGetChannelList struct {
	AgentID   int    `json:"user_agent_id"`
	DateStart string `json:"dt_start"`
	DateEnd   string `json:"dt_end"`
	Offset    int    `json:"offset"`
	PageLimit int    `json:"limit"`
}

type RequestChannelList struct {
	ID          int       `json:"id"`
	AgentID     int       `json:"user_agent_id"`
	ChannelName string    `json:"channel_name"`
	DateCreate  time.Time `json:"dt_create_channel"`
}

type BodyPostLoginAdm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type BodyDelete struct {
	ID      int `json:"id"`
	AgentID int `json:"user_agent_id"`
}

type BodyUpdate struct {
	ID             int    `json:"id"`
	AgentID        int    `json:"user_agent_id"`
	NewNameChannel string `json:"channel_name"`
}

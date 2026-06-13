package domain

type Inbound struct {
	ID          int          `json:"id"`
	Up          int64        `json:"up"`
	Down        int64        `json:"down"`
	Total       int64        `json:"total"`
	Remark      string       `json:"remark"`
	Enable      bool         `json:"enable"`
	Protocol    string       `json:"protocol"`
	Port        int          `json:"port"`
	ClientStats []ClientStat `json:"clientStats"`
}

type ClientStat struct {
	ID         int    `json:"id"`
	InboundID  int    `json:"inboundId"`
	Enable     bool   `json:"enable"`
	Email      string `json:"email"`
	Up         int64  `json:"up"`
	Down       int64  `json:"down"`
	Total      int64  `json:"total"`
	ExpiryTime int64  `json:"expiryTime"`
}

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     T      `json:"obj"`
}

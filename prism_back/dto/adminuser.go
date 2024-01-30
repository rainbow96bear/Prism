package dto

type AdminLoginResult struct {
	Result    bool      `json:"result"`
	IsAdmin   bool      `json:"isAdmin"`
	AdminInfo AdminInfo `json:"admin_info"`
}

type AdminInfo struct {
	Id   string `json:"admin_id"`
	Rank int    `json:"admin_rank"`
}
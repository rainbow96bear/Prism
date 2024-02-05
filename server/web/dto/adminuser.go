package dto

type AdminAuthorization struct {
	IsAdmin   bool      `json:"isAdmin"`
	AdminInfo AdminInfo `json:"admin_info"`
}

type AdminInfo struct {
	Id   string `json:"id"`
	Rank int    `json:"rank"`
}
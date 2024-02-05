package models

type AdminUser struct {
	Id       string `json:"id"`
	Rank     int    `json:"Admin_rank,omitempty"`
	Password string `json:"password,omitempty"`
}
package dto

type User struct {
	Id string `json:"id"`
}

type UserTech struct {
	Name  string `json:"tech_name"`
	Level int    `json:"level"`
}

type TechList struct {
	UserTechList []UserTech `json:"tech_list"`
}
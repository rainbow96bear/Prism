package dto

type Tech struct {
	Id    int    `json:"tech_code,omitempty"`
	Name  string `json:"tech_name"`
	Count int    `json:"count"`
}

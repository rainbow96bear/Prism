package dto

type Tech struct {
	Name  string `json:"tech_name"`
	Count int    `json:"count,omitempty"`
}
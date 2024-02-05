package dto

type Personaldata struct {
	Id                 string   `json:"id"`
	Nickname           string   `json:"nickname"`
	One_line_introduce string   `json:"oneLineIntroduce,omitempty"`
	Hashtag            []string `json:"hashtag,omitempty"`
}

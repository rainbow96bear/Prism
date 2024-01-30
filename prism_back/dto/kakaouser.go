package dto

type KakaoUser struct {
	Id            string `json:"sub"`
	Nickname      string `json:"nickname,omitempty"`
	ProfileImgURL string `json:"picture,omitempty"`
}
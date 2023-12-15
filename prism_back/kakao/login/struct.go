package Login

type KakaoID struct {
	isHave bool
	Value  int
}
type Token struct {
	Access_token             string `json:"access_token"`
	Token_type               string `json:"token_type"`
	Refresh_token            string `json:"refresh_token"`
	Id_token                 string `json:"id_token"`
	Expires_in               int    `json:"expires_in"`
	Scope                    string `json:"scope"`
	Refresh_token_expires_in int    `json:"refresh_token_expires_in"`
}

// kakao user 정보
type User struct {
	User_id     string `json:"sub"`
	Nickname    string `json:"nickname,omitempty"`
	Profile_img string `json:"picture,omitempty"`
}

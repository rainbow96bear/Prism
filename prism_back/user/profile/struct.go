package Profile

type Profile struct {
	Id                 int    `json:"Id"`
	One_line_Introduce string `json:"One_Line_Introduce,omitempty"`
	User_id            string `json:"user_info_User_id"`
}
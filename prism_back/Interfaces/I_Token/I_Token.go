package I_Token

import "net/http"


func GetToken(t I_Token, res http.ResponseWriter, req *http.Request) (I_Token, error) {
	token, err := t.GetToken(res, req)
	if err != nil {
		return nil, err
	}
	return token, err
}


type I_Token interface {
	GetToken(res http.ResponseWriter, req *http.Request) (I_Token, error)
}

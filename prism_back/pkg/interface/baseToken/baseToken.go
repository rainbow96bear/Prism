package basetoken

import "net/http"

type BaseToken interface {
	GetToken(res http.ResponseWriter, req *http.Request) (BaseToken BaseToken, err error)
}

func GetToken(t BaseToken, res http.ResponseWriter, req *http.Request) (BaseToken BaseToken, err error) {
	token, err := t.GetToken(res, req)
	if err != nil {
		return nil, err
	}
	return token, err
}
package common

import (
	"my/util"
	"net/http"
	"sso/user"
)

func Exit(res http.ResponseWriter, req *http.Request) {
	token, _ := req.Cookie("token")
	if token == nil {
		util.WriteJSON(res, "OK")

		return
	}
	state := new(user.State)
	state.Token = token.Value
	state.Del()
	util.WriteJSON(res, "OK")
}

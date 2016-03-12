package common

import (
	"my/util"
	"net/http"
	"sso/user"
	"strconv"
)

// 修改用户
func Edit(rep http.ResponseWriter, req *http.Request) {

	var token string
	cookie, err := req.Cookie("token")
	if err != nil {
		token = req.FormValue("token")
		if token == "" {
			util.WriteJSON(rep, "1")
			return
		}
	} else {
		token = cookie.Value
	}

	u := byToken(token)
	if u == nil {
		util.WriteJSON(rep, "2")
		return
	}
	// conf = util.GetConfig()
	alias := req.FormValue("alias")
	mobile := req.FormValue("mobile")
	wechat := req.FormValue("wechat")
	username := req.FormValue("username")
	password := req.FormValue("password")
	alipay := req.FormValue("alipay")
	qq := req.FormValue("qq")
	email := req.FormValue("email")
	city := req.FormValue("city")

	isOk := util.RegexpMobile.MatchString(mobile)
	if !isOk {
		util.WriteJSON(rep, "3")
		return
	}

	isOk = util.RegexpCommon.MatchString(wechat)
	if !isOk {
		util.WriteJSON(rep, "4")
		return
	}
	alias = util.RegexpFileter.ReplaceAllString(alias, "")

	userRef := new(user.User)
	if password != "" {
		isOk = util.RegexpPassword.MatchString(password)
		if !isOk {
			util.WriteJSON(rep, "5")
			return
		}
		userRef.Password = util.Md5Encode(password)
	}

	userRef.Id = u.Id
	userRef.Alias = alias
	userRef.Mobile = mobile
	userRef.Wechat = wechat
	userRef.Username = username
	userRef.Alipay = alipay
	userRef.QQ, _ = strconv.ParseInt(qq, 10, 64)
	userRef.Email = email
	userRef.City = city
	userRef.Update(*userRef)
	util.WriteJSON(rep, "OK")
}

package common

import (
	"my/util"
	"net/http"
	"sso/user"
	"sync"
	"time"
)

var lockSignin sync.Mutex

// 注册用户
func Signin(w http.ResponseWriter, r *http.Request) {
	lockSignin.Lock()
	alias := r.FormValue("alias")
	mobile := r.FormValue("mobile")
	password := r.FormValue("password")
	wechat := r.FormValue("wechat")
	////validate
	isOk := util.RegexpMobile.MatchString(mobile)
	if !isOk {
		util.WriteJSON(w, "moblie is error.")
		return
	}
	isOk = util.RegexpPassword.MatchString(password)
	if !isOk {
		util.WriteJSON(w, "password is error.")
		return
	}
	isOk = util.RegexpPassword.MatchString(wechat)
	if !isOk {
		util.WriteJSON(w, "wechat is error.")
		return
	}
	alias = util.RegexpFileter.ReplaceAllString(alias, "")
	//
	var userB = &user.User{}
	userB.Mobile = mobile
	userB = userB.ByMobile()
	if userB != nil {
		util.WriteJSON(w, userB)
		return
	}

	var user user.User
	user.Mobile = mobile
	user.Alias = alias
	user.Password = util.Md5Encode(password)
	user.Wechat = wechat
	user.Create = time.Now()
	user.Last = user.Create
	userOk, _, _ := user.Add()
	util.WriteJSON(w, userOk)
	lockSignin.Unlock()
}

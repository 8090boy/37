package controllers

import (
	"encoding/json"
	model "interaction/models"
	"my/util"
	"net/http"
	"net/url"
	"sso/user"
	"strconv"
	"sync"
	"time"
)

var lockSignin sync.Mutex

// No reference, cannot be registered
func Signin(w http.ResponseWriter, r *http.Request) {
	lockSignin.Lock()
	defer lockSignin.Unlock()
	referrerId := r.FormValue("referrerId")
	alias := r.FormValue("alias")
	mobile := r.FormValue("mobile")
	password := r.FormValue("password")
	wechat := r.FormValue("wechat")

	isOk := util.RegexpMobile.MatchString(mobile)
	if !isOk {
		util.WriteJSON(w, 1)
		return
	}
	isOk = util.RegexpPassword.MatchString(password)
	if !isOk {
		util.WriteJSON(w, 2)
		return
	}
	isOk = util.RegexpCommon.MatchString(wechat)
	if !isOk {
		util.WriteJSON(w, 3)
		return
	}
	alias = util.RegexpFileter.ReplaceAllString(alias, "")

	//
	conf = util.GetConfig()
	if referrerId == "" {
		util.WriteJSON(w, 4)
		return
	}

	state := new(user.State).Find(referrerId)
	if state == nil {
		util.WriteJSON(w, 5)
		return
	}
	// 推荐码保质期
	recommandExpiration := conf.Get("common", "recommandExpiration")
	langth, _ := time.ParseDuration(recommandExpiration)

	if time.Now().Local().After(state.Overdue.Add(langth)) {
		state.Del()
		util.WriteJSON(w, 6)
		return
	}

	refUser := findUserById(state.Userjson)

	//Have a reference
	if refUser.Id < 1 {
		state.Del() // 删除推荐码,确保一次性
		util.WriteJSON(w, 7)
		return
	}

	refRelational := &model.Relational{}
	refRelational = refRelational.BySsoId(refUser.Id)
	if refRelational == nil {
		util.WriteJSON(w, 8)
		return
	}
	// 推荐人处于非正常状态
	if refRelational.Status != 1 {
		util.WriteJSON(w, 9)
		return
	}
	//My sso account
	myUser := findUserByMob(mobile)
	if myUser.Id > 0 { //  I have sso account
		relational := &model.Relational{}
		relational = relational.ByMob(mobile)
		if relational != nil {
			if relational.Id > 0 {
				// 用户存在
				util.WriteJSON(w, 10)
				return
			}
		}
		state.Del() // 删除推荐码,确保一次性
		// Have a sso account, there is no account number 37
		// Increase the 37 account
		joinRelational(myUser, refRelational)
		util.WriteJSON(w, "ok")
		return
	}
	// Without sso account
	postValues := url.Values{}
	postValues.Add("mobile", mobile)
	postValues.Add("wechat", wechat)
	postValues.Add("alias", alias)
	postValues.Add("password", password)
	ssoUrl := conf.Get("sso", "url")
	tmpUser, _ := util.Post(ssoUrl+"/reg", postValues)
	err := json.Unmarshal(tmpUser, myUser)

	if err != nil {
		util.WriteJSON(w, 11)
		return
	}
	state.Del() // 删除推荐码,确保一次性
	// Increase the 37 account
	joinRelational(myUser, refRelational)
	util.WriteJSON(w, "ok")
}

func joinRelational(user *user.User, refUser *model.Relational) bool {
	if user == nil {
		return false
	}

	// Increase the 37 account
	relational := model.Relational{}
	relational.SsoId = user.Id
	relational.Mobile = user.Mobile
	relational.Referrer = strconv.FormatInt(refUser.Id, 10)
	relational.Create = time.Now()
	relational.Prev = relational.Create
	_, err := relational.Add()
	if err != nil {
		relational.Edit()
		if err != nil {
			//			fmt.Println(err)
			return false
		}
	}
	return true
}

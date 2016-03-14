package common

import (
	"io"
	"my/util"
	"net/http"
	"sso/user"
	"strings"
)

func ByIds(res http.ResponseWriter, req *http.Request) {
	idsStr := req.URL.Query().Get("ids")
	ids := strings.Split(idsStr, "|")
	users := make([]*user.User, len(ids))
	for i := 0; i < len(ids); i++ {
		userTmp := new(user.User)
		userTmp = userTmp.ById(ids[i])
		if userTmp != nil {
			if userTmp.Id > 0 {
				users[i] = userTmp
			}
		}
	}
	util.WriteJSON(res, users)

}
func Find(res http.ResponseWriter, req *http.Request) {
	//来路校验 目前只有37互动的ip
	key := req.FormValue("k")
	val := req.FormValue("v")
	if key == "" || val == "" {
		util.WriteJSON(res, "")
		return
	}
	user := &user.User{}
	switch strings.ToLower(key) {
	case "mob": //mobile
		user = findMob(val)
	case "em": //email
		user.Email = val
		user = user.ByEmail()
	case "id":

		user = user.ById(val)
	}

	if user == nil {
		util.WriteJSON(res, "")
		return
	}
	if user.Id < 1 {
		util.WriteJSON(res, "")
		return
	}
	util.WriteJSON(res, user)
}

func ByToken(res http.ResponseWriter, req *http.Request) {

	token := req.Header.Get("token")
	if token == "" {
		io.WriteString(res, "")
		return
	}
	state := &user.State{}
	state = state.Find(token)

	if state == nil {
		res.Write([]byte(""))
		return
	}
	mobArr := strings.Split(state.Userjson, "_")
	mob := mobArr[0]
	userA := findMob(mob)
	util.WriteJSON(res, userA)
	state.DelOverdue(state.Userjson)
}

func byToken(token string) *user.User {
	state := new(user.State)
	state = state.Find(token)
	if state == nil {
		return nil
	}
	mobArr := strings.Split(state.Userjson, "_")
	mob := mobArr[0]
	userA := findMob(mob)
	return userA
}

func findMob(mobile string) *user.User {
	isOk := util.RegexpMobile.MatchString(mobile)
	if !isOk {
		return nil
	}
	//根据手机查找用户

	user := new(user.User)
	user.Mobile = mobile
	user = user.ByMobile()
	if user == nil {
		return nil
	}

	return user
}

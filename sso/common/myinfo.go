package common

import (
	"encoding/json"
	"my/util"
	"net/http"
	"regexp"
	"sso/user"
)

// get my infomation
func Myinfo(res http.ResponseWriter, req *http.Request) {

	callback := req.FormValue("cb")
	var sweet map[string]interface{} = make(map[string]interface{})
	sweet["state"] = 0
	// 特殊字符替换
	reg := regexp.MustCompile(`"|'|<|\s`)
	callback = reg.ReplaceAllString(callback, "")
	if len(callback) > 10 {
		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}
	//my user info
	stat, userA := validateUserInfo(req)

	if stat == false {
		sweet["state"] = 1
		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}
	sweet["u"] = *userA
	sweet["state"] = 2
	all_info, _ := json.Marshal(sweet)
	if callback == "null" {
		util.WriteJSONP(res, string(all_info))
		return
	}

	util.WriteJSONP(res, callback+"("+string(all_info)+")")
	return
}

// 校验用户请求信息
func validateUserInfo(req *http.Request) (bool, *user.User) {
	cookie := req.URL.Query().Get("token")

	if cookie == "" {
		return false, nil
	}
	userA := byToken(cookie)
	//my user info
	if userA == nil {
		return false, nil
	}
	if userA.Id == 0 {
		return false, nil
	}

	return true, userA
}

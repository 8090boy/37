package controllers

import (
	"encoding/json"

	"my/util"
	"net/http"
	"regexp"
)

// get my infomation
func Myinfo(res http.ResponseWriter, req *http.Request) {
	conf = GetConfig()
	callback := req.FormValue("cb")
	var sweet map[string]interface{} = make(map[string]interface{})
	sweet["state"] = 0
	// 特殊字符替换
	reg := regexp.MustCompile(`"|'|<|\s`)
	callback = reg.ReplaceAllString(callback, "")
	if len(callback) > 2 {
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
	util.WriteJSONP(res, callback+"("+string(all_info)+")")
	return
}

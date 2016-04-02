package controllers

import (
	model "interaction/models"
	"my/util"
	"sso/user"
	"strconv"
	"strings"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

// 朋友圈
func Friendster(rep rest.ResponseWriter, req *rest.Request) {
	token, err := req.Cookie("token")
	if err != nil || token.Value == "" {
		rep.WriteJson("1")
		return
	}
	user := new(user.User)
	byToken(token.Value, user)
	if user == nil {
		rep.WriteJson("2")
		return
	}

	myRela := new(model.Relational).FindBySsoId(user.Id)
	myFirstRecommended, mySecondRecommended := findRelationalsForRelationalId(myRela.Id)
	result := make(map[string]string)
	tmpA := ""
	tmpB := ""
	for _, fr := range myFirstRecommended {

		if fr == nil {
			continue
		}
		if fr.Id > 0 {
			fId := strconv.FormatInt(fr.Id, 10)
			tmpA = tmpA + "|" + fId + "-" + fr.Mobile
		}
	}

	for _, frr := range mySecondRecommended {

		if frr == nil {
			continue
		}
		if frr.Id > 0 {
			tmpB = tmpB + "|" + frr.Referrer + "-" + frr.Mobile
		}
	}
	result["f"] = tmpA
	result["s"] = tmpB
	rep.WriteJson(result)
}

func findRelationalsForRelationalId(relasid int64) ([]*model.Relational, []*model.Relational) {
	relaTmp := new(model.Relational)

	firstLRelas, firstLRCount := relaTmp.FindRecommended(relasid)
	if firstLRCount == 0 {
		return nil, nil
	}
	seconedRelas := make([]*model.Relational, len(firstLRelas))
	for _, r := range firstLRelas {

		if r.Id == 0 {
			continue
		}

		tmpRelas, count := relaTmp.FindRecommended(r.Id)
		if count > 0 {
			for _, rr := range tmpRelas {
				if rr.Id > 0 {
					seconedRelas = append(seconedRelas, rr)
				}
			}
		}
	}

	return firstLRelas, seconedRelas
}

// 获取我的关系户
func MyRelationUser(rep rest.ResponseWriter, req *rest.Request) {
	mob := req.PathParam("mob")

	if mob == "" {
		rep.WriteJson("")
		return
	}

	refUuser := findUserByMob(mob)
	myInfo := findUser(req)
	if refUuser.Id == 0 {
		rep.WriteJson("")
		return
	}

	if refUuser == nil || myInfo == nil {
		rep.WriteJson("")
		return
	}
	myRela := new(model.Relational).BySsoId(myInfo.Id)
	refRela := new(model.Relational).ByMob(mob)
	if refRela == nil || myRela == nil {
		rep.WriteJson("")
		return
	}
	myRelaIdStr := strconv.FormatInt(myRela.Id, 10)

	if myRelaIdStr == refRela.Referrer {
		rep.WriteJson(refUuser)
		return
	}
	refId, _ := strconv.ParseInt(refRela.Referrer, 10, 64)
	refTwoLayerRela := myRela.ById(refId)
	if refTwoLayerRela == nil {
		rep.WriteJson("")
		return
	}
	if myRelaIdStr == refTwoLayerRela.Referrer {
		rep.WriteJson(refUuser)
		return
	}

	rep.WriteJson("")
	return

}

// 邀请的动态码
func InvitationCode(rep rest.ResponseWriter, req *rest.Request) {
	conf = GetConfig()
	userA := findUser(req)
	result := ""
	if userA == nil {
		rep.WriteJson(result)
		return
	}
	ref := util.Rand().Hex()
	result = strings.Split(ref, "-")[0]

	state := new(user.State)
	day := conf.Get("common", "recommandExpiration")
	refDate, _ := time.ParseDuration(day)
	sysNow := time.Now().Local()
	state.Overdue = sysNow.Add(refDate)
	state.Token = result
	state.Userjson = strconv.FormatInt(userA.Id, 10)
	state.Add()
	state.DelMore()
	rep.WriteJson(result)

}

// 修改自己信息
func Edit(rep rest.ResponseWriter, req *rest.Request) {
	token, err := req.Cookie("token")
	if err != nil || token.Value == "" {
		rep.WriteJson("1")
		return
	}
	user := new(user.User)
	byToken(token.Value, user)
	if user == nil {
		rep.WriteJson("2")
		return
	}
	ssoUrl := conf.Get("sso", "url")
	editUrl := conf.Get("sso", "edit")
	url := ssoUrl + editUrl

	rela := new(model.Relational).FindBySsoId(user.Id)
	req.ParseForm()
	if rela != nil && rela.Id > 0 {
		freetime := req.FormValue("freetime")
		freetime = util.RegexpFileter.ReplaceAllString(freetime, "")
		rela.Freetime = freetime
		rela.Edit()
	}
	req.PostForm.Add("token", token.Value)
	util.Post(url, req.PostForm)
	rep.WriteJson("OK")
}

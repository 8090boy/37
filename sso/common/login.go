package common

import (
	"my/util"
	"net/http"
	"sso/user"
	"strconv"
	"strings"
	"time"
)

const (
	FailUrl  = "/failed.html"
	IndexUrl = "/index.html"
)

var conf *util.Config = nil

func SetConfig(filepath string) *util.Config {
	if conf != nil {
		return conf
	}
	conf = new(util.Config)
	conf.Filepath = filepath
	return conf
}

func GetConfig() *util.Config {
	return conf
}

// 找回密码
func Restore(res http.ResponseWriter, req *http.Request) {
	toEmail := req.FormValue("email")
	isok := util.RegexpEmail.MatchString(toEmail)
	if !isok {
		util.WriteJSON(res, "1")
		return
	}
	// 查找用户
	userA := new(user.User)
	userA.Email = toEmail
	userA = userA.ByEmail()
	if userA == nil {
		util.WriteJSON(res, "2")
		return
	}

	// 创建邮件
	tmpStr := userA.Mobile
	if userA.Alias != "" {
		tmpStr = userA.Alias
	}

	if sendMail(toEmail, tmpStr) != nil {
		util.WriteJSON(res, "fail")
		return
	}

	util.WriteJSON(res, "ok")
}

func sendMail(to, username string) error {
	// 获取配置
	conf = GetConfig()
	host := conf.Get("email", "serveraddr")
	user := conf.Get("email", "user")
	password := conf.Get("email", "password")
	datetime := time.Now().Format(util.TimeFormat)
	//
	randCode := util.Rand()
	serverName := conf.Get("37client", "serverName")
	url := serverName + "?t=" + randCode.Hex()
	subject := "3737.io 密码重置"
	//
	body := `
<includetail>
<table width="80%" border="0" cellspacing="0" cellpadding="0" align="center" style="background:#FF9800">
<tbody>
<tr>
<td  style="  color:#ffffff; font-size:2em;">尊敬的3737互动会员：` + username + ` <hr>
<span>您在` + datetime + `提交的找回密码，点击这里<a href=\"` + url + `\" target="_blank" style="color:#fff; font-size:1.4em;  text-decoration:underline">密码重置</a>！</span>                    </td>
</tr>
</tbody>
</table>
</includetail>
`
	err := util.SendToMail(user, password, host, to, subject, body)
	return err
}

// 登录
func Login(res http.ResponseWriter, req *http.Request) {
	conf = GetConfig()
	// urlRef := conf.Get("37client", "serverName") + conf.Get("37client", "index")
	//loginIndexUrl := conf.Get("37client", "protocol") + urlRef
	loginOverdue := conf.Get("sysinfo", "loginOverdue")
	//
	if req.Method != "POST" {
		util.WriteJSON(res, 1)
		//http.Redirect(res, req, loginIndexUrl, http.StatusMovedPermanently)
		return
	}
	redirectUrl := req.FormValue("redirecturl")
	if redirectUrl == "" {
		util.WriteJSON(res, 2)
		//		http.Redirect(res, req, loginIndexUrl, http.StatusMovedPermanently)
		return
	}
	tmpStrRef := strings.Split(redirectUrl, "?msg=")
	if len(tmpStrRef) > 1 {
		redirectUrl = tmpStrRef[0]
	}

	refName := req.FormValue("username")
	password := req.FormValue("password")
	password = util.Md5Encode(password)
	//数据安全校验、过滤
	mobOk := util.RegexpMobile.MatchString(refName)
	emailOk := util.RegexpEmail.MatchString(refName)
	//
	userA := new(user.User)
	if mobOk {
		userA.Mobile = refName
		userA = userA.ByMobile()
	} else if emailOk {
		userA.Email = refName
		userA = userA.ByEmail()
	} else {
		userA.Username = refName
		userA = userA.ByUsername()
	}
	// 用户不存在
	if userA == nil || userA.Id == 0 {
		util.WriteJSON(res, 3)
		return
	}
	// 密码不对
	if userA.Password != password {
		util.WriteJSON(res, 4)
		return
	}
	//
	userjson := userA.Mobile + "_" + strconv.FormatInt(userA.Id, 10)
	var ref util.UUID = util.Rand()
	token := ref.Hex()
	//记录登录用户id
	state := &user.State{}
	state.DelOverdue(userjson)
	state.Token = token
	state.Userjson = userjson

	refTime, _ := time.ParseDuration(loginOverdue)
	state.Overdue = time.Now().Add(refTime)
	state.Add()
	//
	loginOverdueWeb := conf.Get("sysinfo", "loginOverdueWeb")
	//OkUrl := urlRef + "?token=" + state.Token + "&url=" + redirectUrl
	//http.Redirect(res, req, OkUrl, http.StatusMovedPermanently)
	refObj := "{status:'ok',token:'" + state.Token + "',overdue:'" + loginOverdueWeb + "'}"
	util.WriteJSON(res, refObj)
}

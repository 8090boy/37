package controllers

import (
	"net/http"
	"strings"
)

// Must be sso request and provide information about redirectUrl and token
func Login(res http.ResponseWriter, req *http.Request) {
	conf = GetConfig()
	token := req.URL.Query().Get("token")
	redirectUrl := req.URL.Query().Get("url")
	cookieToken := new(http.Cookie)
	cookieToken.Name = "token"
	if redirectUrl == "" || token == "" {
		cookieToken.Value = "logout"
	} else {
		cookieToken.Value = token
	}
	redirectUrl = redirectUrl + ""
	refDomain := conf.Get("37client", "url")
	cookieToken.Domain = strings.Split(refDomain, "//")[1]
	http.SetCookie(res, cookieToken)
	http.Redirect(res, req, redirectUrl, http.StatusMovedPermanently)
}

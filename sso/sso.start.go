package main

import (
	"fmt"
	"my/util"
	"net/http"
	"os"
	"sso/common"
)

var conf *util.Config
var wd, _ = os.Getwd()

func init() {

	prefileConf := util.NewConfig(wd + util.GetSysSplit() + "prefile.ini")
	model := prefileConf.Get("env", "model")
	confFileName := "conf_pro.ini"
	switch model {
	case "uat":
		confFileName = "conf_uat.ini"
	case "dev":
		confFileName = "conf_dev.ini"
	}
	confPath := wd + util.GetSysSplit() + "sso" + util.GetSysSplit() + confFileName
	conf = common.SetConfig(confPath)
	if conf.Get("sysinfo", "showSQL") == "true" {
		util.InitDB(conf, true)
	} else {
		util.InitDB(conf, false)
	}

	if conf.Get("sysinfo", "initializedDatabase") == "true" {
		common.InitUser()
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd+util.GetSysSplit()+"sso"+util.GetSysSplit()+"static"))))
	mux.HandleFunc("/edit", common.Edit)
	mux.HandleFunc("/reg", common.Signin)
	mux.HandleFunc("/login", common.Login)
	mux.HandleFunc("/restore", common.Restore)
	mux.HandleFunc("/exit", common.Exit)
	mux.HandleFunc("/find", common.Find)
	mux.HandleFunc("/byids", common.ByIds)
	mux.HandleFunc("/byToken", common.ByToken)
	mux.HandleFunc("/myinfo", common.Myinfo)
	listenAndServe := conf.Get("sysinfo", "port")
	fmt.Printf("%v SSO Start OK.\n", listenAndServe)
	http.ListenAndServe(listenAndServe, mux)

}

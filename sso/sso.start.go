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
var hostAndPort string

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
	hostAndPort = conf.Get("sysinfo", "port")
	if conf.Get("sysinfo", "showSQL") == "true" {
		util.InitDB(conf, true)
	} else {
		util.InitDB(conf, false)
	}

	if conf.Get("sysinfo", "initializedDatabase") == "true" {
		common.InitUser()
	}
	fmt.Printf("sso: %v, %v\n", model, hostAndPort)
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
	http.ListenAndServe(hostAndPort, mux)

}

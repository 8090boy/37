package main

import (
	"fmt"
	controller "hundred/controllers"
	"log"
	"my/util"
	"net/http"
	"os"

	"github.com/ant0ine/go-json-rest/rest"
)

var conf = new(util.Config)
var wd, _ = os.Getwd()
var hostAndPort string

func init() {

	//	prefileConf := util.NewConfig("F:\\GOPATH\\src\\prefile.ini")
	prefileConf := util.NewConfig(wd + util.GetSysSplit() + "prefile.ini")

	confFileName := "conf_pro.ini"
	model := prefileConf.Get("env", "model")

	switch model {
	case "uat":
		confFileName = "conf_uat.ini"
	case "dev":
		confFileName = "conf_dev.ini"
	}
	//	confPath := wd + util.GetSysSplit() + confFileName
	confPath := wd + util.GetSysSplit() + "hundred" + util.GetSysSplit() + confFileName

	conf = controller.SetConfig(confPath)
	hostAndPort = conf.Get("sysinfo", "port")
	if conf.Get("common", "showSQL") == "true" {
		util.InitDB(conf, true)
	} else {
		util.InitDB(conf, false)
	}

	if conf.Get("common", "initializedDatabase") == "true" {
		controller.InitData()
	}
	fmt.Printf("hundred: %v, %v\n", model, hostAndPort)
}
func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.ContentTypeCheckerMiddleware{})
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			stat := true
			switch origin {
			case "http://localhost":
				stat = true
			case "http://www.3737.io":
				stat = true
			case "http://uat.3737.io":
				stat = true
			case "http://dev.3737.io":
				stat = true
			}
			return stat
		},
		AllowedMethods:                []string{"GET", "POST"},
		AllowedHeaders:                []string{"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	api.Use(&rest.JsonpMiddleware{
		CallbackNameKey: "ok",
	})

	router, err := rest.MakeRouter(
		rest.Post("/my/edit", controller.Edit),
		rest.Get("/my/relation/:mob", controller.MyRelationUser),
		rest.Get("/my/code", controller.InvitationCode),
		rest.Get("/my/friendster", controller.Friendster),
		rest.Get("/todo/list", controller.TodoList),
		rest.Get("/todo/submit/:to", controller.SubmitTodo),
		rest.Get("/todo/not/:to", controller.NotTodo),
		rest.Get("/task/new", controller.NewTask),
		rest.Get("/task/find/:auid", controller.FindTask),
		rest.Get("/task/submit/:auid", controller.SubmitTask),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	api.SetApp(router)
	http.HandleFunc("/myinfo", controller.Myinfo)
	http.HandleFunc("/interaction", controller.Myrelational)
	http.HandleFunc("/login", controller.Login)   //login
	http.HandleFunc("/signin", controller.Signin) //reg
	http.Handle("/v1/", http.StripPrefix("/v1", api.MakeHandler()))

	http.ListenAndServe(hostAndPort, nil)

}

package models

import (
	"fmt"
	"my/util"
	"os"
	"testing"
)

var conf = new(util.Config)
var wd, _ = os.Getwd()

func init() {

	confPath := "F:" + util.GetSysSplit() + "GOPATH" + util.GetSysSplit() + "src" + util.GetSysSplit() + "hundred" + util.GetSysSplit() + "conf_dev.ini"
	conf = setConfig(confPath)

	if conf.Get("common", "showSQL") == "true" {
		util.InitDB(conf, true)
	} else {
		util.InitDB(conf, false)
	}

}

func TestUpgrage(t *testing.T) {

	aud := new(Audit).ByUpgargeMonad(4)
	fmt.Println(aud)
}

func setConfig(filepath string) *util.Config {
	conf = new(util.Config)
	conf.Filepath = filepath
	return conf
}

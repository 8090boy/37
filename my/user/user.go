package user

import (
	"my/util"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine
var conf *util.Config

func init() {
	wd, _ := os.Getwd()
	confUrl := wd + util.GetSysSplit() + "my" + util.GetSysSplit() + "user" + util.GetSysSplit() + "user.ini"

	conf = util.SetConfig(confUrl)
	initDbEngine(conf)
}

func initDbEngine(conf *util.Config) {
	db = new(xorm.Engine)
	url := conf.Get("database", "url")
	name := conf.Get("database", "name")
	DBtype := conf.Get("database", "dbtype")
	password := conf.Get("database", "password")
	dbUrl := strings.Join([]string{strings.Join([]string{name, password}, ":"), url}, "@")
	db, _ = xorm.NewEngine(DBtype, dbUrl)
	db.ShowSQL(true)
}

type State struct {
	Token    string
	Userjson string
}

func (this *State) Find(token string) {
	sql := "select * from user_status as U where U.token =? limit 1"
	db.Query(sql, token)

}

func (this *State) Add() *State {
	db.Insert(this)
	return this

}

func (this *State) Del() {
	db.Delete(this)
}

func (this *State) DelAll() {
	sql := "DELETE FROM state"
	db.Exec(sql)

}

package util

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Eng *xorm.Engine

func InitDB(conf *Config, show bool) {
	Eng = new(xorm.Engine)
	url := conf.Get("database", "url")
	name := conf.Get("database", "name")
	DBtype := conf.Get("database", "dbtype")
	password := conf.Get("database", "password")
	dbUrl := strings.Join([]string{strings.Join([]string{name, password}, ":"), url}, "@")
	var err error
	Eng, err = xorm.NewEngine(DBtype, dbUrl)

	Eng.TZLocation, _ = time.LoadLocation(LoadLocation)
	if err != nil {
		fmt.Printf(err.Error())
	}
	Eng.ShowSQL(show)
}

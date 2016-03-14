package user

import (
	"fmt"
	"my/util"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type State struct {
	// 随机码
	Token string
	// 参考信息
	Userjson string
	// 时间
	Overdue time.Time
}

func (this *State) Find(token string) *State {
	sql := "select * from state as U where U.token =? and overdue > SYSDATE() limit 1"
	resultsSlice, _ := util.Eng.Query(sql, token)
	if len(resultsSlice) == 0 {
		return nil
	}

	data := new(State)
	var err error
	Slice := resultsSlice[0]
	for k, v := range Slice {
		val := string(v)
		switch k {
		case "token":
			data.Token = val
		case "userjson":
			data.Userjson = val
		case "overdue":
			data.Overdue, err = time.Parse(util.TimeFormat, val)
		}

	}
	if err != nil {
		return nil
	}
	return data
}

func (this *State) Add() *State {
	_, err := util.Eng.Insert(this)
	if err != nil {
		fmt.Println(err)
	}

	return this

}

// 删除以往的token
func (this *State) DelOverdue(userjson string) {

	sql := "delete from state where overdue < SYSDATE() and userjson =?"
	_, err := util.Eng.Exec(sql, userjson)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *State) Del() {
	sql := "delete from state where token =?"
	_, err := util.Eng.Exec(sql, this.Token)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *State) DelMore() {
	sql := "delete from state where userjson=? and overdue < SYSDATE()"
	_, err := util.Eng.Exec(sql, this.Userjson)
	if err != nil {
		fmt.Println(err)
	}
}

func (cate *State) DelAll() {
	sql := "DELETE FROM state;"
	_, err := util.Eng.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}

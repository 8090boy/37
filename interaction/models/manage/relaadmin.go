package manage

import (
	"fmt"
	"my/util"
	"strconv"
)

// relational 管理顶级帐号
type Relaadmin struct {
	Relaid int64 // 被管理的
	Ssoid  int64 // 谁管
	Income int64 // 管理者收入
}

func (cate *Relaadmin) Add() {
	util.Eng.Insert(cate)

}

//
func (cate *Relaadmin) DelAll() {
	sql := "DELETE FROM relaadmin"
	util.Eng.Exec(sql)

}

func (data *Relaadmin) DelBySsoid(ssoid int64) {
	data.Ssoid = ssoid
	util.Eng.Delete(data)

}

func (data *Relaadmin) DelByRelaid(relaid int64) {
	data.Relaid = relaid
	util.Eng.Delete(data)

}

// 更新特殊账户金额
func (data *Relaadmin) UpdateWhereColName(relaid, ssoid int64) {
	util.Eng.Update(data, &Relaadmin{Relaid: relaid, Ssoid: ssoid})

}

func (ra *Relaadmin) FindByRelaId(relationalid int64) *Relaadmin {
	sql := "select * from relaadmin as U where U.relaid =? limit 1"
	slice, _ := util.Eng.Query(sql, relationalid)

	if len(slice) == 0 {

		ra = nil
		return ra
	}

	ra = ra.Compound(slice[0])
	return ra
}

func (ra *Relaadmin) FindBySsoId(relationalid int64) *Relaadmin {
	sql := "select * from relaadmin as U where U.ssoid =? limit 1"
	slice, _ := util.Eng.Query(sql, relationalid)
	if len(slice) == 0 {

		return nil
	}

	return ra.Compound(slice[0])
}

func (ra *Relaadmin) All(startId, count int) []Relaadmin {
	sql := "select * from user where id >= ? limit ?"
	resultsSlice, err := util.Eng.Query(sql, startId, count)
	datas := make([]Relaadmin, count)
	if err != nil || len(resultsSlice) == 0 {

		return datas
	}

	for index, _ := range resultsSlice {
		user := new(Relaadmin)
		user = ra.Compound(resultsSlice[index])
		datas[index] = *user
	}
	return datas
}

func (re *Relaadmin) Compound(Slice map[string][]byte) *Relaadmin {
	user := new(Relaadmin)
	var err error

	for k, v := range Slice {
		val := string(v)
		switch k {
		case "ssoid":
			user.Ssoid, err = strconv.ParseInt(val, 10, 64)
		case "relaid":
			user.Relaid, err = strconv.ParseInt(val, 10, 64)
		case "income":
			user.Income, err = strconv.ParseInt(val, 10, 64)
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	return user
}

package models

import (
	"fmt"
	"my/util"
	"strconv"
	"time"
)

// 审核任务流水表 audit
type Audit struct {
	Id int64
	//审核者sso id
	Sso int64
	// 审核者是否特殊账户,0不是特殊账户，1是特殊账户
	Special int
	// 审核者在产生这个待办时单子收款次数
	Count int
	//审核者的单子id
	MonadId int64
	//审核者id
	RelationalId int64
	//提交者sso id
	ProposerSso int64
	//提交者(id)
	ProposerRelationalId int64
	//提交者单子(id)
	ProposerMonadId int64
	// 提交者在产生这个待办时单子收款次数
	ProposerCount int
	//状态：0未确认，1已确认，2待提交
	Status int
	//是新单1,不是新单0,
	Isnewmonad int
	//提交日期
	Create time.Time
	//审核通过日期
	Operation time.Time
}

func NewAudit() *Audit {
	audit := new(Audit)
	audit.Create = time.Now()
	audit.Operation = audit.Create
	audit.Status = 0
	return audit
}

// form  发起者
// to    审核者
func AddAudit(formUserId, formMonadId, formRelaId, toUserId, toMonadId, toRelaId int64) {
	audit := NewAudit()
	// 申请人
	audit.ProposerSso = formUserId
	audit.ProposerMonadId = formMonadId
	audit.ProposerRelationalId = formRelaId
	// 待办接受人
	audit.Sso = toUserId
	audit.MonadId = toMonadId
	audit.RelationalId = toRelaId
	//
	audit.Add()
}

func (cate *Audit) Add() (int64, error) {
	refId, err := util.Eng.Insert(cate)
	if err != nil {
		return 0, err
	}
	return refId, nil
}

func DelAllAudit() {
	sql := "DELETE FROM audit"
	sql1 := "ALTER TABLE audit AUTO_INCREMENT=1;"
	_, err := util.Eng.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	_, err = util.Eng.Exec(sql1)
	if err != nil {
		fmt.Println(err)
	}
}

func (data *Audit) Del(id int64) {
	data.Id = id
	util.Eng.Delete(data)
}

func (cate *Audit) ById(id int64) (*Audit, error) {
	_, err := util.Eng.Id(id).Get(cate)
	if err != nil {
		return nil, err
	}
	return cate, nil
}
func (data *Audit) Edit() {

	_, err := util.Eng.Id(data.Id).Update(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (data *Audit) UpdateByColName(colName ...string) {
	_, err := util.Eng.Id(data.Id).Cols(colName...).Update(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (e *Audit) All(startId, count int) ([]Audit, error) {
	sql := "select * from audit where id >= ? limit ?"
	resultsSlice, err := util.Eng.Query(sql, startId, count)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]Audit, count)

	for i, slice := range resultsSlice {
		data := new(Audit)
		data.compound(slice, data)
		if data.Id >= 1 {
			datas[i] = *data
		}
	}
	return datas, nil
}

// 根据升级单子的id
func (e *Audit) ByUpgargeMonad(id int64) *Audit {
	sql := "select * from audit where proposer_monad_id = ?"
	resultsSlice, err := util.Eng.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]*Audit, len(resultsSlice))
	count := 0
	if len(resultsSlice) == 0 {
		return nil
	}
	for i, slice := range resultsSlice {

		data := new(Audit)
		data.compound(slice, data)

		if data.Id >= 1 {
			datas[i] = data
			count++
		}
	}
	return datas[0]
}

// 根据接受者sso id
func (e *Audit) BySso(id int64, stat, special int) ([]*Audit, int) {
	sql := "select * from audit where sso = ? and status =? and special=?"
	resultsSlice, err := util.Eng.Query(sql, id, stat, special)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]*Audit, len(resultsSlice))
	count := 0
	if len(resultsSlice) == 0 {
		return datas, count
	}
	for i, slice := range resultsSlice {

		data := new(Audit)
		data.compound(slice, data)

		if data.Id >= 1 {
			datas[i] = data
			count++
		}
	}
	return datas, count
}

// 根据提交者rela id,status,special
func (e *Audit) ByPropRela(id int64, status, special int) ([]*Audit, int) {
	sql := "select * from audit where proposer_relational_id = ? and status =? and special=?"
	resultsSlice, err := util.Eng.Query(sql, id, status, special)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]*Audit, len(resultsSlice))
	count := 0
	if len(resultsSlice) == 0 {
		return datas, count
	}
	for i, slice := range resultsSlice {

		data := new(Audit)
		data.compound(slice, data)

		if data.Id >= 1 {
			datas[i] = data
			count++
		}
	}
	return datas, count
}

// 寻找提交者没有被审核的任务
func (e *Audit) NotOk(id int64) ([]*Audit, int) {
	sql := "select * from audit where proposer_relational_id = ? and status !=1"
	resultsSlice, err := util.Eng.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]*Audit, len(resultsSlice))
	count := 0
	if len(resultsSlice) == 0 {
		return datas, count
	}
	for i, slice := range resultsSlice {

		data := new(Audit)
		data.compound(slice, data)

		if data.Id >= 1 {
			datas[i] = data
			count++
		}
	}
	return datas, count
}

//  获取提交者的任务
func (e *Audit) AuditsByPropRela(id int64) ([]*Audit, int) {
	sql := "select * from audit where proposer_relational_id = ? "
	resultsSlice, err := util.Eng.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]*Audit, len(resultsSlice))
	count := 0
	if len(resultsSlice) == 0 {
		return datas, count
	}
	for i, slice := range resultsSlice {
		data := new(Audit)
		data.compound(slice, data)
		if data.Id >= 1 {
			datas[i] = data
			count++
		}
	}

	return datas, count
}

func (e *Audit) ByRela(id int64) ([]*Audit, int) {
	sql := "select * from audit where relational_id = ? and status =0 and special =0"
	resultsSlice, err := util.Eng.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]*Audit, len(resultsSlice))
	count := 0
	if len(resultsSlice) == 0 {
		return datas, count
	}
	for i, slice := range resultsSlice {

		data := new(Audit)
		data.compound(slice, data)

		if data.Id >= 1 {
			datas[i] = data
			count++
		}
	}
	return datas, count
}

func (audit *Audit) compound(slice map[string][]byte, aud *Audit) {

	var err error
	for k, v := range slice {
		val := string(v)
		switch k {
		case "id":
			aud.Id, err = strconv.ParseInt(val, 10, 64)
		case "sso":
			aud.Sso, err = strconv.ParseInt(val, 10, 64)
		case "count":
			aud.Count, err = strconv.Atoi(val)
		case "special":
			aud.Special, err = strconv.Atoi(val)
		case "relational_id":
			aud.RelationalId, err = strconv.ParseInt(val, 10, 64)
		case "proposer_sso":
			aud.ProposerSso, err = strconv.ParseInt(val, 10, 64)
		case "proposer_monad_id":
			aud.ProposerMonadId, err = strconv.ParseInt(val, 10, 64)
		case "proposer_relational_id":
			aud.ProposerRelationalId, err = strconv.ParseInt(val, 10, 64)
		case "proposer_count":
			aud.ProposerCount, err = strconv.Atoi(val)
		case "status":
			aud.Status, err = strconv.Atoi(val)
		case "isnewmonad":
			aud.Isnewmonad, err = strconv.Atoi(val)
		case "create":
			aud.Create, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "operation":
			aud.Operation, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if aud.Id < 0 {
		aud = nil
	}

}

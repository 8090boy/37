package models

import (
	"fmt"
	"my/util"
	"strconv"
	"time"
)

// 单子、主单
type Monad struct {
	//ID:
	Id int64
	//是否主单：0不是、1是
	IsMain int
	//主单号：子单默认是主单号
	MainMonad int64
	//状态:
	//状态与rela一致，0未激活、1正常、2冻结、3违规、4任务太多，9出局、9子单出局
	//冻结是可以解冻的，短期停止收益；
	//因此冻结因任务太多而冻结是可以解冻，短期停止收益；
	//违规就是封单，不能参与任何游戏活动，人为设置违规。
	//结束是自己主单出局时下面的子单的正常结束
	State int
	//级别:0-7， null、初、中、高
	Class int
	//收款次数
	Count int
	//任务次数
	Task int
	//37 Relational ID
	Pertain int64
	//父级单子ID	 :
	ParentMonad int64
	//应出了多少单解冻
	UnfreezePeriodCount int
	//创建日期  :
	Create time.Time
	//冻结时间：
	Freeze time.Time
	//解冻时间：
	UnFreeze time.Time
	//升级提交时间	 :
	Upgrade time.Time
	//审核通过日期 :
	Audit time.Time
	//审核通过各级数量: 0 12 6 8 6 5 50
	AuditNumber string
}

func NewMonad() *Monad {
	cate := new(Monad)
	cate.Upgrade = time.Now()
	cate.Create = time.Now()
	cate.Class = 0
	cate.State = 0
	return cate
}

func (cate *Monad) Add() (int64, error) {
	_, err := util.Eng.Insert(cate)
	if err != nil {
		return 0, err
	}
	return cate.Id, nil
}

func DelAllMonad() {
	sql := "DELETE FROM monad"
	sql1 := "ALTER TABLE monad AUTO_INCREMENT=1;"
	_, err := util.Eng.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	_, err = util.Eng.Exec(sql1)
	if err != nil {
		fmt.Println(err)
	}
}

func (data *Monad) Del(id int64) error {
	data.Id = id
	_, err := util.Eng.Delete(data)
	return err
}

func (cate *Monad) ById(id int64) *Monad {
	sql := "select * from monad as M where M.id=? limit 1"
	resultsSlice, err := util.Eng.Query(sql, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(resultsSlice) == 0 {
		return nil
	}
	m := cate.compoundNew(resultsSlice[0])
	if m == nil {
		return nil
	}
	return m
}

func (cate *Monad) FindById(id int64) *Monad {
	sql := "select * from monad as M where M.id=? limit 1"
	resultsSlice, err := util.Eng.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	if len(resultsSlice) == 0 {
		cate = nil
	}
	cate = cate.compoundNew(resultsSlice[0])
	return cate
}

func (data *Monad) Edit() {
	_, err := util.Eng.Id(data.Id).Update(data)
	if err != nil {
		fmt.Println(err)
	}
}

// 多个id查询
func (cate *Monad) ByIds(ids string) ([]*Monad, int) {
	sql := "select * from monad as M where M.id in ( ? )"
	result, err := util.Eng.Query(sql, ids)
	if err != nil {
		fmt.Println(err)
	}
	Monads := make([]*Monad, 10)
	no := 0
	for i, slice := range result {
		m := cate.compoundNew(slice)
		if m.Id > 0 {
			no = no + 1
			Monads[i] = m
		}
	}
	return Monads, no
}

// 更具rela id，monad class 寻找单子
func (cate *Monad) MonadsByPertain(RelationalId int64, class int) ([]*Monad, int) {
	sql := "select * from monad as M where M.pertain = ? and class = ?"
	resultsSlice, err := util.Eng.Query(sql, RelationalId, class)
	if err != nil {
		fmt.Println(err)

	}
	Monads := make([]*Monad, len(resultsSlice))
	no := 0
	for i, Slice := range resultsSlice {
		m := cate.compoundNew(Slice)
		if m.Id > 0 {
			no = no + 1
			Monads[i] = m
		}
	}
	return Monads, no
}

func (cate *Monad) MonadChildsByParentMonad(parent_monad int64) ([]*Monad, int) {
	sql := "select * from monad as M where M.parent_monad = ?"
	resultsSlice, err := util.Eng.Query(sql, parent_monad)
	if err != nil {
		fmt.Println(err)
	}
	Monads := make([]*Monad, 10)
	no := 0
	for i, Slice := range resultsSlice {
		m := cate.compoundNew(Slice)
		if m.Id > 0 {
			no = no + 1
			Monads[i] = m
		}
	}
	return Monads, no
}

func (e *Monad) compoundNew(Slice map[string][]byte) *Monad {
	data := new(Monad)
	var err error
	for k, v := range Slice {
		val := string(v)
		switch k {
		case "id":
			data.Id, err = strconv.ParseInt(val, 10, 64)
		case "main_monad_no":
			data.MainMonad, err = strconv.ParseInt(val, 10, 64)
		case "pertain":
			data.Pertain, err = strconv.ParseInt(val, 10, 64)
		case "parent_monad":
			data.ParentMonad, err = strconv.ParseInt(val, 10, 64)
		case "is_main":
			data.IsMain, err = strconv.Atoi(val)
		case "state":
			data.State, err = strconv.Atoi(val)
		case "class":
			data.Class, err = strconv.Atoi(val)
		case "count":
			data.Count, err = strconv.Atoi(val)
		case "task":
			data.Task, err = strconv.Atoi(val)
		case "unfreeze_period_count":
			data.UnfreezePeriodCount, err = strconv.Atoi(val)
		case "create":
			data.Create, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "freeze":
			data.Freeze, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "un_freeze":
			data.UnFreeze, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "upgrade":
			data.Upgrade, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "audit":
			data.Audit, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "audit_number":
			data.AuditNumber = val
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	return data
}

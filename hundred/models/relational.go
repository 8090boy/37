package models

import (
	"fmt"
	"my/util"
	"strconv"
	"sync"
	"time"
)

// 用户关系表 relational
//
type Relational struct {
	//本轮游戏37 ID
	Id int64
	//自己手机号码
	Mobile string
	//推荐人
	Referrer string
	//当前37号
	CurrentMonad int64
	//收入
	Income int64
	//支出
	Spending int64
	//亏损
	Loss int64
	//角色：1会员2普通管理员（处理违规，处理投诉）、3超级会员、4股东只有查看功能、或其它
	Role int
	//状态与主单一致，0未激活、1正常、2冻结、3违规、4任务太多，9出局、9子单出局
	Status int
	// 20元收款次数
	OneCount int
	// 50元收款次数
	TwoCount int
	// 100元收款次数
	ThreeCount int
	//推荐总数
	RecommandTotal int
	//自己用户ID > SSO系统用户id
	SsoId int64
	//上次登录信息
	PrevInfo string
	//历史登录信息
	HistoryTrack string
	//历史37号
	HistoryMonads string
	//我的空闲时间
	Freetime string
	//创建日期 :
	Create time.Time
	//上次登录时间
	Prev time.Time
	//上次出单时间
	PrevNewMonad time.Time
}

func (cate *Relational) Add() (int64, error) {
	refId, err := util.Eng.Insert(cate)
	if err != nil {
		return 0, err
	}
	return refId, nil
}

func DelAllRelational() {
	sql := "DELETE FROM relational"
	sql1 := "ALTER TABLE relational AUTO_INCREMENT=1;"
	_, err := util.Eng.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	_, err = util.Eng.Exec(sql1)
	if err != nil {
		fmt.Println(err)
	}
}

func (data *Relational) Del(cid string) error {

	data.Id, _ = strconv.ParseInt(cid, 10, 64)
	_, err := util.Eng.Delete(data)
	return err
}

func (cate *Relational) FindReferrer(relaId string) *Relational {
	sql := "select * from relational where referrer = ? limit 1"
	resultsSlice, err := util.Eng.Query(sql, relaId)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(resultsSlice) == 0 {
		return nil
	}
	cate = cate.CompoundSingle(resultsSlice[0])
	return cate
}

func (rela *Relational) FindByReferrer(relaid int64) []*Relational {
	sql := "select * from relational where referrer = ?"
	resultsSlice, err := util.Eng.Query(sql, relaid)
	if err != nil {
		fmt.Println(err)
	}

	datas := make([]*Relational, len(resultsSlice))
	for i, slice := range resultsSlice {
		rela = rela.CompoundSingle(slice)
		if rela.Id > 0 {
			datas[i] = rela
		}
	}
	return datas
}

func (rela *Relational) FindRecommended(relaid int64) ([]*Relational, int) {
	sql := "select * from relational where referrer = ?"
	resultsSlice, err := util.Eng.Query(sql, relaid)
	if err != nil {
		fmt.Println(err)
	}

	datas := make([]*Relational, len(resultsSlice))
	for i, slice := range resultsSlice {
		rela = rela.CompoundSingle(slice)
		if rela.Id > 0 {
			datas[i] = rela
		}
	}
	return datas, len(datas)
}

func (cate *Relational) ById(cid int64) *Relational {
	sql := "select * from relational where id = ? limit 1"
	resultsSlice, err := util.Eng.Query(sql, cid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(resultsSlice) == 0 {
		return nil
	}
	cate = cate.CompoundSingle(resultsSlice[0])
	return cate
}

func (cate *Relational) ByIds(ids string) ([]*Relational, int) {
	sql := "select * from relational as M where M.id in ( ? )"
	result, err := util.Eng.Query(sql, ids)
	if err != nil {
		fmt.Println(err)
	}

	cates := make([]*Relational, 10)
	no := 0
	for i, slice := range result {
		rela := cate.CompoundSingle(slice)
		if rela != nil {
			if rela.Id > 0 {
				no = no + 1
				cates[i] = rela
			}
		}
	}
	return cates, no
}

func (cate *Relational) FindById(cid int64) *Relational {
	sql := "select * from relational where id = ? limit 1"
	resultsSlice, err := util.Eng.Query(sql, cid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(resultsSlice) == 0 {
		return nil
	}

	return cate.CompoundSingle(resultsSlice[0])
}

func (cate *Relational) BySsoId(cid int64) *Relational {
	if cid == 0 {
		return nil
	}
	sql := "select * from relational where sso_id = ? limit 1"
	resultsSlice, err := util.Eng.Query(sql, cid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(resultsSlice) == 0 {
		return nil
	}

	return cate.CompoundSingle(resultsSlice[0])
}

func (cate *Relational) FindBySsoId(cid int64) *Relational {
	sql := "select * from relational where sso_id = ? limit 1"
	resultsSlice, err := util.Eng.Query(sql, cid)
	if err != nil {
		fmt.Println(err)
	}
	if len(resultsSlice) == 0 {
		return nil
	}

	return cate.CompoundSingle(resultsSlice[0])

}

func (cate *Relational) ByMob(mob string) *Relational {

	sql := "select * from relational where mobile = ? limit 1"
	reSlice, err := util.Eng.Query(sql, mob)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(reSlice) == 0 {
		return nil
	}
	return cate.CompoundSingle(reSlice[0])
}

func (data *Relational) Edit() {
	editSync1 := new(sync.Mutex)
	editSync1.Lock()
	fmt.Println("-------11111111111111111------")
	_, err := util.Eng.Id(data.Id).Update(data)
	fmt.Println("-------22222222222222222------")
	defer editSync1.Unlock()
	if err != nil {
		fmt.Println(err)
	}
}

func (data *Relational) UpdateByColsName(names ...string) {

	_, err := util.Eng.Id(data.Id).Cols(names...).Update(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (e *Relational) All(startId, count int) ([]Relational, error) {
	sql := "select * from issuance where id >= ? limit ?"
	resultsSlice, err := util.Eng.Query(sql, startId, count)
	if err != nil {
		fmt.Println(err)
	}
	datas := make([]Relational, count)

	for i, slice := range resultsSlice {
		data := new(Relational)
		data = e.CompoundSingle(slice)
		if data != nil {
			if data.Id > 0 {
				datas[i] = *data
			}
		}
	}
	return datas, nil
}

// 返回一个
func (cate *Relational) CompoundSingle(slice map[string][]byte) *Relational {
	if len(slice) == 0 {
		return nil
	}

	var err error
	data := new(Relational)
	for k, v := range slice {
		val := string(v)
		switch k {
		case "id":
			data.Id, err = strconv.ParseInt(val, 10, 64)
		case "income":
			data.Income, err = strconv.ParseInt(val, 10, 64)
		case "spending":
			data.Spending, err = strconv.ParseInt(val, 10, 64)
		case "loss":
			data.Loss, err = strconv.ParseInt(val, 10, 64)
		case "parent":
			data.Mobile = val
		case "freetime":
			data.Freetime = val
		case "mobile":
			data.Mobile = val
		case "referrer":
			data.Referrer = val
		case "current_monad":
			data.CurrentMonad, err = strconv.ParseInt(val, 10, 64)
		case "one_count":
			data.OneCount, err = strconv.Atoi(val)
		case "two_count":
			data.TwoCount, err = strconv.Atoi(val)
		case "three_count":
			data.ThreeCount, err = strconv.Atoi(val)
		case "role":
			data.Role, err = strconv.Atoi(val)
		case "status":
			data.Status, err = strconv.Atoi(val)
		case "recommand_total":
			data.RecommandTotal, err = strconv.Atoi(val)
		case "sso_id":
			data.SsoId, err = strconv.ParseInt(val, 10, 64)
		case "prev_info":
			data.PrevInfo = val
		case "history_track":
			data.HistoryTrack = val
		case "history_monads":
			data.HistoryMonads = val
		case "create":
			data.Create, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "prev":
			data.Prev, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		case "prev_new_monad":
			data.PrevNewMonad, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		}
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return data
}

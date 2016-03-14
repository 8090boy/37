package models

import (
	"fmt"
	"my/util"
	"strconv"
	"time"
)

// 用户消息表 message
type Message struct {
	Id int64
	//单子id
	MId int64
	//单子级别,这个单子收款次数
	MClass int
	//用户Id
	RId int64
	//消息类别 1升级，2警示
	Typee int
	//消息内容
	Content string
	//状态：查看/未查看
	Status int
	//提交审核日期
	Create time.Time
}

func NewMessage() *Message {
	mes := new(Message)
	mes.Create = time.Now()
	return mes
}

func (cate *Message) Add() *Message {
	_, err := util.Eng.Insert(cate)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cate
}

func DelAllMessage() {
	sql := "DELETE FROM message"
	sql1 := "ALTER TABLE message AUTO_INCREMENT=1;"
	_, err := util.Eng.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	_, err = util.Eng.Exec(sql1)
	if err != nil {
		fmt.Println(err)
	}
}
func (data *Message) Del(id int64) {
	sql := "DELETE FROM `message` WHERE `id` = ?"
	util.Eng.Exec(sql, id)
}

func (cate *Message) ById() {
	if cate.Id == 0 {
		cate = nil
		return
	}
	_, err := util.Eng.Id(cate.Id).Get(cate)
	if err != nil {
		fmt.Println(err)
	}
}

func (data *Message) Edit(id int64) {
	if data.Id == 0 {
		return
	}
	_, err := util.Eng.Id(id).Update(data)
	if err != nil {
		fmt.Println(err)
	}

}

func (mes *Message) ByMId(mId int64) {

	sql := "select * from message where m_id = ? limit 1"
	results, err := util.Eng.Query(sql, mId)
	if err != nil || len(results) < 1 {
		fmt.Println(err)
	}
	for _, slice := range results {
		mes.compound(slice)
	}

}

func (e *Message) ByRelaId(rId int64, statu int) []*Message {
	var count int = 100
	sql := "select * from message where r_id = ? and status =?  limit ?"
	results, err := util.Eng.Query(sql, rId, statu, count)
	datas := make([]*Message, len(results))
	if err != nil {
		fmt.Println(err)
		return datas
	}
	for i, slice := range results {

		mes := new(Message)
		mes.compound(slice)
		if mes.Id >= 1 {
			datas[i] = mes
		}
	}

	return datas
}

func (e *Message) All(startId, count int) []Message {
	sql := "select * from message where id >= ? limit ?"
	results, err := util.Eng.Query(sql, startId, count)
	datas := make([]Message, count)
	if err != nil {
		fmt.Println(err)
		return datas
	}
	for i, slice := range results {
		mes := new(Message)
		mes.compound(slice)
		if mes.Id >= 1 {
			datas[i] = *mes
		}
	}
	return datas
}

func (data *Message) compound(slice map[string][]byte) {
	var err error
	for k, v := range slice {
		val := string(v)
		switch k {
		case "id":
			data.Id, err = strconv.ParseInt(val, 10, 64)
		case "m_id":
			data.MId, err = strconv.ParseInt(val, 10, 64)
		case "m_class":
			data.MClass, err = strconv.Atoi(val)
		case "r_id":
			data.RId, err = strconv.ParseInt(val, 10, 64)
		case "typee":
			data.Typee, err = strconv.Atoi(val)
		case "content":
			data.Content = val
		case "status":
			data.Status, err = strconv.Atoi(val)
		case "create":
			data.Create, err = time.ParseInLocation(util.TimeFormat, val, util.Loca)
		}
		if err != nil {
			break
		}
	}
}

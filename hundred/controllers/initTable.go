package controllers

import (
	"fmt"
	model "interaction/models"
	manage "interaction/models/manage"

	"time"
)

// 初始化数据
func InitData() {
	model.DelAllAudit()
	model.DelAllMessage()
	model.DelAllMonad()
	model.DelAllRelational()
	initTJZ("18620131415", "25", 35)
	initSpaceMember("17750662398")                           //空号
	initRelationalAndMonadData("18650710067", "13790226216") // 股东
	initRelationalAndMonadData("18059244379", "13790226216") // 股东
}

//初始化直接加入37
func initTJZ(mobile, referrerId string, parentMonadId int64) {
	conf = GetConfig()
	relational := &model.Relational{}
	relational.Mobile = mobile
	relational.Referrer = referrerId
	relational.Create = time.Now()
	relational.Prev = relational.Create
	relational.PrevNewMonad = relational.Create
	relational.Status = 1
	//
	myUser := findUserByMob(mobile)
	relational.SsoId = myUser.Id
	_, err := relational.Add()
	if err != nil {
		fmt.Println(err)
	}
	//
	monad := model.NewMonad()
	monad.State = 1
	monad.Class = 6
	monad.MainMonad = 1
	monad.IsMain = 1
	monad.Pertain = relational.Id
	monad.ParentMonad = parentMonadId
	refId, _ := monad.Add()
	//
	relational.CurrentMonad = refId
	relational.Edit()
}

// 创建系统股东号码
// 新加股东号码，管理者号码
func initRelationalAndMonadData(topMob, refMob string) {
	conf = GetConfig()
	mobile := topMob
	relational := &model.Relational{}
	relational.Mobile = mobile
	relational.Referrer = "top"
	relational.Create = time.Now()
	relational.Prev = relational.Create
	relational.Status = 1
	//
	myUser := findUserByMob(mobile)
	relational.SsoId = myUser.Id
	_, err := relational.Add()
	if err != nil {
		fmt.Println(err)
	}
	//
	monad := model.NewMonad()
	monad.State = 1
	monad.Class = 9
	monad.MainMonad = 1
	monad.IsMain = 1
	monad.Pertain = relational.Id
	monad.ParentMonad = 0 //顶级单父级必须为0
	refId, _ := monad.Add()
	//
	relational.CurrentMonad = refId
	relational.Edit()
	//
	relaAdmin := new(manage.Relaadmin)
	refSsoId := findUserByMob(refMob)
	relaAdmin.Ssoid = refSsoId.Id
	relaAdmin.Relaid = relational.Id
	relaAdmin.Add()
}

// 设置空号
func initSpaceMember(mob string) {
	relaAdmin := new(manage.Relaadmin)
	relaAdmin.DelAll()
	refSsoId := findUserByMob(mob)
	relaAdmin.Ssoid = refSsoId.Id
	relaAdmin.Relaid = 0
	relaAdmin.Add()
}

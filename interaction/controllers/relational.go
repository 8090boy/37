package controllers

import (
	"encoding/json"
	"fmt"
	"sso/user"
	"sync"

	model "interaction/models"
	manage "interaction/models/manage"
	"my/util"
	"net/http"
	"strconv"
	"time"
)

var lockValidate sync.Mutex

// 获取自己37相关数据
func Myrelational(res http.ResponseWriter, req *http.Request) {
	lockValidate.Lock()
	defer lockValidate.Unlock()
	conf = util.GetConfig()
	callback := req.FormValue("cb")
	var sweet map[string]interface{} = make(map[string]interface{})
	sweet["state"] = 0
	//my user info
	stat, myInfo := validateUserInfo(req)

	if !stat {
		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}
	//
	sweet["state"] = 1
	//my relational
	relational := new(model.Relational)
	relational = relational.BySsoId(myInfo.Id)

	if relational == nil {
		// is  Specificity account
		spcif := new(manage.Relaadmin).FindBySsoId(myInfo.Id)
		if spcif != nil {
			if spcif.Ssoid > 0 && myInfo.Id == spcif.Ssoid {
				sweet["s"] = 2
				sweet["income"] = spcif.Income
				audits, _ := new(model.Audit).BySso(myInfo.Id, 0, 1)
				if len(audits) > 0 {
					aus, count := new(model.Audit).BySso(spcif.Ssoid, 0, 1)
					if count > 0 {
						sweet["todos"] = findAuditorInfoForIds(aus)
					}
				}
			}
		}

		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}
	// 是特殊账户
	if relational.Referrer == "top" {
		sweet["s"] = 1
	}
	// 是否出局
	if relational.Status == RELA_STATUS_DISCARD || relational.Status == RELA_STATUS_Retired {
		sweet["state"] = 2
		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}
	sweet["r"] = *relational
	unfreezeDatetime := conf.Get("common", "unfreezeDatetime")
	refHours, _ := time.ParseDuration(unfreezeDatetime)   // 未出单// s秒,m分,h小时,月
	prevNewMonad := relational.PrevNewMonad.Add(refHours) // 剩余时间将冻结
	sweet["pnm"] = prevNewMonad
	// 有收过款的、不是股东、出过一次子单的间隔一定时间没有出单 ,会冻结
	//my main monad
	myMainmonad := new(model.Monad)
	if relational.CurrentMonad > 0 {
		myMainmonad = myMainmonad.ById(relational.CurrentMonad)
	}
	if myMainmonad == nil || myMainmonad.State == RELA_STATUS_Retired || myMainmonad.State == RELA_STATUS_DISCARD {
		sweet["state"] = 99
		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}

	// 有收益的普通会员
	if relational.Income > 0 && relational.Referrer != "top" {
		// 检查自己状态
		time.LoadLocation(util.LoadLocation)
		sysNow := time.Now().Local()
		//		unfreezeDatetime := conf.Get("common", "unfreezeDatetime")
		retirementDatetime := conf.Get("common", "retirementDatetime")
		//;h hour
		//;m minutes
		//;s second
		//	refHours, _ := time.ParseDuration(unfreezeDatetime)           // 未出单// s秒,m分,h小时,月
		refUnFreezeHours, _ := time.ParseDuration(retirementDatetime) // 需要做完任务
		//	unFreezeTime := relational.PrevNewMonad.Add(refHours)         // 解冻期
		// 解冻时间期限
		retirementDate := myMainmonad.UnFreeze.Add(refUnFreezeHours) // 非正常出局期
		// 更新玩啥37系统的时间
		relational.Prev = time.Now().Local()
		// 是冻结状态再检查是否该出局
		if relational.Status == RELA_STATUS_FREEZE {
			// 是否非正常出局
			if sysNow.After(retirementDate) { // 符合非正常出局条件
				relational.Status = RELA_STATUS_DISCARD
				myMainmonad.State = RELA_STATUS_DISCARD
				//	fmt.Println("非正常出局,手机单独记录，防止再次注册")
				relational.Edit()
				myMainmonad.Edit()
			}
		}

		// 正常状态时未完成任务，可能被冻结
		if relational.Status == RELA_STATUS_NORMAL || relational.Status == RELA_STATUS_FOUR {
			updateStatusIfTasks(relational, myMainmonad)
		}

		// 正常状态时，可能长时间未出单需要冻结
		//		if relational.Status == RELA_STATUS_NORMAL {
		//			if sysNow.After(unFreezeTime) {
		//				unfreezeMonadCount := conf.Get("common", "unfreezeMonadCount")
		//				unfmc, _ := strconv.Atoi(unfreezeMonadCount)
		//				relational.Status = RELA_STATUS_FREEZE
		//				myMainmonad.State = RELA_STATUS_FREEZE
		//				myMainmonad.UnfreezePeriodCount = unfmc
		//				myMainmonad.UnFreeze = sysNow.Add(refUnFreezeHours)
		//				myMainmonad.Freeze = sysNow
		//				relational.Edit()
		//				myMainmonad.Edit()
		//			}
		//		}

	}
	// 有收入的股东和普通会员都要定是产生单字
	if relational.Income > 0 {
		// 正常状态时需要自动出单
		if relational.Status == RELA_STATUS_NORMAL {
			_autoNewMonad(myInfo, relational, myMainmonad)
		}
	}

	if myMainmonad.Id > 0 {
		sweet["m"] = *myMainmonad
	}
	//my audits
	var audit model.Audit
	//
	// 出单信息
	//
	// 自己的单子未被别人审核状况
	myAudits, myCount := audit.ByPropRela(relational.Id, 0, 0)
	// 自己单子级别为0
	myMonads, myMonadsCount := myMainmonad.MonadsByPertain(relational.Id, 0)
	tmpAu := new(model.Audit)
	if myMonadsCount == 0 || myCount == 0 {
		tmpAu = nil
	} else {
		tmpAu = findMonadByMyauditsInMyMonadsAndClassIsZero(myAudits, myMonads)
	}
	//
	if tmpAu == nil { // 已出的新单子全部被对方审核通过
		sweet["pi"] = false //可以出新单子
	} else { // 有一条没有被审核的新单子
		tmpUser, tmpRela, _ := findURM(tmpAu.RelationalId)
		sweet["pi"] = resultAuditerInfo(tmpRela, tmpUser) // 返回对方的信息
		// 对方推荐人信息
		if tmpRela.Referrer == "top" { //股东
			sweet["pri"] = findSpecificReferrer(tmpRela.Id)
		} else {
			tmpRelaId, _ := strconv.ParseInt(tmpRela.Referrer, 10, 64)
			tmpUserRef, tmpRelaRef, _ := findURM(tmpRelaId)
			if tmpUserRef != nil && tmpRelaRef != nil {
				sweet["pri"] = resultAuditerInfo(tmpRelaRef, tmpUserRef) // 返回对方推荐人信息
			}
		}
	}

	// 我的任务
	myTask, count := new(model.Audit).AuditsByPropRela(relational.Id)
	if count > 0 {
		sweet["tasks"] = myTask
	}
	aus, _ := new(model.Audit).ByRela(relational.Id)
	if len(aus) > 0 {
		sweet["todos"] = aus
	}
	sweet["interval"] = conf.Get("common", "interval")
	sweet["state"] = 100
	all_info, _ := json.Marshal(sweet)
	util.WriteJSONP(res, callback+"("+string(all_info)+")")
}

// 自动出单
// 是否应该出单
// 出单了就返回 true,没有出单返回 false
func _autoNewMonad(myUser *user.User, myRelational *model.Relational, myMainMonad *model.Monad) bool {
	conf = util.GetConfig()
	// 收入大于支出时，需要出单
	incomeIsOk := myRelational.Income > myRelational.Spending
	// 上次出单时间多久？，需要出单
	interval, _ := time.ParseDuration(conf.Get("common", "interval"))
	newPointer := myRelational.PrevNewMonad.Add(interval)
	now := time.Now().Local()
	prevNewMonadDateIsOk := now.After(newPointer)
	// 是否应该出单
	if !incomeIsOk && !prevNewMonadDateIsOk {
		fmt.Println("create monad false. 1")
		return false
	}
	// 出单
	myMonad := model.NewMonad()
	myMonad.Pertain = myRelational.Id
	myMonad.MainMonad = myRelational.CurrentMonad
	//
	parentRela, parMonad, flag := newSub(myMonad, myRelational, myMainMonad)
	// 没有位置
	if !flag {
		fmt.Println("create monad false. 2")
		return false
	}
	// 更新自己出单时间
	myRelational.PrevNewMonad = newPointer
	myRelational.Edit()
	// 因为对方上级单子处于冻结状态
	parMainMonad := new(model.Monad).ById(parentRela.CurrentMonad)
	if parentRela.Referrer == "top" {
		// add audit
		createAuditForNewMonad(myMonad, parMonad, myRelational, parentRela, 0, 0)
		fmt.Println("create monad true. 3")
		return true
	}
	state := false
	state = state || parMonad.State == RELA_STATUS_FREEZE
	state = state || parMonad.State == RELA_STATUS_FOUR
	state = state || parentRela.Status == RELA_STATUS_FREEZE
	state = state || parentRela.Status == RELA_STATUS_FOUR
	if parentRela.SsoId > 1 {
		state = state || parMainMonad.Class > 6
	}
	if state {
		parentRela.Loss = parentRela.Loss + INCOME[0]
		parentRela.Edit()
		// 指定帐号
		specialUserId := int64(3)
		// add audit
		createAuditForNewMonad(myMonad, parMonad, myRelational, parentRela, specialUserId, 0)
	} else {
		createAuditForNewMonad(myMonad, parMonad, myRelational, parentRela, 0, 0)
	}
	fmt.Println("create monad true. 4")
	return true
}

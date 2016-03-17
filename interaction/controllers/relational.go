package controllers

import (
	"encoding/json"
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
	stat, userA := validateUserInfo(req)

	if !stat {
		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}
	//
	sweet["state"] = 1
	//my relational
	relational := new(model.Relational)
	relational = relational.BySsoId(userA.Id)

	if relational == nil {
		// is  Specificity account
		spcif := new(manage.Relaadmin).FindBySsoId(userA.Id)
		if spcif != nil {
			if spcif.Ssoid > 0 && userA.Id == spcif.Ssoid {
				sweet["s"] = 2
				sweet["income"] = spcif.Income
				audits, _ := new(model.Audit).BySso(userA.Id, 0, 1)
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
		unfreezeDatetime := conf.Get("common", "unfreezeDatetime")
		retirementDatetime := conf.Get("common", "retirementDatetime")
		//;h hour
		//;m minutes
		//;s second
		refHours, _ := time.ParseDuration(unfreezeDatetime)           // 未出单// s秒,m分,h小时,月
		refUnFreezeHours, _ := time.ParseDuration(retirementDatetime) // 需要做完任务
		unFreezeTime := relational.PrevNewMonad.Add(refHours)         // 解冻期
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

		// 正常状态时，可能长时间未出单需要冻结
		if relational.Status == RELA_STATUS_NORMAL {
			if sysNow.After(unFreezeTime) {
				unfreezeMonadCount := conf.Get("common", "unfreezeMonadCount")
				unfmc, _ := strconv.Atoi(unfreezeMonadCount)
				relational.Status = RELA_STATUS_FREEZE
				myMainmonad.State = RELA_STATUS_FREEZE
				myMainmonad.UnfreezePeriodCount = unfmc
				myMainmonad.UnFreeze = sysNow.Add(refUnFreezeHours)
				myMainmonad.Freeze = sysNow
				relational.Edit()
				myMainmonad.Edit()
			}
		}

		// 正常状态时未完成任务，可能被冻结
		if relational.Status == RELA_STATUS_NORMAL || relational.Status == RELA_STATUS_FOUR {
			updateStatusIfTasks(relational, myMainmonad)
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

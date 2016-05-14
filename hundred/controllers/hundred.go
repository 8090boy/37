package controllers

import (
	"encoding/json"
	"fmt"

	"sso/user"
	"strings"
	"sync"

	model "hundred/models"
	manage "hundred/models/manage"
	"my/util"
	"net/http"
	"strconv"
	"time"
)

// 获取自己37相关数据
func Myrelational(res http.ResponseWriter, req *http.Request) {

	conf = GetConfig()
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
	sweet["state"] = 1 //正常账户
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

	if myMainmonad == nil || relational.Status == RELA_STATUS_Retired || relational.Status == RELA_STATUS_DISCARD {
		sweet["m"] = *myMainmonad
		sweet["state"] = 9
		all_info, _ := json.Marshal(sweet)
		util.WriteJSONP(res, callback+"("+string(all_info)+")")
		return
	}
	// 升级主单
	moandUpgrade(*myMainmonad)
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
	// 有收入的会员都要产生单子
	if relational.Income > 0 {
		if relational.Income > taskSum(relational) {
			// 正常状态时需要自动出单
			if relational.Status == RELA_STATUS_NORMAL {

				_autoCreateMonad(myInfo, relational, myMainmonad)

			}
		}

	}
	if myMainmonad.Id > 0 {
		sweet["m"] = *myMainmonad
	}

	//
	// 出单信息
	//
	//my audits
	var audit model.Audit
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
		if strings.ToLower(tmpRela.Referrer) == "top" { //股东
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

var updateLock *sync.Mutex

// monad升级
func moandUpgrade(monad model.Monad) bool {
	updateLock = new(sync.Mutex)
	updateLock.Lock()
	if monad.Id == 0 {
		return false
	}
	if monad.IsMain == 1 {
		// 需要推荐人员数量限制
		isOk, _, _ := mainMonadTask(&monad)
		if isOk == false { // 推荐人数不够
			return false
		}
	}
	rela := new(model.Relational).ById(monad.Pertain)
	// 收入大于 支出金额，才能产生任务
	if !assertIncomeGTspending(*rela, monad) {
		return false
	}
	// 产生升级任务
	// 产生升级任务
	//升级到几级
	targetLayer := monad.Class + 1
	// 付出金额 pay out
	income := INCOME[targetLayer]
	// 出一次单增加一次任务
	monad.Task = monad.Task + 1

	monad.Edit()
	updateLock.Unlock()
	targetMonad := findParentMonad(&monad, targetLayer)
	var targetRelaAmin *manage.Relaadmin
	// 审核方单子不存在
	if targetMonad == nil {
		// 设置收款人为运营组帐号
		// 给运营组帐号添加待办
		targetRelaAmin = targetRelaAmin.FindByRelaId(0)
		createAudit(&monad, nil, targetRelaAmin.Ssoid, false)
		return true
	}

	// 真正的收款人信息
	_, targetRela, targetMainMonad := findURM(targetMonad.Pertain)
	// 收款方主单或子单，rela状态不正常
	tarMainMoSata := targetRela.Status != 1
	// 收款方主或子单级别  小于 付款方单子级别
	tarMainMoClass := (targetMainMonad.Class < monad.Class) || (targetMonad.Class < monad.Class)
	// 是符合要求
	if tarMainMoSata || tarMainMoClass {
		// 由于以上两个条件不符合，真正的收款方需要增加损失
		targetRela.Loss = targetRela.Loss + income
		targetRela.UpdateByColsName("loss")
		if strings.ToLower(targetRela.Referrer) == "top" {
			targetRelaAmin = targetRelaAmin.FindByRelaId(targetRela.Id) // 是股东就用股东所对应的管理者
			createAudit(&monad, nil, targetRelaAmin.Ssoid, false)
			return true
		} else {
			targetRelaAmin = targetRelaAmin.FindByRelaId(0) // 不是股东就特定给0好id的管理者
			createAudit(&monad, nil, targetRelaAmin.Ssoid, false)
			return true
		}
	}
	createAudit(&monad, targetMonad, 0, false)
	return true
}

//收入 大于 总支出
// monad 当前收款单子
func assertIncomeGTspending(rela model.Relational, monad model.Monad) bool {
	// 重点走下这个流程，仔细分析
	if rela.Income == 0 {
		return false
	}

	if rela.Status != 1 {
		return false
	}

	if rela.CurrentMonad == 0 {
		return false
	}

	if currentMondUnfinished(&monad) {
		return false
	}

	// 级别为1级，并且收过两次款
	if monad.Class == 1 {
		if monad.Count == 1 {
			return true
		} else {
			return false
		}
	}

	// 是主单升级时
	if (rela.CurrentMonad == monad.Id) && (monad.IsMain == 1) {
		mulriple, _ := strconv.Atoi(conf.Get("common", "mulriple"))
		if monad.Class < 3 {
			return true
		}
	}

	// 预计支出金额
	refSpending := INCOME[monad.Class+1]
	// +实际已经支出
	spendingSum := rela.Spending + refSpending
	// +待确认的支出
	spendingSum = spendingSum + taskSum(&rela)
	// 总收入大于总支出
	if rela.Income > spendingSum {
		return true
	}
	return false
}

// 时间限制出单
func spaceOfTime(myRelational *model.Relational) bool {

	conf = GetConfig()
	// 收入大于支出时，需要出单
	incomeIsOk := myRelational.Income > myRelational.Spending
	// 上次出单时间多久？，需要出单
	interval, _ := time.ParseDuration(conf.Get("common", "interval"))
	newPointer := myRelational.PrevNewMonad.Add(interval)
	now := time.Now().Local()
	isOk := now.After(newPointer)
	// 是否应该出单
	if !incomeIsOk || !isOk {
		return false
	}
	// 更新自己出单时间
	myRelational.PrevNewMonad = newPointer
	myRelational.Edit()
	return true
}

// 自动出单
// 是否应该出单
// 出单了就返回 true,没有出单返回 false
func _autoCreateMonad(myUser *user.User, myRelational *model.Relational, myMainMonad *model.Monad) bool {
	// 时间限制
	sta := spaceOfTime(myRelational)
	if !sta {
		return false
	}
	// 出单总量
	conf = GetConfig()
	monadCount, _ := strconv.Atoi(conf.Get("common", "monadCount"))
	if myRelational.MonadCount == monadCount {
		return false
	}

	// 出单
	myMonad := model.NewMonad()
	myMonad.Pertain = myRelational.Id
	myMonad.MainMonad = myRelational.CurrentMonad
	parentRela, parMonad, flag := newSub(myMonad, myRelational, myMainMonad)
	// 出单失败
	if !flag {
		return false
	}
	myRelational.MonadCount = myRelational.MonadCount + 1
	myRelational.Edit()
	// 因为对方上级单子处于冻结状态
	parMainMonad := new(model.Monad).ById(parentRela.CurrentMonad)
	if parentRela.Referrer == "top" {
		// add audit
		createAudit(myMonad, parMonad, 0, true)
		return true
	}
	state := false
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
		createAudit(myMonad, nil, specialUserId, true)
	} else {
		createAudit(myMonad, parMonad, 0, true)
	}
	return true
}

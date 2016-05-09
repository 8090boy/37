package controllers

import (
	model "hundred/models"
	"hundred/models/manage"
	"sso/user"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

var submitLock *sync.Mutex

// 提交待办，审核单子
// 审核单子增加自己收入金额，增加别人的支出
// 设置审核状态或移除、删除审核信息
// 激活、解冻别人
func SubmitTodo(rep rest.ResponseWriter, req *rest.Request) {
	submitLock = new(sync.Mutex)
	submitLock.Lock()
	defer submitLock.Unlock()
	conf = GetConfig()
	// 返回信息
	result := make(map[string]interface{})
	to := req.PathParam("to")
	if to == "" {
		rep.WriteJson("0")
		return
	}
	cookie, _ := req.Cookie("token")
	token := cookie.Value
	// token 校验
	userA := new(user.User)
	byToken(token, userA)
	//
	audit := new(model.Audit)
	id, _ := strconv.ParseInt(to, 10, 64)
	audit, _ = audit.ById(id)
	if userA.Id != audit.Sso {
		rep.WriteJson("1")
		return
	}
	// 审核单子
	isOK := receiveAudit(*audit, *userA)
	if !isOK {
		rep.WriteJson("3")
		return
	}
	if audit.Special == 1 {
		rep.WriteJson("ok")
		return
	}
	//
	// 产生自己的升级任务 开始
	// 产生自己的升级任务 开始
	// 产生自己的升级任务 开始
	//
	myAuMonad := new(model.Monad).ById(audit.MonadId)
	// 任何单子级别大于6将不产生任务
	if myAuMonad.Class == 7 {
		result["influence"] = false
		rep.WriteJson(result)
		return
	}
	// 收款单子增加一次收入
	myAuMonad.Count = myAuMonad.Count + 1
	myAuMonad.Edit()
	isOk := moandUpgrade(*myAuMonad) // 我的收款单子

	if isOk {
		result["influence"] = true
	} else {
		result["influence"] = false
	}
	rep.WriteJson(result)
	return

}

// 没有收到红包
func NotTodo(rep rest.ResponseWriter, req *rest.Request) {
	// 返回信息
	//	result := make(map[string]interface{})
	to := req.PathParam("to")
	if to == "" {
		rep.WriteJson("0")
		return
	}
	cookie, _ := req.Cookie("token")
	token := cookie.Value
	// token 校验
	userA := new(user.User)
	byToken(token, userA)
	//
	audit := new(model.Audit)
	id, _ := strconv.ParseInt(to, 10, 64)
	audit, _ = audit.ById(id)
	if userA.Id != audit.Sso {
		rep.WriteJson("1")
		return
	}
	audit.Status = 2
	audit.Edit()
	rep.WriteJson("ok")
}

// 获取todo审核列表
func TodoList(rep rest.ResponseWriter, req *rest.Request) {
	so := req.URL.Query().Get("_")
	if so == "" {
		rep.WriteJson("")
		return
	}
	ssoIds := strings.Split(so, "|")
	if len(ssoIds) == 0 {
		rep.WriteJson("")
		return
	}
	results := make(map[string]interface{})
	for _, auid := range ssoIds {
		refId, _ := strconv.ParseInt(auid, 10, 64)
		audit, _ := new(model.Audit).ById(refId)

		if audit != nil {
			result := ""
			targetUser, _, _ := findURM(audit.ProposerRelationalId)
			result += targetUser.Mobile + "|"
			result += targetUser.Wechat + "|"
			result += targetUser.Alias + "|"
			refIndex := 0
			refIndex = audit.ProposerCount + 1
			sum := strconv.FormatInt(INCOME[refIndex], 10)
			result += sum + "|"
			result += audit.Create.String() + "|"
			//
			fromMonad := new(model.Monad).ById(audit.ProposerMonadId)
			result += strconv.Itoa(fromMonad.Class) + "|"
			result += strconv.Itoa(fromMonad.IsMain)
			results[auid] = result
			//

		}
	}
	///
	rep.WriteJson(results)
}

// receive audit
func receiveAudit(audit model.Audit, user user.User) bool {
	if user.Id != audit.Sso {
		return false
	}
	income := INCOME[audit.ProposerCount+1]
	// 查找出对方单子信息
	spendersMonad := new(model.Monad).ById(audit.ProposerMonadId)
	myRela := new(model.Relational)
	if myRela.Status == RELA_STATUS_Retired {
		return false
	}
	// 审核者账户
	// 是否特殊账户
	if audit.Special == 1 {
		// 特殊账户收入增加
		myRelaAdmin := new(manage.Relaadmin).FindBySsoId(user.Id)
		myRelaAdmin.Income = myRelaAdmin.Income + income
		myRelaAdmin.UpdateWhereColName(myRelaAdmin.Relaid, myRelaAdmin.Ssoid)

	} else {
		// 自己收入金额增加

		myRela = myRela.ById(audit.RelationalId)
		myRela.Income = myRela.Income + income
		// 自己是否该出局了
		maxIncomeRef := conf.Get("common", "maxIncome")
		maxIncome, _ := strconv.ParseInt(maxIncomeRef, 10, 64)
		if myRela.Income >= maxIncome {
			myRela.Status = RELA_STATUS_Retired
		}
		myRela.Edit()
	}
	//
	// 对方账户
	spendersRela := new(model.Relational).ById(audit.ProposerRelationalId)
	// 对方主单
	spenderMainMonad := findParentMainMonad(audit.ProposerRelationalId)

	if spendersRela.OneCount == RELA_STATUS_UNBORN {
		spendersRela.Status = 1
	}
	// 对方支出增加
	spendersRela.Spending = spendersRela.Spending + income
	// 对方刚加入不增加支出
	if spendersMonad.Id == spendersRela.CurrentMonad {
		if spendersRela.Spending == income && spendersRela.Income == 0 {
			spendersRela.Spending = 0
		}
	}
	// 别人的主单，关系户状态
	if spenderMainMonad != nil {
		// 因为未完成任务 4
		if spendersRela.Status == RELA_STATUS_FOUR {
			updateStatusIfTasks(spendersRela, spenderMainMonad)
		}
		// 因为对方单子是冻结状态2，并且是出新单
		if spendersRela.Status == RELA_STATUS_FREEZE && spendersMonad.Class == 0 {
			spenderMainMonad.UnfreezePeriodCount = spenderMainMonad.UnfreezePeriodCount - 1
			if spenderMainMonad.UnfreezePeriodCount == 0 {
				sysNow := time.Now().Local()
				// 解冻时间期限
				isOk := sysNow.Before(spenderMainMonad.UnFreeze)
				// 是否解冻对方主单
				if isOk { // 解冻期之内
					spendersRela.Status = RELA_STATUS_NORMAL
				}
			}
		}
		//没有激活时需要激活 0
		if spendersRela.Status == RELA_STATUS_UNBORN {
			spendersRela.Status = RELA_STATUS_NORMAL

		}
	}
	// 对方单子小于7级时加1级
	if spendersMonad.Class < 7 {
		spendersMonad.Class = spendersMonad.Class + 1
	}
	if spendersMonad.IsMain == 0 {
		spendersRela.Status = 1
	}
	spenderMainMonad.Edit()
	spendersMonad.Edit()
	spendersRela.Edit()
	// 删除待确认信息
	audit.Del(audit.Id)

	// 自己出局了，不会产生升级了
	if myRela.Status == RELA_STATUS_Retired {
		return false
	}

	if audit.Special == 1 {
		return true
	}
	return true
}

// 此单子有任务未完成
func currentMondUnfinished(mon *model.Monad) bool {
	aud := new(model.Audit).ByUpgargeMonad(mon.Id)
	if aud == nil {
		return false
	}
	if aud.Id > 0 {
		return true
	}
	return false
}

// 根据audits数组获取人对方信息
func findAuditorInfoForIds(aus []*model.Audit) map[string]interface{} {

	results := make(map[string]interface{})
	for _, audit := range aus {
		if audit != nil {
			if audit.Id > 0 {
				result := ""
				targetUser, _, _ := findURM(audit.ProposerRelationalId)
				result += targetUser.Mobile + "|"
				result += targetUser.Wechat + "|"
				result += targetUser.Alias + "|"
				refIndex := 0
				refIndex = audit.ProposerCount + 1
				sum := strconv.FormatInt(INCOME[refIndex], 10)
				result += sum + "|"
				result += audit.Create.String()
				ref := strconv.FormatInt(audit.Id, 10)
				results[ref] = result
			}
		}
	}
	return results
}

//  转换为审核者的推荐人信息信息
func changeAuditorReferrerInfo(au *model.Audit) string {
	resultStr := ""
	if au.Special == 1 {
		return resultStr
	}
	auditorRela := new(model.Relational).ById(au.RelationalId)
	var referrerUser = new(user.User)
	if auditorRela.Referrer == "top" {
		relaAdmi := new(manage.Relaadmin).FindByRelaId(auditorRela.Id)
		tmpId := strconv.FormatInt(relaAdmi.Ssoid, 10)
		referrerUser = findUserById(tmpId)
		if referrerUser.Id == 0 {
			referrerUser = findUserById("1")
		}
		resultStr += referrerUser.Mobile + "|"
		resultStr += referrerUser.Wechat + "|"
		resultStr += referrerUser.Alias + "|"
		resultStr += "全天"
	} else {
		tmpId, _ := strconv.ParseInt(auditorRela.Referrer, 10, 64)
		auditorRelaReferrer := auditorRela.ById(tmpId)
		userRefId := strconv.FormatInt(auditorRelaReferrer.SsoId, 10)
		referrerUser = findUserById(userRefId)
		resultStr += referrerUser.Mobile + "|"
		resultStr += referrerUser.Wechat + "|"
		resultStr += referrerUser.Alias + "|"
		resultStr += auditorRelaReferrer.Freetime
	}
	return resultStr
}

//  转换为审核者信息
func changeAuditorInfo(au *model.Audit) string {
	resultStr := ""
	var auUserId string
	if au.Special == 1 {
		relaAdmi := new(manage.Relaadmin).FindBySsoId(au.Sso)
		auUserId = strconv.FormatInt(relaAdmi.Ssoid, 10)
	} else {
		auUserId = strconv.FormatInt(au.Sso, 10)
	}
	auditorUser := findUserById(auUserId)
	consume := INCOME[au.ProposerCount+1]
	resultStr += auditorUser.Mobile + "|"
	resultStr += auditorUser.Wechat + "|"
	resultStr += auditorUser.Alias + "|"
	resultStr += strconv.FormatInt(consume, 10) + "|"
	if au.Special == 0 {
		auditorRela := new(model.Relational).ById(au.RelationalId)
		resultStr += auditorRela.Freetime
	} else {
		resultStr += "全天"
	}
	return resultStr
}

// 转换为提交者信息
func changeProposerInfo(au *model.Audit) string {
	resultStr := ""
	refId := strconv.FormatInt(au.ProposerSso, 10)
	tmpUser := findUserById(refId)
	consume := INCOME[au.ProposerCount+1]
	resultStr += tmpUser.Mobile + "|"
	resultStr += tmpUser.Wechat + "|"
	resultStr += tmpUser.Alias + "|"
	resultStr += strconv.FormatInt(consume, 10) + "|"
	if au.Special == 0 {
		tmpRela := new(model.Relational).ById(au.RelationalId)
		resultStr += tmpRela.Freetime
	}
	return resultStr
}

package controllers

import (
	"encoding/json"
	model "interaction/models"
	manage "interaction/models/manage"
	"math"
	"my/util"
	"net/http"
	"sso/user"
	"strconv"
	"strings"

	"github.com/ant0ine/go-json-rest/rest"
)

var conf *util.Config = nil

func SetConfig(filepath string) *util.Config {
	if conf != nil {
		return conf
	}
	conf = new(util.Config)
	conf.Filepath = filepath
	return conf
}

func GetConfig() *util.Config {
	return conf
}

// 校验用户请求信息
func validateUserInfo(req *http.Request) (bool, *user.User) {

	cookie, err := req.Cookie("token")
	if cookie == nil {
		return false, nil
	}
	if err != nil || cookie.Value == "" {
		return false, nil
	}
	conf = GetConfig()
	getTokenUrl := conf.Get("sso", "url") + conf.Get("sso", "byToken")
	userInfoByte := util.GetUserInfo(getTokenUrl, cookie.Value)

	if len(userInfoByte) < 1 {
		return false, nil
	}
	//my user info
	userA := &user.User{}
	json.Unmarshal(userInfoByte, &userA)
	if userA == nil {
		return false, nil
	}
	if userA.Id == 0 {
		return false, nil
	}

	return true, userA
}

// 获取自己需要审核的待办
func findMyTodos(myRelaId int64) map[string]string {
	aus, ausCount := new(model.Audit).ByRela(myRelaId)
	sweet := make(map[string]string)

	if ausCount == 0 {
		return nil
	}
	for _, task := range aus {
		if task == nil {
			continue
		}
		if task.Id == 0 {
			continue
		}
		id := strconv.FormatInt(task.Id, 10)
		sweet[id] = changeProposerInfo(task)
	}
	return sweet

}

// 检查任务是否完成，未完成将设置为冻结状态，完了就设置为正常
func updateStatusIfTasks(rela *model.Relational, mainM *model.Monad) (*model.Relational, *model.Monad, int, int, int) {
	firstT, secondT, threeT := findTaskForRelaId(rela.Id)
	if mainM == nil || rela == nil {
		return rela, mainM, 0, 0, 0
	}

	mesA := findAuditsLen(firstT)
	mesB := findAuditsLen(secondT)
	mesC := findAuditsLen(threeT)

	limitA, _ := strconv.Atoi(conf.Get("common", "taskLimitA"))
	limitB, _ := strconv.Atoi(conf.Get("common", "taskLimitB"))
	limitC, _ := strconv.Atoi(conf.Get("common", "taskLimitC"))
	//	fmt.Printf("1层：%v;----2层：%v;----3层：%v;\n", mesA, mesB, mesC)
	//	fmt.Printf("1.limit：%v;----2.limit：%v;----3.limit：%v;\n", limitA, limitB, limitC)

	if mesA >= limitA || mesB >= limitB || mesC >= limitC {
		rela.Status = RELA_STATUS_FOUR
		mainM.State = RELA_STATUS_FOUR
	}
	if rela.Status == RELA_STATUS_FOUR {
		if mesA < limitA && mesB < limitB && mesC < limitC {
			rela.Status = RELA_STATUS_NORMAL
			mainM.State = RELA_STATUS_NORMAL
		}

	}

	rela.Edit()
	mainM.Edit()

	return rela, mainM, mesA, mesB, mesC
}

func findAuditsLen(aus []*model.Audit) int {
	result := 0

	for _, au := range aus {
		if au == nil {
			continue
		}
		if au.Id > 0 {
			result++
		}
	}
	return result
}

// 寻找提交者的任务
func findTaskForRelaId(relaid int64) (firstT, secondT, threeT []*model.Audit) {

	aus, ausCount := new(model.Audit).NotOk(relaid)
	// 检查任务情况是否需要冻结
	//	fmt.Println(len(aus))
	//	fmt.Println(aus)
	firstT = make([]*model.Audit, ausCount)
	secondT = make([]*model.Audit, ausCount)
	threeT = make([]*model.Audit, ausCount)
	if ausCount == 0 {
		return firstT, secondT, threeT
	}
	for k, m := range aus {
		if m == nil {
			continue
		}
		if m.Id == 0 {
			continue
		}
		switch m.ProposerCount {
		case 7:
			threeT[k] = m
		case 6:
			threeT[k] = m
		case 5:
			threeT[k] = m
		case 4:
			secondT[k] = m
		case 3:
			secondT[k] = m
		case 2:
			firstT[k] = m
		case 1:
			firstT[k] = m
		case 0:
			firstT[k] = m
		}
	}
	return firstT, secondT, threeT

}

// 根据未被对方审核的(自己)所有单子和
//自己所有单子查找级别为0的审核记录
func findMonadByMyauditsInMyMonadsAndClassIsZero(myAudits []*model.Audit, myMonads []*model.Monad) *model.Audit {
	if len(myAudits) == 0 || len(myMonads) == 0 {
		return nil
	}

	for i := 0; i < len(myAudits); i++ {
		for j := 0; j < len(myMonads); j++ {
			myA := myAudits[i]
			myM := myMonads[j]
			if myA == nil || myM == nil {
				continue
			}
			if myA.ProposerMonadId == myM.Id && myM.Class == 0 {
				return myA
			}
		}
	}
	return nil
}

// 循环查找上几层单子
func findParentMonad(monad *model.Monad, layer int) *model.Monad {
	targetMonad := new(model.Monad)
	targetMonad = monad
	for k := 1; k <= layer; k++ {
		tmpMon := targetMonad.ById(targetMonad.ParentMonad)
		if tmpMon == nil {
			return nil
		}
		targetMonad = tmpMon
	}
	return targetMonad
}

// 根据 id 查找user info
func findUserById(id string) *user.User {
	conf := GetConfig()
	if id == "top" {
		return nil
	}
	tmpUser := new(user.User)
	ssoUrl := conf.Get("sso", "url")
	tmpUserByte := util.Get(ssoUrl + "/find?k=id&v=" + id)
	json.Unmarshal(tmpUserByte, &tmpUser)
	return tmpUser
}

// 根据 mob 查找user info
func findUserByMob(mob string) *user.User {
	conf := GetConfig()
	tmpUser := new(user.User)
	ssoUrl := conf.Get("sso", "url")
	tmpUserByte := util.Get(ssoUrl + "/find?k=mob&v=" + mob)

	json.Unmarshal(tmpUserByte, &tmpUser)
	return tmpUser
}

// 根据token查找user info

func byToken(token string, userRef *user.User) {
	conf := GetConfig()
	// Access token corresponding user information
	byTokenUrl := conf.Get("sso", "url") + conf.Get("sso", "byToken")
	userInfoByte := util.GetUserInfo(byTokenUrl, token)
	if len(userInfoByte) < 1 {
		//		fmt.Println("not user info.")
		return
	}
	json.Unmarshal(userInfoByte, &userRef)
}

// // 手机，微信，昵称，空余时间
func resultAuditerInfo(rela *model.Relational, userA *user.User) map[string]string {
	info := make(map[string]string)
	info["mob"] = userA.Mobile
	info["alias"] = userA.Alias
	info["wechat"] = userA.Wechat
	info["free"] = rela.Freetime
	return info
}

// 返回用户手机，微信，昵称，剩余时间
func resultUserInfo(userA *user.User) map[string]string {
	info := make(map[string]string)
	info["mob"] = userA.Mobile
	info["alias"] = userA.Alias
	info["wechat"] = userA.Wechat
	info["free"] = "7:00-23:00点，微信无应答，可来电！"
	return info
}

// 获取指定sso id的帐号
func resultAssignUserInfo(soid int64) map[string]string {
	ssoid := strconv.FormatInt(soid, 10)
	user := findUserById(ssoid)
	info := make(map[string]string)
	info["mob"] = user.Mobile
	info["alias"] = user.Alias
	info["wechat"] = user.Wechat
	info["free"] = "7:00-23:00点，微信无应答，可来电！"
	return info
}

// 根据top rela id 查找管理人员的信息
func findManageUserInfoByTopRelaId(topRelaId int64) map[string]string {
	info := make(map[string]string)
	info["mob"] = ""
	info["alias"] = ""
	info["wechat"] = ""
	info["free"] = "7:00-23:00点，微信无应答，可来电！"
	manageUser := new(manage.Relaadmin).FindByRelaId(topRelaId)
	if manageUser == nil {
		return info
	}
	userId := strconv.FormatInt(manageUser.Ssoid, 10)
	user := findUserById(userId)
	info["mob"] = user.Mobile
	info["alias"] = user.Alias
	info["wechat"] = user.Wechat
	return info
}

// 根据股东rela账户查找出特定特定推荐人帐号
func findSpecificReferrer(relaid int64) map[string]string {
	info := make(map[string]string)
	specRef := new(manage.Relaadmin)
	specRef = specRef.FindByRelaId(relaid)
	if specRef == nil {
		return info
	}
	ssoid := strconv.FormatInt(specRef.Ssoid, 10)
	userinfo := findUserById(ssoid)
	info["mob"] = userinfo.Mobile
	info["alias"] = userinfo.Alias
	info["wechat"] = userinfo.Wechat
	info["free"] = "7:00-23:00点，微信无应答，可来电！"
	return info
}

// 根据单子寻找主人user,rela,mainMonad
func findURM(relationalId int64) (*user.User, *model.Relational, *model.Monad) {
	rela := new(model.Relational)
	rela = rela.ById(relationalId)
	if rela == nil {
		return nil, nil, nil
	}
	if rela.Id == 0 {
		return nil, nil, nil
	}
	ssoId := strconv.FormatInt(rela.SsoId, 10)
	tmpUser := findUserById(ssoId)
	//
	tmpMainMonad := new(model.Monad)
	tmpMainMonad = tmpMainMonad.ById(rela.CurrentMonad)
	return tmpUser, rela, tmpMainMonad
}

// 根据user 信息找推荐人user信息
func findUserInfoByUser(userA *user.User) *user.User {
	tmpRela := new(model.Relational)
	tmpRela.ById(userA.Id)
	return findUserById(tmpRela.Referrer)
}

// 更具request 寻找user
func findUser(req *rest.Request) *user.User {
	conf = GetConfig()

	cookie, _ := req.Cookie("token")
	if cookie == nil {
		return nil
	}

	byTokenUrl := conf.Get("sso", "url") + conf.Get("sso", "byToken")
	userInfoByte := util.GetUserInfo(byTokenUrl, cookie.Value)
	userA := &user.User{}
	json.Unmarshal(userInfoByte, &userA)
	if userA == nil {
		return nil
	}
	return userA

}

// 根据token 查找 user、relational、mainMonad
func findUserInfo(req *rest.Request) map[string][]byte {
	conf = GetConfig()
	result := make(map[string][]byte)
	result["user"] = nil
	result["relational"] = nil
	result["mainMonad"] = nil
	//
	cookie, _ := req.Cookie("token")
	if cookie == nil {
		return result
	}

	byTokenUrl := conf.Get("sso", "url") + conf.Get("sso", "byToken")
	userInfoByte := util.GetUserInfo(byTokenUrl, cookie.Value)
	userA := &user.User{}
	json.Unmarshal(userInfoByte, &userA)

	if userA == nil {
		return result
	}
	result["user"], _ = json.Marshal(userA)
	relational := new(model.Relational)
	if userA.Mobile != "" {
		relational = relational.ByMob(userA.Mobile)

	}
	if relational == nil {
		//		fmt.Println("relational is not 3")
		return result
	}
	result["relational"], _ = json.Marshal(relational)
	if relational.CurrentMonad < 1 {
		//		fmt.Println("Current Main Monad is not.")
		return result
	}
	myMainMonad := model.NewMonad()
	myMainMonad = myMainMonad.ById(relational.CurrentMonad)
	if myMainMonad != nil {
		result["mainMonad"], _ = json.Marshal(myMainMonad)
	}
	return result
}

// 寻找单子的下层单子
func findChildsByMonad(mainMonad *model.Monad) *model.Monad {
	myRelaId := mainMonad.Pertain
	monads := []*model.Monad{mainMonad}
	dismissal, _ := strconv.Atoi(conf.Get("common", "dismissal")) // 7
	mulriple, _ := strconv.Atoi(conf.Get("common", "mulriple"))   // 3
	for k := 1; k <= dismissal; k++ {

		childs, absent, isEnd := findChilds(monads, k, mulriple)

		if isEnd {
			monads = childs
			continue
		}
		if absent.Pertain != myRelaId { // 空缺的父级不是自己，返回空缺位
			return absent
		}
		k = 1
		monads = childs
		//childs, absent, isEnd = findChilds(childs, 0, mulriple)

	}
	return nil
}

//根据主单向下找空位
func findMonadChilds(mainMonad *model.Monad) *model.Monad {
	monads := []*model.Monad{mainMonad}
	dismissal, _ := strconv.Atoi(conf.Get("common", "dismissal")) // 7
	mulriple, _ := strconv.Atoi(conf.Get("common", "mulriple"))   // 3
	for k := 1; k <= dismissal; k++ {
		childs, absent, isEnd := findChilds(monads, k, mulriple)
		if isEnd {
			monads = childs
		} else {
			return absent
		}
	}
	return nil
}

// 根据多个monads寻找下级monads
func findChilds(monads []*model.Monad, layer, mulriple int) ([]*model.Monad, *model.Monad, bool) {
	length := math.Pow(float64(3), float64(layer))
	count := int(length)
	allChilds := make([]*model.Monad, count)
	absent := &model.Monad{}
	for j := 0; j < len(monads); j++ {
		var childs []*model.Monad
		var sum int
		if monads[j] == nil {
			continue
		}
		childs, sum = monads[j].MonadChildsByParentMonad(monads[j].Id)
		if sum == mulriple {
			for k := 0; k < len(childs); k++ {
				allChilds = append(allChilds, childs[k])
			}
		} else {
			for k := 0; k < len(childs); k++ {
				if childs[k] != nil {
					if childs[k].Id > 0 {
						allChilds = append(allChilds, childs[k])
					}
				}
			}

			if monads[j] != nil && monads[j].Id > 0 {
				absent = monads[j]
				return allChilds, absent, false
			}
		}
	}
	return allChilds, absent, true
}

// 根据rela id获取主单
func findParentMainMonad(RelaId int64) *model.Monad {
	rela := &model.Relational{}
	rela = rela.ById(RelaId) //父级rela
	if rela == nil {
		//		fmt.Printf("find rela = nil , but id = %v\n", RelaId)
		return nil
	}
	mainMonad := &model.Monad{}
	mainMonad = mainMonad.ById(rela.CurrentMonad)
	if mainMonad == nil {
		return nil
	}
	return mainMonad
}

// 寻找提交者信息，根据mes id ，返回monad,relational,message
func findMonadByMessageId(mesId string) (*model.Monad, *model.Relational, *model.Message) {

	mes := new(model.Message)
	mesIdint64, _ := strconv.ParseInt(mesId, 10, 64)
	mes.Id = mesIdint64
	mes.ById()
	//
	myMonad := new(model.Monad)
	myMonad = myMonad.ById(mes.MId)
	//
	myRela := new(model.Relational)
	myRela = myRela.FindById(mes.RId)
	//
	return myMonad, myRela, mes
}

// 主单需要限制升级
func mainMonadTask(myAuMonad *model.Monad, refMon *model.Monad) (isOk bool, fCount, sCount int) {
	isOk = false
	fCount, sCount = findRecommandInfo(myAuMonad.Pertain, refMon)

	switch myAuMonad.Task + 1 {
	case 3:
		if fCount > 0 {
			isOk = true
		}
	case 4:
		if (fCount + sCount) > 2 {
			isOk = true
		}
	case 5:
		if (fCount + sCount) > 4 {
			isOk = true
		}
	case 6:
		if ((fCount + sCount) > 6) && (fCount > 1) && (sCount > 1) {
			isOk = true
		}
	}

	return
}

// 根据某个用户获取待办详细列表
func findTodoDetailListForAudits(aus []*model.Audit) ([]*model.Relational, []*model.Monad, []*user.User) {
	if len(aus) == 0 {
		return nil, nil, nil
	}
	var ssoIdStr, monadIdStr, relaIdStr string

	for _, a := range aus {
		if a != nil {
			if a.Id > 0 {
				ssoId := strconv.FormatInt(a.Id, 10)
				monId := strconv.FormatInt(a.MonadId, 10)
				relaId := strconv.FormatInt(a.RelationalId, 10)
				//mesId := strconv.FormatInt(a.Messageid, 10)
				//
				ssoIdStr += ssoId + "|"
				monadIdStr += monId + "|"
				relaIdStr += relaId + "|"
				//messIdStr += mesId + "|"

			}
		}
	}
	ssoIds := strings.Split(ssoIdStr, "|")
	monadIds := strings.Split(monadIdStr, "|")
	relaIds := strings.Split(relaIdStr, "|")
	//messIds := strings.Split(messIdStr, "|")
	return findTodoDetailList(ssoIds, monadIds, relaIds)

}

// 获取待办详细列表(详细审核列表)
func findTodoDetailList(ssoIds, monadIds, relaIds []string) ([]*model.Relational, []*model.Monad, []*user.User) {
	monads, _ := new(model.Monad).ByIds(strings.Join(monadIds, ","))
	relas, _ := new(model.Relational).ByIds(strings.Join(relaIds, ","))
	//
	ssoUrl := conf.Get("sso", "url")
	pageUrl := conf.Get("sso", "byIds")
	url := ssoUrl + pageUrl + "?ids=" + strings.Join(ssoIds, "|")
	//
	ssos := util.Get(url)
	users := make([]*user.User, len(ssoIds))
	json.Unmarshal(ssos, &users)

	return relas, monads, users

}

// 是否存在一条未审核的0级单子
func canCreateMonad(relationalId int64) (bool, *model.Audit) {
	audit := &model.Audit{}
	monad := &model.Monad{}
	// 自己的单子未被别人审核状况
	myAudits, myCount := audit.ByPropRela(relationalId, 0, 0)
	if myCount == 0 {
		return true, nil
	}
	// 自己单子级别为0
	myMonads, myMonadsCount := monad.MonadsByPertain(relationalId, 0)
	if myMonadsCount == 0 {
		return true, nil
	}
	tmpAu := new(model.Audit)
	tmpAu = findMonadByMyauditsInMyMonadsAndClassIsZero(myAudits, myMonads)
	if tmpAu == nil { // 已出的新单子全部被对方审核通过
		return true, nil //可以出新单子
	} else { // 有一条没有被审核的新单子
		return false, tmpAu //不可以出单
	}
}

package controllers

import (
	"encoding/json"

	model "interaction/models"
	"my/util"
	"sso/user"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

var lock sync.Mutex

// 获取升级任务需要的对方信息
// myMonad是需要升级的单子
func FindTask(rep rest.ResponseWriter, req *rest.Request) {
	stat, auid, _ := getParamForReq(req, "auid")
	if !stat {
		rep.WriteJson("")
		return
	}
	// 获取提交者
	refId, _ := strconv.ParseInt(auid, 10, 64)
	myTask, _ := new(model.Audit).ById(refId)
	//
	//
	result := make(map[string]interface{})
	// 审核者信息
	result["pi"] = changeAuditorInfo(myTask)
	// 审核者推荐人信息
	result["pri"] = changeAuditorReferrerInfo(myTask)
	rep.WriteJson(result)
}

// 根据auid id 提交任务
func SubmitTask(rep rest.ResponseWriter, req *rest.Request) {
	stat, auid, myUser := getParamForReq(req, "auid")
	if !stat {
		rep.WriteJson("1")
		return
	}
	refId, _ := strconv.ParseInt(auid, 10, 64)
	audit, _ := new(model.Audit).ById(refId)
	// 安全提交
	if myUser.Id != audit.ProposerSso {
		rep.WriteJson("2")
		return
	}
	audit.Status = 0
	audit.Create = time.Now().Local()
	audit.UpdateByColName("status", "create")
	rep.WriteJson("ok")
}

// 出单
func NewTask(res rest.ResponseWriter, req *rest.Request) {

	result := findUserInfo(req)
	lock.Lock()
	defer lock.Unlock()
	myUser := new(user.User)
	myRelational := new(model.Relational)
	myMainMonad := new(model.Monad)
	json.Unmarshal(result["user"], &myUser)
	json.Unmarshal(result["relational"], &myRelational)
	json.Unmarshal(result["mainMonad"], &myMainMonad)
	//
	// Add monad start
	//
	resultClent := make(map[string]interface{}) // 返回信息
	myMonad := model.NewMonad()
	myMonad.Pertain = myRelational.Id
	myMonad.IsMain = 0
	myMonad.MainMonad = myRelational.CurrentMonad
	//
	parentRela := new(model.Relational)
	parMonad := new(model.Monad)
	flag := false
	//var state bool
	if myRelational.CurrentMonad == 0 {
		parentRela, parMonad, flag = newMain(myMonad, myRelational) // 推荐人rela，主单
		resultClent["isMain"] = true
		// 出主单失败
		if !flag {
			resultClent["fail"] = true
			resultClent["status"] = 5
			res.WriteJson(resultClent)
			return
		}
	} else {
		if myMainMonad.State == 0 {
			resultClent["fail"] = true
			resultClent["status"] = 1
			res.WriteJson(resultClent)
			return
		}
		if b, _ := canCreateMonad(myRelational.Id); !b {
			resultClent["fail"] = true
			resultClent["status"] = 2
			res.WriteJson(resultClent)
			return
		}

		parentRela, parMonad, flag = newSub(myMonad, myRelational, myMainMonad)

		// 没有位置
		if !flag {
			resultClent["fail"] = true
			resultClent["status"] = 4
			res.WriteJson(resultClent)
			return
		}
	}

	// 因为对方上级单子处于冻结状态
	parMainMonad := new(model.Monad).ById(parentRela.CurrentMonad)
	if parentRela.Referrer != "top" {
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
			// 组装对方信息
			resultClent["pi"] = resultAssignUserInfo(specialUserId)
			resultClent["status"] = 5
			res.WriteJson(resultClent)
			return
		}
	}

	// 寻找父级主人
	ssoid := strconv.FormatInt(parentRela.SsoId, 10)
	parentUser := findUserById(ssoid)
	// add audit
	createAuditForNewMonad(myMonad, parMonad, myRelational, parentRela, 0, 0)
	// 组装对方信息
	resultClent["pi"] = resultAuditerInfo(parentRela, parentUser)
	// 对方推荐人信息
	if parentRela.Referrer == "top" { //股东
		resultClent["pri"] = findSpecificReferrer(parentRela.Id)
	} else {
		tmpRelaId, _ := strconv.ParseInt(parentRela.Referrer, 10, 64)
		tmpUserRef, tmpRelaRef, _ := findURM(tmpRelaId)
		if tmpUserRef != nil && tmpRelaRef != nil {
			resultClent["pri"] = resultAuditerInfo(tmpRelaRef, tmpUserRef)
		}
	}
	resultClent["status"] = 9
	res.WriteJson(resultClent)

}

// 出主单，返回推荐信息
func newMain(myMonad *model.Monad, myRela *model.Relational) (*model.Relational, *model.Monad, bool) {
	myMonad.IsMain = 1 //设置为主单
	//获取父级主单
	cid, _ := strconv.ParseInt(myRela.Referrer, 10, 64)
	parentMainMonad := findParentMainMonad(cid)
	if parentMainMonad == nil {
		return nil, nil, false
	}
	if parentMainMonad.Id == 0 {
		return nil, nil, false
	}
	// 在父级主单下找空位
	parMonad := findMonadChilds(parentMainMonad)
	if parMonad == nil {
		return nil, nil, false
	}
	if parMonad.Id == 0 {
		return nil, nil, false
	}
	myMonad.ParentMonad = parMonad.Id

	myMainMonadId, _ := myMonad.Add()
	myRela.CurrentMonad = myMonad.Id
	myRela.PrevNewMonad = time.Now().Local()
	myRela.Edit()
	myMonad.MainMonad = myMainMonadId
	myMonad.Edit()
	// 添加审核任务
	parentRela := myRela.ById(parMonad.Pertain)
	return parentRela, parMonad, true
}

// 出子单
func newSub(myMonad *model.Monad, myRela *model.Relational, myMainMonad *model.Monad) (*model.Relational, *model.Monad, bool) {
	var parMonad *model.Monad = nil

	overall := conf.Get("common", "overall")
	// 全局查找单子空位
	if strings.ToLower(overall) == "true" {
		// 寻找本条线顶级37账号
		var topRelational *model.Relational = nil
		topRelational = findTopRelational(myRela)
		if topRelational == nil {
			//	fmt.Println("topRelational is nil")
			return nil, nil, false
		}
		topMonad := myMonad.ById(topRelational.CurrentMonad)
		if topMonad == nil {
			//	fmt.Println("topMonad is nil")
			return nil, nil, false
		}

		parMonad = findMonadChilds(topMonad)
		if myRela.Id == parMonad.Pertain && strings.ToLower(myRela.Referrer) != "top" {
			//	fmt.Println("myRela.Id == parMonad.Pertain")
			return nil, nil, false
		}

		// parMonad = findChildsByMonad(topMonad) // 在自己归属的主线下面找空位
	} else { // 在自己下面找空位
		parMonad = findChildsByMonad(myMainMonad) // 在自己主单下面找空位
	}

	if parMonad == nil {
		return nil, nil, false
	}
	if parMonad.Id == 0 {
		return nil, nil, false
	}
	myMonad.ParentMonad = parMonad.Id
	myMonad.Add()
	parentRela := myRela.ById(parMonad.Pertain)
	return parentRela, parMonad, true
}

func findTopRelational(myRela *model.Relational) *model.Relational {
	refRela := new(model.Relational)
	if myRela == nil {
		return nil
	}
	if myRela.Id == 0 {
		return nil
	}
	if strings.ToLower(myRela.Referrer) == "top" {
		return myRela
	}
	refRela = myRela
	for {

		refRelaIdId, _ := strconv.ParseInt(refRela.Referrer, 10, 64)
		refRela = myRela.ById(refRelaIdId)
		if refRela == nil {
			return nil
		}
		if refRela.Id == 0 {
			return nil
		}
		if strings.ToLower(refRela.Referrer) == "top" {
			return refRela
		}

	}
	return refRela
}

// 校验用户请求信息
func validateRole(req *rest.Request) (bool, *user.User) {
	cookie, err := req.Cookie("token")
	if err != nil || cookie.Value == "" {
		return false, nil
	}
	conf = util.GetConfig()
	byTokenUrl := conf.Get("sso", "url") + conf.Get("sso", "byToken")
	userInfoByte := util.GetUserInfo(byTokenUrl, cookie.Value)
	if len(userInfoByte) < 1 {
		return false, nil
	}
	//my user info
	userA := &user.User{}
	json.Unmarshal(userInfoByte, &userA)
	if userA == nil {
		return false, nil
	}
	return true, userA
}

// 获取get请求参数
func getParamForReq(req *rest.Request, paramName string) (bool, string, *user.User) {
	sta, userRef := validateRole(req)
	if sta {
		auidRef := req.PathParam(paramName)
		return true, auidRef, userRef
	}
	return false, "", userRef
}

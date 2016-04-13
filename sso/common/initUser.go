package common

import (
	"my/util"
	"sso/user"
	"time"
)

func InitUser() {
	new(user.State).DelAll()
	new(user.User).DelAll()
	newUserTemplate("18620131415", "18620131415")
	newMangeUserTemplate("13790226216", "13790226216")
	whiteUserTemplate("17750662398", "17750662398")
	newUserTemplate("18059244379", "18059244379")
	newUserTemplate("18650710067", "18650710067")
	//	newUserTemplate("18620131415", "78Rt56[+99")
	//	newMangeUserTemplate("13790226216", "78Rt56[+99")
	//	whiteUserTemplate("17750662398", "3737.io@mm")
	//	newUserTemplate("18059244379", "3737.io@mxg")
	//	newUserTemplate("18650710067", "3737.io@mff")
}

// 管理者帐号
func newMangeUserTemplate(mob, mm string) {
	user := new(user.User)
	user.Id = 0
	user.Username = "test"
	user.Password = util.Md5Encode(mm)
	user.Alias = "简洁的代言"
	user.Mobile = mob
	user.Alipay = mob + "@qq.com"
	user.Wechat = "heimawangzi_com"
	user.QQ = 611041314
	user.Email = mob + "@3737.io"
	user.City = "北京"
	user.Address = "中关村055号"
	user.Sex = 1
	user.Identity = "6688"
	user.Create = time.Now()
	user.Last = time.Now()
	user.Add()
}

func newUserTemplate(mob, mm string) {
	user := new(user.User)
	user.Id = 0
	user.Username = mob
	user.Password = util.Md5Encode(mm)
	user.Alias = mob
	user.Mobile = mob
	user.Alipay = mob + "@3737.io"
	user.Wechat = mob
	user.QQ = 611041314
	user.Email = mob + "@3737.io"
	user.City = "北京"
	user.Address = "中关村051号"
	user.Sex = 1
	user.Identity = "3737.io"
	user.Create = time.Now()
	user.Last = time.Now()
	user.Add()
}

// 空号信息
func whiteUserTemplate(mob, mm string) {
	user := new(user.User)
	user.Id = 0
	user.Username = mob
	user.Password = util.Md5Encode(mm)
	user.Alias = "千山万水"
	user.Mobile = mob
	user.Alipay = mob + "@3737.io"
	user.Wechat = mob
	user.QQ = 611041314
	user.Email = mob + "@3737.io"
	user.City = "北京"
	user.Address = "中关村01号"
	user.Sex = 1
	user.Identity = "3737.io"
	user.Create = time.Now()
	user.Last = time.Now()
	user.Add()
}

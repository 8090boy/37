package user

import (
	"fmt"
	"my/util"
	"strconv"
	"time"
)

// SSO User
type User struct {
	Id       int64     //json:",string"`
	Username string    //用户名
	Password string    //密码
	Alias    string    //昵称
	Mobile   string    //手机号
	Alipay   string    //支付宝
	Wechat   string    //微信
	QQ       int64     `json:",string"` //QQ
	Email    string    //邮箱
	City     string    //城市
	Address  string    //详细地址
	Sex      int       `json:",string"` //性别
	Identity string    //身份证号
	Create   time.Time `json:",string"` //注册日期
	Last     time.Time `json:",string"` //上次登录日期
}

func NewUser() User {
	u := User{}
	return u
}

func (cate *User) Add() (*User, int64, error) {
	refId, err := util.Eng.Insert(cate)
	if err != nil {
		// id is identity
		return nil, 0, err
	}
	return cate, refId, nil
}

func (cate *User) DelAll() {
	sql := "DELETE FROM user;"
	sql1 := "ALTER TABLE user AUTO_INCREMENT=1;"
	_, err := util.Eng.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	_, err = util.Eng.Exec(sql1)
	if err != nil {
		fmt.Println(err)
	}
}

func (data *User) Del(id int64) error {

	data.Id = id
	_, err := util.Eng.Delete(data)
	return err
}

func (cate *User) ById(refId string) *User {

	id, err := strconv.ParseInt(refId, 10, 64)
	if id == 0 {
		return nil
	}
	if err != nil {
		return nil
	}
	_, err = util.Eng.Id(id).Get(cate)
	if err != nil {
		return nil
	}
	return cate
}

func (data *User) Edit() error {
	if data.Id == 0 {
		return nil
	}
	fmt.Println(data)
	_, err := util.Eng.Id(data.Id).Update(data)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (data *User) Update(dataRef User) error {
	if dataRef.Id == 0 {
		return nil
	}

	_, err := util.Eng.Id(dataRef.Id).Update(dataRef)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (cate *User) ByMobile() *User {
	sql := "select * from user as U where U.mobile =?  limit 1"
	slice, _ := util.Eng.Query(sql, cate.Mobile)
	if len(slice) == 0 {
		return nil
	}
	cate = Compound(slice, 0)
	return cate
}

func (user *User) ByEmail() *User {
	sql := "select * from user as U where U.email  =? limit 1"
	Slice, _ := util.Eng.Query(sql, user.Email)
	if len(Slice) == 0 {
		return nil
	}
	user = Compound(Slice, 0)
	return user
}

func (user *User) ByUsername() *User {
	sql := "select * from user as U where U.username =?  limit 1"
	Slice, _ := util.Eng.Query(sql, user.Username)
	if len(Slice) == 0 {
		return nil
	}
	user = Compound(Slice, 0)
	return user
}

func Compound(resultsSlice []map[string][]byte, no int) *User {
	user := new(User)
	var err error

	Slice := resultsSlice[no]
	for k, v := range Slice {
		val := string(v)
		switch k {
		case "id":
			user.Id, err = strconv.ParseInt(val, 10, 64)
		case "username":
			user.Username = val
		case "password":
			user.Password = val
		case "alias":
			user.Alias = val
		case "mobile":
			user.Mobile = val
		case "alipay":
			user.Alipay = val
		case "wechat":
			user.Wechat = val
		case "qq":
			user.QQ, err = strconv.ParseInt(val, 10, 64)
		case "email":
			user.Email = val
		case "city":
			user.City = val
		case "address":
			user.Address = val
		case "sex":
			user.Sex, err = strconv.Atoi(val)
		case "identity":
			user.Identity = val
		case "create":
			user.Create, err = time.Parse(util.TimeFormat, val)
		case "last":
			user.Last = time.Now()
		}
		if err != nil {
			return nil
		}
	}

	return user
}

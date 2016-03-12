package main

import (
	"fmt"
	"my/util"
	"time"
)

func main() {
	user := "platform3737@163.com"
	password := "16Ui89rnMpEry"
	host := "smtp.163.com:25"
	to := "0411dapeng@163.com"

	subject := "3737.io 密码重置"
	//
	datetime := time.Now().Format(util.TimeFormat)
	body := `
<includetail>
<table width="100%" height="100%" border="0" cellspacing="0" cellpadding="0" align="center" style="background:#FF9800">
<tbody>
<tr>
<td width="323" style="  color:#ffffff;   font-size:2em;">尊敬的3737.io会员：Akgl <hr>
<span>您在` + datetime + `提交的找回密码，点击这里<a href="http://3737.io/?new=13416536546513" target="_blank" style="color:#fff; font-size:1.4em;  text-decoration:underline">密码重置</a>！</span>                    </td>
</tr>
</tbody>
</table>
</includetail>
`

	err := util.SendToMail(user, password, host, to, subject, body)
	if err != nil {
		fmt.Printf("Send mail error ! detail: %v\n", err)
	}
}
